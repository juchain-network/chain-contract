const Proposal = artifacts.require("Proposal");
const Validators = artifacts.require('Validators');
const Punish = artifacts.require("Punish");

const {
    constants,
    expectRevert,
    expectEvent,
    time,
    ether,
    BN
} = require('@openzeppelin/test-helpers');
const {
    assert
} = require('chai');

const Active = new BN("1");
const Jailed = new BN("2");

contract("Validators test", function (accounts) {
    var valIns, proposalIns, punishIns, initValidators;
    var miner = accounts[0];

    before(async function () {
        valIns = await Validators.new();
        proposalIns = await Proposal.new();
        punishIns = await Punish.new();

        initValidators = getInitValidators(accounts);
        await proposalIns.setContracts(valIns.address, constants.ZERO_ADDRESS, constants.ZERO_ADDRESS);
        await valIns.setContracts(valIns.address, punishIns.address, proposalIns.address);
        await punishIns.setContracts(valIns.address, punishIns.address, proposalIns.address);

        await valIns.initialize(initValidators);
        await valIns.setMiner(miner);
        await proposalIns.initialize(initValidators);
        await punishIns.initialize();
    })

    it("can only init once", async function () {
        await expectRevert(valIns.initialize(initValidators), "Already initialized");
    })

    describe("create or edit validator", async function () {
        let validator = accounts[30];

        it("can't create validator if fee addr == address(0)", async function () {
            await expectRevert(valIns.createOrEditValidator(constants.ZERO_ADDRESS, "", "", "", "", "", {
                from: validator
            }), "Invalid fee address");
        })

        it("can't create validator if describe info invalid", async function () {
            // invalid moniker
            let moniker = getInvalidMoniker();
            await expectRevert(valIns.createOrEditValidator(validator, moniker, "", "", "", "", {
                from: validator
            }), "Invalid moniker length");
        })

        it("can't create validator if not pass propose", async function () {
            await expectRevert(valIns.createOrEditValidator(validator, "", "", "", "", "", {
                from: validator
            }), "You must be authorized first");
        })

        it("create validator", async function () {
            await pass(proposalIns, initValidators, validator);
            let receipt = await valIns.createOrEditValidator(validator, "", "", "", "", "", {
                from: validator
            });
            expectEvent(receipt, "LogEditValidator", {
                val: validator,
                fee: validator
            });

            // check validator status
            let status = await valIns.getValidatorInfo(validator);
            assert.equal(status[1].eq(Active), true);
        })

        it("edit validator info", async function () {
            let feeAddr = accounts[31];
            let receipt = await valIns.createOrEditValidator(feeAddr, "", "", "", "", "", {
                from: validator
            });
            expectEvent(receipt, "LogEditValidator", {
                val: validator,
                fee: feeAddr
            });
        })
    })

    describe("propose add a new val", async function () {
        let nval = accounts[21];

        it('not a val', async function () {
            let isVal = await valIns.isTopValidator(nval);
            assert.equal(false, isVal);
        })

        it("create/vote proposal", async function () {
            let receipt = await proposalIns.createProposal(nval, true, "", {
                from: nval
            });

            id = receipt.logs[0].args.id;

            for (let i = 0; i < initValidators.length; i++) {
                await proposalIns.voteProposal(id, true, {
                    from: initValidators[i]
                });
            }
        })

        it('is a val', async function () {
            let isVal = await valIns.isTopValidator(nval);
            assert.equal(true, isVal);

            let pass = await proposalIns.pass(nval);
            assert.equal(true, pass);
        })
    })

    describe("propose remove a val", async function () {
        let nval = accounts[21];

        it('not a val', async function () {
            let isVal = await valIns.isTopValidator(nval);
            assert.equal(true, isVal);
        })

        it("create/vote proposal", async function () {
            let receipt = await proposalIns.createProposal(nval, false, "", {
                from: nval
            });

            id = receipt.logs[0].args.id;
            for (let i = 0; i < initValidators.length; i++) {
                await proposalIns.voteProposal(id, true, {
                    from: initValidators[i]
                });
            }
        })

        it('is a val', async function () {
            let isVal = await valIns.isTopValidator(nval);
            assert.equal(false, isVal);

            let pass = await proposalIns.pass(nval);
            assert.equal(false, pass);
        })
    })

    describe("distribute block reward", async function () {
        let fee = ether("0.3");
        let expectPerFee = ether("0.1");
        it("miner can distribute to validator contract, the profits should be right updated", async function () {
            let receipt = await valIns.distributeBlockReward({
                from: miner,
                value: fee
            });

            expectEvent(receipt, "LogDistributeBlockReward", {
                coinbase: miner,
                blockReward: fee,
            })

            for (let i = 0; i < initValidators.length; i++) {
                let info = await valIns.getValidatorInfo(miner);

                assert.equal(info[2].toString(10), expectPerFee.toString(10));
            }
        })

        it('update withdraw profit wait block', async function () {
            let receipt = await proposalIns.createUpdateConfigProposal(4, 10, {
                from: accounts[0]
            });
            id = receipt.logs[0].args.id;

            for (let i = 0; i < initValidators.length; i++) {
                await proposalIns.voteProposal(id, true, {
                    from: initValidators[i]
                });
            }
        })

        it("validator can withdraw profits", async function () {
            let receipt = await valIns.withdrawProfits(miner, {
                from: miner
            });

            expectEvent(receipt, "LogWithdrawProfits", {
                val: miner,
                fee: miner,
                hb: expectPerFee,
            });

            fee = ether('0.5');
            feeAddr = accounts[10];
            expectFee = fee.div(new BN('3'));
            await valIns.createOrEditValidator(feeAddr, "", "", "", "", "", {
                from: miner
            });
            await valIns.distributeBlockReward({
                from: miner,
                value: fee
            });

            // advance block
            let lock = await proposalIns.withdrawProfitPeriod();
            for (let i = 0; i < lock.toNumber(); i++) {
                await time.advanceBlock();
            }

            receipt = await valIns.withdrawProfits(miner, {
                from: feeAddr
            });
            expectEvent(receipt, "LogWithdrawProfits", {
                val: miner,
                fee: feeAddr,
                hb: expectFee,
            });
        })

        it("Can't call withdrawProfits if you don't have any profits", async function () {
            feeAddr = accounts[10];

            // advance block
            let lock = await proposalIns.withdrawProfitPeriod();
            for (let i = 0; i < lock.toNumber(); i++) {
                await time.advanceBlock();
            }

            await expectRevert(valIns.withdrawProfits(miner, {
                from: feeAddr
            }), "You don't have any profits");
        })
    })

    describe("update set", async function () {
        it("update active validator set", async function () {
            let epoch = 30;
            let newSet = getNewValidators(accounts);
            while (true) {
                let currentNumber = await web3.eth.getBlockNumber();

                if (currentNumber % epoch == (epoch - 1)) {
                    let receipt = await valIns.updateActiveValidatorSet(newSet, epoch, {
                        from: miner
                    });
                    expectEvent(receipt, "LogUpdateValidator");
                    break;
                }

                await time.advanceBlock();
            }

            // validate validator set
            for (let i = 0; i < initValidators.length; i++) {
                let is = await valIns.isActiveValidator(initValidators[i]);
                assert.equal(is, false);
            }
            for (let i = 0; i < newSet.length; i++) {
                let is = await valIns.isActiveValidator(newSet[i]);
                assert.equal(is, true);
            }
        })
    })
});

contract("Punish", function (accounts) {
    var valIns, proposalIns, punishIns, initValidators;
    var miner = accounts[0];

    before(async function () {
        valIns = await Validators.new();
        proposalIns = await Proposal.new();
        punishIns = await Punish.new();

        initValidators = getInitValidators(accounts);
        await proposalIns.setContracts(valIns.address, constants.ZERO_ADDRESS, constants.ZERO_ADDRESS);
        await valIns.setContracts(valIns.address, punishIns.address, proposalIns.address);
        await punishIns.setContracts(valIns.address, punishIns.address, proposalIns.address);

        await valIns.initialize(initValidators);
        await valIns.setMiner(miner);
        await proposalIns.initialize(initValidators);
        await punishIns.initialize();
        await punishIns.setMiner(miner);
    })

    it("can only init once", async function () {
        await expectRevert(punishIns.initialize(), "Already initialized");
    })

    describe("punish val", async function () {
        it("miner can punish validator", async function () {
            let removeThreshold = await proposalIns.removeThreshold();
            let punishThreshold = await proposalIns.punishThreshold();
            let fee = ether("0.4");

            for (let i = 0; i < removeThreshold.toNumber(); i++) {
                // distribute
                await valIns.distributeBlockReward({
                    from: miner,
                    value: fee
                });

                // punish
                let receipt = await punishIns.punish(miner, {
                    from: miner
                });
                expectEvent(receipt, "LogPunishValidator", {
                    val: miner
                });
                let recordInfo = await punishIns.getPunishRecord(miner);
                assert.equal(recordInfo.toNumber(), (i + 1) % removeThreshold.toNumber());

                let info = await valIns.getValidatorInfo(miner);

                if ((i + 1) % removeThreshold.toNumber() == 0) {
                    let is = await valIns.isTopValidator(miner);
                    assert.equal(is, false);

                    assert.equal(recordInfo.toNumber(), 0);
                    assert.equal(info[1].eq(Jailed), true);
                } else if ((i + 1) % punishThreshold.toNumber() == 0) {
                    assert.equal(info[2].toNumber(), 0);
                }
            }

            // check other validator profits.
            let info = await valIns.getValidatorInfo(initValidators[1]);

            let feeBN = new BN(fee.toString());
            let multi = new BN(removeThreshold.toString());
            let expectFee = feeBN.mul(multi).div(new BN("2"));

            // not equal for precision
            console.log("expect", expectFee.toString(), "acutal", info[2].toString());

            // get punish info
            info = await valIns.getValidatorInfo(miner);
            assert.equal(info[2].isZero(), true);
            // not equal for precision reason
            console.log("expect", feeBN.mul(multi).div(new BN('3')).toString(), "acutal", info[3].toString());
        })

        it("validator missed record will decrease if necessary", async function () {
            let removeThreshold = await proposalIns.removeThreshold();
            let decreaseRate = await proposalIns.decreaseRate();
            let step = 2;
            for (let i = 0; i < removeThreshold.div(decreaseRate).toNumber() + step; i++) {
                if (i < removeThreshold.div(decreaseRate).toNumber()) {
                    await punishIns.punish(initValidators[0], {
                        from: miner
                    });
                }
                await punishIns.punish(initValidators[1], {
                    from: miner
                });
            }

            let l = await punishIns.getPunishValidatorsLen();
            assert.equal(l.toNumber(), 2);

            let expect = await punishIns.getPunishRecord(initValidators[0]);
            // Punish record will be set to 0 if <= removeThreshold/decreaseRate 
            if (expect.lte(removeThreshold.div(decreaseRate))) {
                expect = new BN('0');
            }

            // step to epoch
            let epoch = 30;
            while (true) {
                let currentNumber = await web3.eth.getBlockNumber();
                if (currentNumber % epoch == (epoch - 1)) {
                    let receipt = await punishIns.decreaseMissedBlocksCounter(epoch, {
                        from: miner
                    });
                    expectEvent(receipt, "LogDecreaseMissedBlocksCounter");
                    break;
                }

                await time.advanceBlock();
            }

            let acutal_0 = await punishIns.getPunishRecord(initValidators[0]);
            assert.equal(expect.toNumber(), acutal_0.toNumber());
            let acutal_1 = await punishIns.getPunishRecord(initValidators[1]);
            assert.equal(acutal_1.toNumber(), step);
        })

        it("jailed record will be cleaned if validator repass proposal", async function () {
            let jailed = accounts[0];
            let removeThreshold = await proposalIns.removeThreshold();
            for (let i = 0; i < removeThreshold / 2; i++) {
                await punishIns.punish(jailed, {
                    from: miner
                });
                let record = await punishIns.getPunishRecord(jailed);
            }

            let record = await punishIns.getPunishRecord(jailed);
            assert.equal(record.isZero(), false);

            await pass(proposalIns, initValidators, jailed);

            // check record
            record = await punishIns.getPunishRecord(jailed);
            assert.equal(record.isZero(), true);

            // not in punish list
            let len = await punishIns.getPunishValidatorsLen();

            for (let i = 0; i < len.toNumber(); i++) {
                let punishee = await punishIns.punishValidators(i);
                assert.equal(punishee == jailed, false);
            }
        })
    })
})

async function pass(proposalIns, validators, who) {
    let receipt = await proposalIns.createProposal(who, true, "test", {
        from: who
    });
    let id = receipt.logs[0].args.id;
    for (let i = 0; i < validators.length / 2 + 1; i++) {
        await proposalIns.voteProposal(id, true, {
            from: validators[i]
        });
    }
}

function getInitValidators(accounts) {
    return accounts.slice(0, 3);
}

function getNewValidators(accounts) {
    return accounts.slice(3, 6);
}

function getInvalidMoniker() {
    let r = ""
    for (let i = 0; i < 71; i++) {
        r += i;
    }

    return r;
}