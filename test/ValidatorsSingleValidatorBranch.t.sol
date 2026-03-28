// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";

contract ValidatorsSingleValidatorBranchTest is BaseSetup {
    address internal onlyValidator;

    function setUp() public {
        onlyValidator = makeAddr("validators-single-validator");
        address[] memory initVals = new address[](1);
        initVals[0] = onlyValidator;
        deploySystem(initVals);
    }

    function testRemoveValidatorIncomingReturnsWhenSingleValidatorRemains() public {
        vm.prank(PUNISH);
        Validators(VALIDATORS).removeValidatorIncoming(onlyValidator);
    }
}
