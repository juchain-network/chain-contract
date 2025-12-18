// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";

contract ValidatorsAdditionalTest is BaseSetup {
    address v1; address v2; address v3;
    
    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
    }

    // Test isValidatorActive function
    function testIsValidatorActive() public {
        // Test for active validator
        bool isActive = Validators(VALIDATORS).isValidatorActive(v1);
        assertTrue(isActive, "v1 should be active validator");

        // Test for non-existent validator
        address nonExistent = makeAddr("nonExistent");
        bool isNonActive = Validators(VALIDATORS).isValidatorActive(nonExistent);
        assertFalse(isNonActive, "nonExistent should not be active validator");
    }

    // Test validateDescription function
    function testValidateDescription() public pure {
        // Test valid description
        bool isValid = Validators(VALIDATORS).validateDescription(
            "Test Validator", // moniker - max 70 chars
            "identity",       // identity - max 3000 chars
            "https://example.com", // website - max 140 chars
            "test@example.com",    // email - max 140 chars
            "Test details"         // details - max 280 chars
        );
        assertTrue(isValid, "Valid description should pass validation");
    }

    // Test validateDescription with invalid inputs
    function testValidateDescriptionInvalidInputs() public {
        // Test invalid moniker (too long)
        vm.expectRevert("Invalid moniker length");
        Validators(VALIDATORS).validateDescription(
            "aTooLongMonikerThatExceedsSeventyCharactersWhichIsTheMaximumAllowedAndShouldTriggerARevert", // > 70 chars
            "identity",
            "https://example.com",
            "test@example.com",
            "Test details"
        );

        // Test invalid identity (too long)
        vm.expectRevert("Invalid identity length");
        // Generate a string longer than 3000 characters
        string memory longIdentity = new string(3001);
        bytes memory longIdentityBytes = bytes(longIdentity);
        for (uint i = 0; i < longIdentityBytes.length; i++) {
            // casting to 'bytes1' is safe because "x" is a single character
            // forge-lint: disable-next-line(unsafe-typecast)
            longIdentityBytes[i] = bytes1("x");
        }
        longIdentity = string(longIdentityBytes);
        Validators(VALIDATORS).validateDescription(
            "Test Validator",
            longIdentity,
            "https://example.com",
            "test@example.com",
            "Test details"
        );

        // Test invalid website (too long)
        vm.expectRevert("Invalid website length");
        Validators(VALIDATORS).validateDescription(
            "Test Validator",
            "identity",
            "https://aVeryLongWebsiteUrlThatExceedsOneHundredAndFortyCharactersWhichIsBeyondTheLimitSetInTheSmartContractCodeAndShouldTriggerARevertWithAppropriateErrorMessage.com",
            "test@example.com",
            "Test details"
        );

        // Test invalid email (too long)
        vm.expectRevert("Invalid email length");
        Validators(VALIDATORS).validateDescription(
            "Test Validator",
            "identity",
            "https://example.com",
            "aVeryLongEmailAddressThatExceedsOneHundredAndFortyCharactersInTotalWhichGoesBeyondTheLimitDefinedInTheSmartContractCodeAndShouldResultInARevertWithProperErrorMessagingForValidationPurpose@example.com",
            "test@example.com"
        );

        // Test invalid details (too long)
        vm.expectRevert("Invalid details length");
        Validators(VALIDATORS).validateDescription(
            "Test Validator",
            "identity",
            "https://example.com",
            "test@example.com",
            "aVeryLongDetailsStringThatExceedsTwoHundredAndEightyCharactersWhichIsTheUpperBoundSetInTheSmartContractSourceCodeAndShouldCauseTheFunctionToRevertWithASpecificErrorIndicatingTheNatureOfTheProblemRatherThanSimplyFailingSilentlyWithoutAnyInformativeFeedbackAndTriggersTheRevertCorrectly"
        );
    }

    // Test getActiveValidators function
    function testGetActiveValidators() public view {
        address[] memory activeValidators = Validators(VALIDATORS).getActiveValidators();
        assertEq(activeValidators.length, 3, "Should have 3 active validators");
        assertEq(activeValidators[0], v1, "First validator should be v1");
        assertEq(activeValidators[1], v2, "Second validator should be v2");
        assertEq(activeValidators[2], v3, "Third validator should be v3");
    }

    // Test getHighestValidators function
    function testGetHighestValidators() public view {
        address[] memory highestValidators = Validators(VALIDATORS).getHighestValidators();
        assertEq(highestValidators.length, 3, "Should have 3 highest validators");
    }

    // Test isTopValidator function
    function testIsTopValidator() public {
        bool isTop = Validators(VALIDATORS).isTopValidator(v1);
        assertTrue(isTop, "v1 should be a top validator");
        
        address nonTop = makeAddr("nonTop");
        bool isNonTop = Validators(VALIDATORS).isTopValidator(nonTop);
        assertFalse(isNonTop, "nonTop should not be a top validator");
    }

    // Test getValidatorDescription function
    function testGetValidatorDescription() public {
        // v1 is already a validator from setUp, just edit its description
        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(
            payable(v1),
            "Test Validator",
            "identity",
            "https://example.com",
            "test@example.com",
            "Test details"
        );

        // Now test getting the description
        (string memory moniker, string memory identity, string memory website, string memory email, string memory details) = 
            Validators(VALIDATORS).getValidatorDescription(v1);
        
        assertEq(moniker, "Test Validator", "Moniker should match");
        assertEq(identity, "identity", "Identity should match");
        assertEq(website, "https://example.com", "Website should match");
        assertEq(email, "test@example.com", "Email should match");
        assertEq(details, "Test details", "Details should match");
    }

    // Test isActiveValidator function
    function testIsActiveValidator() public {
        bool isActive = Validators(VALIDATORS).isActiveValidator(v1);
        assertTrue(isActive, "v1 should be active");

        address inactive = makeAddr("inactive");
        bool isInactive = Validators(VALIDATORS).isActiveValidator(inactive);
        assertFalse(isInactive, "inactive should not be active");
    }

    // Test getActiveValidatorCount function
    function testGetActiveValidatorCount() public view {
        uint256 count = Validators(VALIDATORS).getActiveValidatorCount();
        assertEq(count, 3, "Should have 3 active validators");
    }

    // Test modifier onlyNotUpdated
    function testOnlyNotUpdatedModifier() public {
        // Need to move to a future block since setup might have used current block
        vm.roll(block.number + 1);
        
        // Set v1 as coinbase/miner
        vm.coinbase(v1);
        
        // First call should succeed
        address[] memory newSet = new address[](3);
        newSet[0] = v1;
        newSet[1] = v2;
        newSet[2] = v3;
        
        vm.prank(v1); // v1 is the miner
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, block.number);
        
        // Second call in same block should silently return (not revert)
        vm.prank(v1);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, block.number);
        
        // Test passed if no revert occurred
    }

    // Test getActiveValidatorsWithStakes function
    function testGetActiveValidatorsWithStakes() public view {
        (address[] memory validators, uint256[] memory stakes) = Validators(VALIDATORS).getActiveValidatorsWithStakes();
        assertEq(validators.length, 3, "Should have 3 validators");
        assertEq(stakes.length, 3, "Should have 3 stakes");
        // All genesis validators have 1000 stake by default (as set in BaseSetup)
        assertEq(stakes[0], 100000000000000000000000, "First validator should have 100000000000000000000000 stake");
        assertEq(stakes[1], 100000000000000000000000, "Second validator should have 100000000000000000000000 stake");
        assertEq(stakes[2], 100000000000000000000000, "Third validator should have 100000000000000000000000 stake");
    }
}