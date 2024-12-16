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

contract("Validators test", function (accounts) {
    var valIns, proposalIns, punishIns, initValidators;
    var miner = accounts[0];
    var initLen;

    before(async function () {
        valIns = await Validators.new();
        proposalIns = await Proposal.new();
        punishIns = await Punish.new();

        initValidators = accounts.slice(0, 3);
        initLen = new BN(initValidators.length.toString())

        await proposalIns.setContracts(valIns.address, constants.ZERO_ADDRESS, constants.ZERO_ADDRESS);
        await valIns.setContracts(valIns.address, punishIns.address, proposalIns.address);
        await punishIns.setContracts(valIns.address, punishIns.address, proposalIns.address);

        await valIns.initialize(initValidators);
        await valIns.setMiner(miner);
        await proposalIns.initialize(initValidators);
        await punishIns.setMiner(miner);
        await punishIns.initialize();
    })

    describe("normal case(no validator jailed)", async function () {
        it("reward should be equally distributed to active validators if no stake", async function () {

            let reward = ether('1');
            await valIns.distributeBlockReward({
                from: miner,
                value: reward
            });
            let remain = reward.sub(reward.div(initLen).mul(initLen));

            for (let i = 0; i < initValidators.length; i++) {
                let info = await valIns.getValidatorInfo(initValidators[i]);
                let inPlan = reward.div(initLen);

                if (i == initValidators.length - 1) {
                    assert.equal(info[2].eq(inPlan.add(remain)), true);
                } else {
                    assert.equal(info[2].eq(inPlan), true);
                }
            }
        })
    })

    describe("punish reward should be distributed to others by stake percent", async function () {
        it("remove validator's reward", async function () {
            let punishee = initValidators[0];
            let info = await valIns.getValidatorInfo(punishee);
            let toRemove = info[2];

            let punishThreshold = await proposalIns.punishThreshold();
            let before = await getBefore(valIns, initValidators);

            for (let i = 0; i < punishThreshold.toNumber(); i++) {
                await punishIns.punish(punishee, {
                    from: miner
                });
            }

            // at this time, the profits of punishee will be removed to others.
            info = await valIns.getValidatorInfo(punishee);
            assert.equal(info[2], 0);

            let added = ether('0');
            for (let i = 1; i < initValidators.length; i++) {
                info = await valIns.getValidatorInfo(initValidators[i]);

                let inPlan = toRemove.div(initLen.sub(new BN('1')));
                added = added.add(inPlan);

                if (i == initValidators.length - 1) {
                    assert.equal(info[2].sub(before[i]).eq(inPlan.add(toRemove.sub(added))), true);
                } else {
                    assert.equal(info[2].sub(before[i]).eq(inPlan), true);
                }
            }
        })

        it("jailed validator can't get reward", async function () {
            let punishee = initValidators[0];
            let removeThreshold = await proposalIns.removeThreshold();
            for (let i = 0; i < removeThreshold.toNumber(); i++) {
                await punishIns.punish(punishee, {
                    from: miner
                });
            }

            // at this time, the profit should only sent to not jailed validators.
            let before = await getBefore(valIns, initValidators);
            let reward = ether('1');
            await valIns.distributeBlockReward({
                from: miner,
                value: reward
            });

            let added = ether('0');
            for (let i = 1; i < initValidators.length; i++) {
                let info = await valIns.getValidatorInfo(initValidators[i]);

                let infoPlan = reward.div(initLen.sub(new BN('1')));
                added = added.add(infoPlan);

                if (i == initValidators.length - 1) {
                    assert.equal(info[2].sub(before[i]).eq(infoPlan.add(reward.sub(added))), true);
                } else {
                    assert.equal(info[2].sub(before[i]).eq(infoPlan), true);
                }
            }
        })

        it("jailed validator can't get profits of punish", async function () {
            let punishThreshold = await proposalIns.punishThreshold();
            let punishee = initValidators[1];
            let before = await getBefore(valIns, initValidators);

            for (let i = 0; i < punishThreshold.toNumber(); i++) {
                await punishIns.punish(punishee, {
                    from: miner
                });
            }

            let info = await valIns.getValidatorInfo(initValidators[0]);
            assert.equal(info[2].eq(before[0]), true);

            info = await valIns.getValidatorInfo(initValidators[1]);
            assert.equal(info[2], 0);

            info = await valIns.getValidatorInfo(initValidators[2]);
            assert.equal(info[2].sub(before[2]).eq(before[1]), true);
        })
    })
})

async function getBefore(valIns, vals) {
    let before = [];
    for (let i = 0; i < vals.length; i++) {
        let info = await valIns.getValidatorInfo(vals[i]);
        before.push(info[2]);
    }

    return before
}