// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Staking} from "../contracts/Staking.sol";

contract StakingSingleValidatorBranchTest is BaseSetup {
    address internal onlyValidator;

    function setUp() public {
        onlyValidator = makeAddr("staking-single-validator");
        address[] memory initVals = new address[](1);
        initVals[0] = onlyValidator;
        deploySystem(initVals);
    }

    function testSlashValidatorReturnsZeroWhenLastEffectiveValidatorAtMinimumStake() public {
        vm.prank(PUNISH);
        (uint256 actualSlash, uint256 actualReward) =
            Staking(STAKING).slashValidator(onlyValidator, 1 ether, address(0xBEEF), 0, address(0xdead));

        assertEq(actualSlash, 0);
        assertEq(actualReward, 0);
    }
}
