// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Punish} from "../../contracts/Punish.sol";

// Complete reward distribution testing
contract RewardFoundryTest is BaseSetup {

    address v1; address v2; address v3;
    address miner;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2"); 
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        miner = v1; // Simulate v1 as the block producer (miner)
        vm.coinbase(miner);
        // Give miner enough ETH for testing
        vm.deal(miner, 100 ether);
    }

    function testRewardEquallyDistributedNoStake() public {
        // Corresponds to "reward should go directly to block producer if not jailed"
        // New logic: reward goes directly to the block producer (miner = v1)
        uint256 reward = 1 ether;
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: reward}();

        // Check block producer (v1) receives full reward
        (, , uint256 v1Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1Incoming == reward, "block producer should get full reward");
        
        // Check other validators do not receive rewards
        (, , uint256 v2Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        require(v2Incoming == 0, "v2 should get no reward");
        require(v3Incoming == 0, "v3 should get no reward");
    }

    function testRemoveValidatorReward() public {
        // Corresponds to "remove validator's reward"
        uint256 threshold = Proposal(PROPOSAL).punishThreshold();
        
        // Give v1 some rewards (v1 is miner, i.e. block producer, gets full reward)
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        (, , uint256 toRemoveBefore,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(toRemoveBefore == 1 ether, "v1 should have received the full reward");
        
        // Punish v1 until it is jailed
        for (uint i = 0; i < threshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // The reward of the punished validator should be removed and distributed to other validators
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1After == 0, "punished validator should have no reward");
        
        uint256 inPlan = toRemoveBefore / 2;
        uint256 added = inPlan * 2;
        uint256 remain = toRemoveBefore - added;
        
        require(v2After - v2Before == inPlan, "v2 should get equal share");
        require(v3After - v3Before == inPlan + remain, "v3 should get equal share plus remainder");
    }

    function testJailedValidatorCantGetReward() public {
        // Corresponds to "jailed block producer's reward is distributed to other validators"
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        
        // Punish v1 until it is jailed (v1 is miner, i.e. block producer)
        for (uint i = 0; i < removeThreshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // Record other validators' reward before state
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // Distribute new reward (v1 is jailed, so reward goes to other validators)
        uint256 reward = 1 ether;
        vm.prank(miner); // miner is still v1, but v1 is now jailed
        Validators(VALIDATORS).distributeBlockReward{value: reward}();
        
        // Check jailed block producer (v1) does not get reward
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1After == 0, "jailed block producer should not get reward");
        
        // Check reward is distributed to other active validators (v2 and v3)
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        uint256 inPlan = reward / 2; //  Only 2 active validators (v2 and v3)
        uint256 remain = reward - inPlan * 2;
        
        require(v2After - v2Before == inPlan, "v2 should get equal share");
        require(v3After - v3Before == inPlan + remain, "v3 should get equal share plus remainder");
    }

    function testJailedValidatorCantGetPunishProfits() public {
        // Corresponds to "jailed validator can't get profits of punish"
        uint256 threshold = Proposal(PROPOSAL).punishThreshold();
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        
        // Jail v1
        for (uint i = 0; i < removeThreshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // Give v2 some rewards (v1 is jailed, so reward goes to other validators)
        // Note: miner is still v1, but v1 is jailed, so reward goes to v2 and v3
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        // Since v1 is jailed, reward should be distributed to v2 and v3
        (, , uint256 v1Before,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1Before == 0, "jailed v1 should not get reward");
        require(v2Before > 0 || v3Before > 0, "v2 or v3 should have received reward");
        
        // Punish v2 until it is jailed (v1 is still jailed)
        for (uint i = 0; i < threshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v2);
            vm.roll(block.number + 1);
        }
        
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1After == v1Before, "jailed validator should not benefit from punish"); // Jailed validator should not benefit from punish
        require(v2After == 0, "punished validator should lose reward"); // Punished validator should lose reward
        require(v3After - v3Before == v2Before, "v3 should get all v2's reward"); // v3 should get all v2's reward
    }
}
