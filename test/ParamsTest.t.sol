// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test} from "forge-std/Test.sol";
import {Params} from "../contracts/Params.sol";

contract TestableParams is Params {
    // Helper functions for testing modifiers - not auto-run tests
    function callOnlyMiner() external view onlyMiner returns (bool) {
        return true;
    }
    
    function callOnlyNotInitialized() external view onlyNotInitialized returns (bool) {
        return true;
    }
    
    function callOnlyInitialized() external view onlyInitialized returns (bool) {
        return true;
    }
    
    function callOnlyPunishContract() external view onlyPunishContract returns (bool) {
        return true;
    }
    
    function callOnlyValidatorsContract() external view onlyValidatorsContract returns (bool) {
        return true;
    }
    
    function callOnlyProposalContract() external view onlyProposalContract returns (bool) {
        return true;
    }
    
    function callOnlyStakingContract() external view onlyStakingContract returns (bool) {
        return true;
    }
    
    function callOnlyBlockEpoch(uint256 epoch) external view onlyBlockEpoch(epoch) returns (bool) {
        return true;
    }

    function setEpochForTest(uint256 epoch_) external {
        _initializeEpoch(epoch_);
    }
    
    function initializeForTest() external {
        initialized = true;
    }
}

contract ParamsTest is Test {
    TestableParams public params;
    
    function setUp() public {
        params = new TestableParams();
    }
    
    function testConstants() public view {
        assertEq(params.VALIDATOR_ADDR(), 0x000000000000000000000000000000000000F010);
        assertEq(params.PUNISH_ADDR(), 0x000000000000000000000000000000000000F011);
        assertEq(params.PROPOSAL_ADDR(), 0x000000000000000000000000000000000000F012);
        assertEq(params.STAKING_ADDR(), 0x000000000000000000000000000000000000F013);
    }
    
    function testInitialState() public view {
        assertFalse(params.initialized());
    }
    
    function testOnlyMinerModifier() public {
        // Set block coinbase to test address
        address miner = address(0x123);
        vm.coinbase(miner);
        
        vm.prank(miner);
        assertTrue(params.callOnlyMiner());
        
        // Test failure with non-miner
        vm.prank(address(0x456));
        vm.expectRevert("Miner only");
        params.callOnlyMiner();
    }
    
    function testOnlyNotInitializedModifier() public {
        // Should work when not initialized
        assertTrue(params.callOnlyNotInitialized());
        
        // Initialize and test failure
        params.initializeForTest();
        vm.expectRevert("Already initialized");
        params.callOnlyNotInitialized();
    }
    
    function testOnlyInitializedModifier() public {
        // Should fail when not initialized
        vm.expectRevert("Not initialized yet");
        params.callOnlyInitialized();
        
        // Initialize and test success
        params.initializeForTest();
        assertTrue(params.callOnlyInitialized());
    }
    
    function testOnlyPunishContractModifier() public {
        vm.prank(params.PUNISH_ADDR());
        assertTrue(params.callOnlyPunishContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Punish contract only");
        params.callOnlyPunishContract();
    }
    
    function testOnlyValidatorsContractModifier() public {
        vm.prank(params.VALIDATOR_ADDR());
        assertTrue(params.callOnlyValidatorsContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Validators contract only");
        params.callOnlyValidatorsContract();
    }
    
    function testOnlyProposalContractModifier() public {
        vm.prank(params.PROPOSAL_ADDR());
        assertTrue(params.callOnlyProposalContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Proposal contract only");
        params.callOnlyProposalContract();
    }
    
    function testOnlyStakingContractModifier() public {
        vm.prank(params.STAKING_ADDR());
        assertTrue(params.callOnlyStakingContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Staking contract only");
        params.callOnlyStakingContract();
    }
    
    function testOnlyBlockEpochModifier() public {
        uint256 epoch = 10;
        params.setEpochForTest(epoch);
        
        // Set block number to multiple of epoch
        vm.roll(20); // 20 % 10 == 0
        assertTrue(params.callOnlyBlockEpoch(epoch));
        
        // Set block number to non-multiple of epoch
        vm.roll(25); // 25 % 10 != 0
        vm.expectRevert("Block epoch only");
        params.callOnlyBlockEpoch(epoch);
    }
}
