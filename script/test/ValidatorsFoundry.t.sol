// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Staking} from "../../contracts/Staking.sol";

contract ValidatorsFoundryTest is BaseSetup {

    address miner;
    address v1; address v2; address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        
        // Note: deploySystem() already registers genesis validators via initializeWithValidators()
        // So we don't need to register them again here
        
        miner = v1; // simulate coinbase
        vm.coinbase(miner);
        // Give miner enough ETH for testing
        vm.deal(miner, 100 ether);
    }
    

    function testDistributeBlockRewardEqually() public {
        // send 1 ether from coinbase (v1 is the miner)
        vm.startPrank(miner);
        (bool ok, ) = address(Validators(VALIDATORS)).call{value: 1 ether}(
            abi.encodeWithSelector(Validators.distributeBlockReward.selector)
        );
        vm.stopPrank();
        require(ok, "distribute failed");

        // read validator profits (aacIncoming)
        ( , , uint256 a1,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        ( , , uint256 a2,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        ( , , uint256 a3,,) = Validators(VALIDATORS).getValidatorInfo(v3);

        // New logic: reward goes directly to the block producer (v1)
        require(a1 == 1 ether, "v1 (block producer) should get full reward");
        require(a2 == 0, "v2 should get no reward");
        require(a3 == 0, "v3 should get no reward");
    }

    function testWithdrawProfitsAfterPeriod() public {
        // configure withdrawProfitPeriod small via proposal
        Proposal p = Proposal(PROPOSAL);
        bytes32 id;
        vm.warp(2_000_000);
        
        // Create proposal from v1 (active validator) instead of address(this)
        vm.prank(v1);
        id = p.createUpdateConfigProposal(4, 2);
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);

        // distribute some reward (v1 is the miner, so v1 gets the reward)
        vm.startPrank(miner);
        (bool ok, ) = address(Validators(VALIDATORS)).call{value: 9 ether}(
            abi.encodeWithSelector(Validators.distributeBlockReward.selector)
        );
        vm.stopPrank();
        require(ok, "distribute failed");

        // advance blocks to satisfy withdrawProfitPeriod
        vm.roll(block.number + 3);

    // fee addr defaults to validator addr, must call as fee receiver
    uint256 balBefore = miner.balance;
    vm.prank(miner);
    Validators(VALIDATORS).withdrawProfits(miner);
    uint256 balAfter = miner.balance;
        require(balAfter > balBefore, "profits withdrawn");
    }

    function testAddProfitsToActiveValidatorsWithRemainder() public {
        // Jail v1 (miner), so rewards will be distributed to other validators
        // Use Staking contract's jailValidator function
        vm.prank(VALIDATORS); // Validators contract is allowed to call jailValidator
        Staking(STAKING).jailValidator(v1, 100);

        // Send 3 wei as reward - this won't be divisible by 2 (v2 and v3 are active)
        vm.startPrank(miner);
        (bool ok, ) = address(Validators(VALIDATORS)).call{value: 3 wei}(
            abi.encodeWithSelector(Validators.distributeBlockReward.selector)
        );
        vm.stopPrank();
        require(ok, "distribute failed");

        // Check validator profits (aacIncoming)
        // v1 should get 0 (jailed), v2 should get 2 wei, v3 should get 1 wei
        ( , , uint256 a1,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        ( , , uint256 a2,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        ( , , uint256 a3,,) = Validators(VALIDATORS).getValidatorInfo(v3);

        require(a1 == 0, "v1 (jailed) should get no reward");
        require(a2 == 2 wei, "v2 should get 2 wei");
        require(a3 == 1 wei, "v3 should get 1 wei");
    }
}
