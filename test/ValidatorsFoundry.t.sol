// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";

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
        miner = v1; // simulate coinbase
        vm.coinbase(miner);
    }

    function testDistributeBlockRewardEqually() public {
        // send 1 ether from coinbase
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

    // 1 ether / 3 with integer division, remainder to last non-jailed
    uint256 per = uint256(1 ether) / uint256(3);
    uint256 rem = uint256(1 ether) - per * 2;
    require(a1 == per, "v1 share");
    require(a2 == per, "v2 share");
    require(a3 == rem, "v3 share with remainder");
    }

    function testWithdrawProfitsAfterPeriod() public {
        // configure withdrawProfitPeriod small via proposal
        Proposal p = Proposal(PROPOSAL);
        bytes32 id;
        vm.warp(2_000_000);
        id = keccak256(abi.encodePacked(address(this), uint256(4), uint256(2), block.timestamp));
        p.createUpdateConfigProposal(4, 2);
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);

        // distribute some reward
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
}
