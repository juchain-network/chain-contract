// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/Params.sol";

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
        assertEq(params.ValidatorContractAddr(), 0x000000000000000000000000000000000000f000);
        assertEq(params.PunishContractAddr(), 0x000000000000000000000000000000000000F001);
        assertEq(params.ProposalAddr(), 0x000000000000000000000000000000000000F002);
        assertEq(params.StakingContractAddr(), 0x000000000000000000000000000000000000F003);
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
        vm.expectRevert("Not init yet");
        params.callOnlyInitialized();
        
        // Initialize and test success
        params.initializeForTest();
        assertTrue(params.callOnlyInitialized());
    }
    
    function testOnlyPunishContractModifier() public {
        vm.prank(params.PunishContractAddr());
        assertTrue(params.callOnlyPunishContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Punish contract only");
        params.callOnlyPunishContract();
    }
    
    function testOnlyValidatorsContractModifier() public {
        vm.prank(params.ValidatorContractAddr());
        assertTrue(params.callOnlyValidatorsContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Validators contract only");
        params.callOnlyValidatorsContract();
    }
    
    function testOnlyProposalContractModifier() public {
        vm.prank(params.ProposalAddr());
        assertTrue(params.callOnlyProposalContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Proposal contract only");
        params.callOnlyProposalContract();
    }
    
    function testOnlyStakingContractModifier() public {
        vm.prank(params.StakingContractAddr());
        assertTrue(params.callOnlyStakingContract());
        
        vm.prank(address(0x123));
        vm.expectRevert("Staking contract only");
        params.callOnlyStakingContract();
    }
    
    function testOnlyBlockEpochModifier() public {
        uint256 epoch = 10;
        
        // Set block number to multiple of epoch
        vm.roll(20); // 20 % 10 == 0
        assertTrue(params.callOnlyBlockEpoch(epoch));
        
        // Set block number to non-multiple of epoch
        vm.roll(25); // 25 % 10 != 0
        vm.expectRevert("Block epoch only");
        params.callOnlyBlockEpoch(epoch);
    }
}
