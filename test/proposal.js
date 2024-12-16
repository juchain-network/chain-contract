const MockValidators = artifacts.require("MockValidators");
const Proposal = artifacts.require("Proposal");

const {
    constants,
    expectRevert,
    expectEvent,
    time
} = require('@openzeppelin/test-helpers');
const {
    assert
} = require('chai');

// Test content:
// 1. initialize can only call once
//
// 1. anyone can create a proposal
// 2. one can't create a proposal to propose a already passed user
// 3. detail info can't too long
//
// 1. only validator can vote for a proposal
// 2. validator can only vote once for a proposal
// 3. validator can't vote for a expired proposal
// 4. len(validators)/2+1 vote agree, the proposal will pass
// 5. len(validators)/2+1 vote reject, the proposal will reject

contract("Proposal test", function (accounts) {
    // set validators
    let vals = [];
    for (let i = 0; i < 5; i++) {
        vals.push(accounts[i]);
    }
    var proposalIns;
    var mockVal;

    before(async function () {
        proposalIns = await Proposal.new();
        mockVal = await MockValidators.new(vals, proposalIns.address);
        for (let i = 0; i < vals.length; i++) {
            let exist = await mockVal.isActiveValidator(vals[i]);
            assert.equal(exist, true, "initialize validator failed");
        }

        await proposalIns.setContracts(mockVal.address, constants.ZERO_ADDRESS, constants.ZERO_ADDRESS);
        await proposalIns.initialize(vals);
    });

    it("Init can only call once", async function () {
        await expectRevert(proposalIns.initialize([]), "Already initialized");
    })

    describe("Create proposal", async function () {
        let candidate = accounts[6];
        it('anyone can create proposal', async function () {
            for (let i = 0; i < accounts.length && i < 10; i++) {
                let receipt = await proposalIns.createProposal(candidate, true, "", {
                    from: accounts[i]
                });

                expectEvent(receipt, 'LogCreateProposal', {
                    proposer: accounts[i],
                    dst: candidate,
                });
            }
        })

        it('can"t try to add already exist dst', async function () {
            await expectRevert(proposalIns.createProposal(accounts[4], true, ""), 'Cant"t add a already exist dst or Cant"t remove a not passed dst');
        })

        it('can"t try to remove not exist dst', async function () {
            await expectRevert(proposalIns.createProposal(accounts[6], false, ""), 'Cant"t add a already exist dst or Cant"t remove a not passed dst');
        })

        it("details info can't too long", async function () {
            await expectRevert(proposalIns.createProposal(candidate, true, getInvalidDetails()), "Details too long");
        })
    })

    describe("Vote for proposal(add,pass)", async function () {
        let candidate = accounts[6];
        let proposer = accounts[7];
        let id;

        it("normal vote for a proposal(3 true/2 false)", async function () {
            let receipt = await proposalIns.createProposal(candidate, true, "test", {
                from: proposer
            });
            id = receipt.logs[0].args.id;

            for (let i = 0; i < 3; i++) {
                let receipt = await proposalIns.voteProposal(id, true, {
                    from: accounts[i]
                });
                expectEvent(receipt, 'LogVote', {
                    id: id,
                    voter: accounts[i],
                    auth: true
                });

                if (i == 2) {
                    expectEvent(receipt, 'LogPassProposal', {
                        id: id
                    });
                    // check candidate is pass or not
                    let isPass = await proposalIns.pass(candidate);
                    assert.equal(true, isPass);
                }
            }

            receipt = await proposalIns.voteProposal(id, false, {
                from: accounts[3]
            });
            expectEvent(receipt, 'LogVote', {
                id: id,
                voter: accounts[3],
                auth: false
            });
        })

        it("only validator can vote for a proposal", async function () {
            await expectRevert(proposalIns.voteProposal(id, false, {
                from: accounts[6]
            }), "Validator only");
        })

        it("validator can only vote for a proposal once", async function () {
            await expectRevert(proposalIns.voteProposal(id, false, {
                from: accounts[1]
            }), "You can't vote for a proposal twice");
        })

        it("validator can't vote for proposal if it is expired", async function () {
            let step = await proposalIns.proposalLastingPeriod();
            await time.increase(step);
            await expectRevert(proposalIns.voteProposal(id, false, {
                from: accounts[4]
            }), "Proposal expired");
        })

        it("Validate candidate's info", async function () {
            // check proposal info
            let proposalInfo = await proposalIns.proposals(id);
            assert.equal(proposalInfo.proposer, proposer);
            assert.equal(proposalInfo.dst, candidate);
            let resultInfo = await proposalIns.results(id);
            assert.equal(resultInfo.agree.toNumber(), 3);
            assert.equal(resultInfo.reject.toNumber(), 1);
            assert.equal(resultInfo.resultExist, true);
            // ensure candidate is passed
            let pass = await proposalIns.pass(candidate);
            assert.equal(pass, true);
        })
    })
    describe("Vote for proposal(remove,pass)", async function () {
        let candidate = accounts[6];
        let proposer = accounts[7];
        let id;

        it("normal vote for a proposal(3 true/2 false)", async function () {
            let receipt = await proposalIns.createProposal(candidate, false, "test", {
                from: proposer
            });
            id = receipt.logs[0].args.id;

            for (let i = 0; i < 3; i++) {
                let receipt = await proposalIns.voteProposal(id, true, {
                    from: accounts[i]
                });
                expectEvent(receipt, 'LogVote', {
                    id: id,
                    voter: accounts[i],
                    auth: true
                });

                if (i == 2) {
                    expectEvent(receipt, 'LogPassProposal', {
                        id: id
                    });
                    // check candidate is pass or not
                    let isPass = await proposalIns.pass(candidate);
                    assert.equal(false, isPass);
                }
            }

            receipt = await proposalIns.voteProposal(id, false, {
                from: accounts[3]
            });
            expectEvent(receipt, 'LogVote', {
                id: id,
                voter: accounts[3],
                auth: false
            });
        })
    })

    describe("Vote for proposal(reject)", async function () {
        let candidate = accounts[8];
        let proposer = accounts[9];
        let id;

        it("normal vote(2 agree, 3 reject)", async function () {
            let receipt = await proposalIns.createProposal(candidate, true, "test", {
                from: proposer
            });
            id = receipt.logs[0].args.id;

            for (let i = 0; i < 2; i++) {
                let receipt = await proposalIns.voteProposal(id, true, {
                    from: accounts[i]
                });
                expectEvent(receipt, 'LogVote', {
                    id: id,
                    voter: accounts[i],
                    auth: true
                });

                if (i == 2) {
                    expectEvent(receipt, 'LogPassProposal', {
                        id: id
                    });
                }
            }

            for (let i = 2; i < 5; i++) {
                let receipt = await proposalIns.voteProposal(id, false, {
                    from: accounts[i]
                });
                expectEvent(receipt, 'LogVote', {
                    id: id,
                    voter: accounts[i],
                    auth: false
                });

                if (i == 4) {
                    expectEvent(receipt, 'LogRejectProposal', {
                        id: id
                    });
                }
            }
        })
    })

    describe("Create/Vote config update proposal", async function () {
        let proposer = accounts[7];
        let id;

        it("normal vote for a proposal(4 true/2 false)", async function () {
            list = [{
                    cid: 0,
                    value: 100
                },
                {
                    cid: 1,
                    value: 200
                },
                {
                    cid: 2,
                    value: 300
                },
                {
                    cid: 3,
                    value: 400
                },
                {
                    cid: 4,
                    value: 500
                },
                {
                    cid: 5,
                    value: 600
                },
                {
                    cid: 6,
                    value: "0x854bcf3915c629f63eb83e3c922dfa591151f29c",
                },
            ]

            for (let k = 0; k < list.length; k++) {
                let receipt = await proposalIns.createUpdateConfigProposal(list[k].cid, list[k].value.toString(), {
                    from: proposer
                });
                id = receipt.logs[0].args.id;

                for (let i = 0; i < 4; i++) {
                    let receipt = await proposalIns.voteProposal(id, true, {
                        from: accounts[i]
                    });

                    expectEvent(receipt, 'LogVote', {
                        id: id,
                        voter: accounts[i],
                        auth: true
                    });

                    if (i == 2) {
                        expectEvent(receipt, 'LogPassProposal', {
                            id: id
                        });
                    }
                }

                // check the config is update
                let cid = list[k].cid;
                let expect = list[k].value;
                let actual;
                if (cid == 0) {
                    actual = await proposalIns.proposalLastingPeriod();
                } else if (cid == 1) {
                    actual = await proposalIns.punishThreshold();
                } else if (cid == 2) {
                    actual = await proposalIns.removeThreshold();
                } else if (cid == 3) {
                    actual = await proposalIns.decreaseRate();
                } else if (cid == 4) {
                    actual = await proposalIns.withdrawProfitPeriod();
                } else if (cid == 5) {
                    actual = await proposalIns.increasePeriod();
                } else if (cid == 6) {
                    actual = await proposalIns.receiverAddr();
                }
                assert.equal(actual.toString().toLowerCase(), expect.toString().toLowerCase());
            }
        })
    })

    describe("Set val unpass", async function () {
        let candidate = accounts[1];

        it("only validator can set val unpass", async function () {
            await expectRevert(
                proposalIns.setUnpassed(candidate),
                "Validators contract only"
            )
        })
        it("validator contract can set val unpass", async function () {
            let before = await proposalIns.pass(candidate);
            assert.equal(before, true);

            await mockVal.setUnpassed(candidate);

            let after = await proposalIns.pass(candidate);
            assert.equal(after, false);
        })
    })
});

function getInvalidDetails() {
    var res = ""
    for (let i = 0; i < 3005; i++) {
        res += i;
    }

    return res;
}