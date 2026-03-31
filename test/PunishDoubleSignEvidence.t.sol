// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test} from "forge-std/Test.sol";
import {BaseSetup} from "./BaseSetup.t.sol";
import {Punish} from "../contracts/Punish.sol";
import {Staking} from "../contracts/Staking.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

contract PunishDoubleSignEvidenceTest is Test, BaseSetup {
    uint256 private constant SIGNER_KEY = 0xA11CE;

    address internal signer;
    bytes[] internal goHeaders1;
    bytes[] internal goHeaders2;
    address[] internal goSigners;
    uint256[] internal goHeights;

    function setUp() public virtual {
        _loadGoVectors();
        signer = vm.addr(SIGNER_KEY);
        address[] memory initVals = _buildInitValidators();
        deploySystem(initVals);

        // Fund staking contract to pay rewards/burn
        vm.deal(STAKING, 200000 ether);

        // Ensure non-epoch block
        vm.roll(1000);
    }

    function testSubmitDoubleSignEvidenceValid() public virtual {
        uint256 height = block.number - 1;

        bytes memory header1 = _buildSignedHeader(height, signer, bytes32(uint256(1)), bytes8(uint64(1)), SIGNER_KEY);
        bytes memory header2 = _buildSignedHeader(height, signer, bytes32(uint256(2)), bytes8(uint64(2)), SIGNER_KEY);

        (uint256 selfStakeBefore,,,,,,,,,) = Staking(STAKING).getValidatorInfo(signer);
        address reporter = vm.addr(0xB0B);
        uint256 reporterBalanceBefore = reporter.balance;

        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);

        (uint256 selfStakeAfter,,,,,,,,,) = Staking(STAKING).getValidatorInfo(signer);
        uint256 slashAmount = Proposal(PROPOSAL).doubleSignSlashAmount();
        uint256 rewardAmount = Proposal(PROPOSAL).doubleSignRewardAmount();

        uint256 actualSlash = slashAmount > selfStakeBefore ? selfStakeBefore : slashAmount;
        uint256 actualReward = rewardAmount > actualSlash ? actualSlash : rewardAmount;

        assertEq(selfStakeAfter, selfStakeBefore - actualSlash, "selfStake should be slashed");
        assertEq(reporter.balance, reporterBalanceBefore + actualReward, "reporter should be rewarded");

        vm.expectRevert("Already punished");
        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testSubmitDoubleSignEvidenceSignerMismatchCoinbase() public {
        uint256 height = block.number - 1;
        address fakeCoinbase = address(0xBEEF);

        bytes memory header1 =
            _buildSignedHeader(height, fakeCoinbase, bytes32(uint256(1)), bytes8(uint64(1)), SIGNER_KEY);
        bytes memory header2 =
            _buildSignedHeader(height, fakeCoinbase, bytes32(uint256(2)), bytes8(uint64(2)), SIGNER_KEY);

        vm.expectRevert("Signer != coinbase");
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testSubmitDoubleSignEvidenceInvalidHeaderLength() public {
        bytes memory header = _buildHeaderWithItemCount(21);

        vm.expectRevert("Invalid header length");
        Punish(PUNISH).submitDoubleSignEvidence(header, header);
    }

    function testSubmitDoubleSignEvidenceZeroNumberLongExtra() public virtual {
        vm.roll(1);
        bytes memory extraTrimmed = new bytes(80);
        bytes memory header1 =
            _buildSignedHeaderWithExtra(0, signer, bytes32(uint256(1)), bytes8(uint64(1)), SIGNER_KEY, extraTrimmed);
        bytes memory header2 =
            _buildSignedHeaderWithExtra(0, signer, bytes32(uint256(2)), bytes8(uint64(2)), SIGNER_KEY, extraTrimmed);

        address reporter = vm.addr(0xB0D);
        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testSubmitDoubleSignEvidenceGoVectors() public {
        if (goHeaders1.length == 0) {
            return;
        }
        require(goHeaders1.length == goHeaders2.length, "vector length mismatch");
        require(goHeaders1.length == goSigners.length, "vector length mismatch");
        require(goHeaders1.length == goHeights.length, "vector length mismatch");

        address reporter = vm.addr(0xB0C);
        uint256 slashAmount = Proposal(PROPOSAL).doubleSignSlashAmount();
        uint256 rewardAmount = Proposal(PROPOSAL).doubleSignRewardAmount();

        for (uint256 i = 0; i < goHeaders1.length; i++) {
            uint256 height = goHeights[i];
            address expectedSigner = goSigners[i];

            uint256 rollTo = height + 1;
            if (rollTo % TEST_EPOCH == 0) {
                rollTo += 1;
            }
            vm.roll(rollTo);

            (uint256 selfStakeBefore,,,,,,,,,) = Staking(STAKING).getValidatorInfo(expectedSigner);
            uint256 reporterBalanceBefore = reporter.balance;

            vm.prank(reporter);
            Punish(PUNISH).submitDoubleSignEvidence(goHeaders1[i], goHeaders2[i]);

            (uint256 selfStakeAfter,,,,,,,,,) = Staking(STAKING).getValidatorInfo(expectedSigner);
            uint256 actualSlash = slashAmount > selfStakeBefore ? selfStakeBefore : slashAmount;
            uint256 actualReward = rewardAmount > actualSlash ? actualSlash : rewardAmount;

            assertEq(selfStakeAfter, selfStakeBefore - actualSlash, "selfStake should be slashed");
            assertEq(reporter.balance, reporterBalanceBefore + actualReward, "reporter should be rewarded");
        }
    }

    function _loadGoVectors() internal {
        goSigners.push(0x0CA50Bf78F617bCB2D4a5ab64A70fB2E6000F340);
        goHeights.push(100);
        goHeaders1.push(
            hex"f90256a09d58f64f0b6f2b94f5e3e8b3d4b98e3b93a6f91c17e3f4d64b9d9df4e3fd1e01a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940ca50bf78f617bcb2d4a5ab64a70fb2e6000f340a01111111111111111111111111111111111111111111111111111111111111111a0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026483e4e1c080846553f100b861111111111111111111111111111111111111111111111111111111111111111115843d331c9acf4a65a6ac8f005c85168e08830ca2eb555b43dbeba94dc32648290d0d457aa422ac0e7b7fa6f79c13723bb71962f2df9eaba81b53dae96e707600a00000000000000000000000000000000000000000000000000000000000000000880000000000000000"
        );
        goHeaders2.push(
            hex"f90256a09d58f64f0b6f2b94f5e3e8b3d4b98e3b93a6f91c17e3f4d64b9d9df4e3fd1e01a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940ca50bf78f617bcb2d4a5ab64a70fb2e6000f340a02222222222222222222222222222222222222222222222222222222222222222a0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026483e4e1c080846553f100b86111111111111111111111111111111111111111111111111111111111111111117d246a04d654bf61a442fda3dc8207bcdc3b670d7eeb8889f696b0a92db487ee1b52e1c84b400a5a43fdb9f6a0365b7e857f3d1ff286941ec546ce2b9124feea01a00000000000000000000000000000000000000000000000000000000000000000880000000000000000"
        );

        goSigners.push(0x0CA50Bf78F617bCB2D4a5ab64A70fB2E6000F340);
        goHeights.push(101);
        goHeaders1.push(
            hex"f9029fa09d58f64f0b6f2b94f5e3e8b3d4b98e3b93a6f91c17e3f4d64b9d9df4e3fd1e01a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940ca50bf78f617bcb2d4a5ab64a70fb2e6000f340a01111111111111111111111111111111111111111111111111111111111111111a0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026583e4e1c080846553f100b8611111111111111111111111111111111111111111111111111111111111111111af5e1475c349ac57b51664b0890beb1e085e9e6e22bc796f99b6028c6c4d2c166b222bb7b424ae8e538d113239b29239b10622ff5238db90109a6db75c93e00f00a00000000000000000000000000000000000000000000000000000000000000000880000000000000000843b9aca00a033333333333333333333333333333333333333333333333333333333333333330102a04444444444444444444444444444444444444444444444444444444444444444"
        );
        goHeaders2.push(
            hex"f9029fa09d58f64f0b6f2b94f5e3e8b3d4b98e3b93a6f91c17e3f4d64b9d9df4e3fd1e01a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940ca50bf78f617bcb2d4a5ab64a70fb2e6000f340a02222222222222222222222222222222222222222222222222222222222222222a0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026583e4e1c080846553f100b8611111111111111111111111111111111111111111111111111111111111111111507dcfd149fa8964cb49a4b035850102b1147ad298da46b82181688e5de5257004eb5441b5522218cb15278cbabf0e1f2fddbe4bbeff81b458e636e2f4c405fe01a00000000000000000000000000000000000000000000000000000000000000000880000000000000000843b9aca00a033333333333333333333333333333333333333333333333333333333333333330102a04444444444444444444444444444444444444444444444444444444444444444"
        );
    }

    function _buildInitValidators() internal view returns (address[] memory initVals) {
        uint256 maxLen = goSigners.length + 1;
        address[] memory temp = new address[](maxLen);
        uint256 count = 0;

        for (uint256 i = 0; i < goSigners.length; i++) {
            address validator = goSigners[i];
            require(validator != address(0), "Invalid signer");
            bool exists = false;
            for (uint256 j = 0; j < count; j++) {
                if (temp[j] == validator) {
                    exists = true;
                    break;
                }
            }
            if (!exists) {
                temp[count] = validator;
                count++;
            }
        }

        bool hasSigner = false;
        for (uint256 j = 0; j < count; j++) {
            if (temp[j] == signer) {
                hasSigner = true;
                break;
            }
        }
        if (!hasSigner) {
            temp[count] = signer;
            count++;
        }

        initVals = new address[](count);
        for (uint256 k = 0; k < count; k++) {
            initVals[k] = temp[k];
        }
    }

    function _buildSignedHeader(uint256 number, address coinbase, bytes32 mixDigest, bytes8 nonce, uint256 signerKey)
        internal
        pure
        returns (bytes memory)
    {
        bytes memory extraTrimmed = new bytes(32);
        return _buildSignedHeaderWithExtra(number, coinbase, mixDigest, nonce, signerKey, extraTrimmed);
    }

    function _buildSignedHeaderWithExtra(
        uint256 number,
        address coinbase,
        bytes32 mixDigest,
        bytes8 nonce,
        uint256 signerKey,
        bytes memory extraTrimmed
    ) internal pure returns (bytes memory) {
        bytes memory rlpForSeal = _buildHeaderRlp(number, coinbase, mixDigest, nonce, extraTrimmed);
        bytes32 sealHash = keccak256(rlpForSeal);

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerKey, sealHash);
        bytes memory sig = abi.encodePacked(r, s, v);
        bytes memory extraWithSig = bytes.concat(extraTrimmed, sig);

        return _buildHeaderRlp(number, coinbase, mixDigest, nonce, extraWithSig);
    }

    function _buildHeaderRlp(uint256 number, address coinbase, bytes32 mixDigest, bytes8 nonce, bytes memory extra)
        internal
        pure
        returns (bytes memory)
    {
        bytes[] memory items = new bytes[](15);
        items[0] = _encodeBytes(abi.encodePacked(bytes32(uint256(1)))); // ParentHash
        items[1] = _encodeBytes(abi.encodePacked(bytes32(uint256(2)))); // UncleHash
        items[2] = _encodeBytes(abi.encodePacked(coinbase)); // Coinbase
        items[3] = _encodeBytes(abi.encodePacked(bytes32(uint256(3)))); // Root
        items[4] = _encodeBytes(abi.encodePacked(bytes32(uint256(4)))); // TxHash
        items[5] = _encodeBytes(abi.encodePacked(bytes32(uint256(5)))); // ReceiptHash
        items[6] = _encodeBytes(abi.encodePacked(bytes32(uint256(6)))); // Bloom
        items[7] = _encodeUint(1); // Difficulty
        items[8] = _encodeUint(number); // Number
        items[9] = _encodeUint(30_000_000); // GasLimit
        items[10] = _encodeUint(0); // GasUsed
        items[11] = _encodeUint(12345); // Time
        items[12] = _encodeBytes(extra); // Extra
        items[13] = _encodeBytes(abi.encodePacked(mixDigest)); // MixDigest
        items[14] = _encodeBytes(abi.encodePacked(nonce)); // Nonce

        return _encodeList(items);
    }

    function _buildHeaderWithItemCount(uint256 count) internal pure returns (bytes memory) {
        bytes[] memory items = new bytes[](count);
        for (uint256 i = 0; i < count; i++) {
            items[i] = _encodeBytes("");
        }
        return _encodeList(items);
    }

    function _encodeList(bytes[] memory items) internal pure returns (bytes memory out) {
        uint256 payloadLen = 0;
        for (uint256 i = 0; i < items.length; i++) {
            payloadLen += items[i].length;
        }
        bytes memory prefix = _encodeListPrefix(payloadLen);
        out = new bytes(prefix.length + payloadLen);
        uint256 dest = _copyBytes(out, 0, prefix);
        for (uint256 i = 0; i < items.length; i++) {
            dest = _copyBytes(out, dest, items[i]);
        }
    }

    function _encodeUint(uint256 value) internal pure returns (bytes memory) {
        if (value == 0) {
            return hex"80";
        }
        bytes memory data = _toBinary(value);
        return _encodeBytes(data);
    }

    function _encodeBytes(bytes memory data) internal pure returns (bytes memory out) {
        uint256 len = data.length;
        if (len == 1 && uint8(data[0]) <= 0x7f) {
            return data;
        }
        if (len <= 55) {
            out = new bytes(1 + len);
            // forge-lint: disable-next-line(unsafe-typecast)
            out[0] = bytes1(uint8(0x80 + len));
            _copyBytesSlice(out, 1, data, 0, len);
            return out;
        }
        bytes memory lenBytes = _toBinary(len);
        out = new bytes(1 + lenBytes.length + len);
        out[0] = bytes1(uint8(0xb7 + lenBytes.length));
        _copyBytesSlice(out, 1, lenBytes, 0, lenBytes.length);
        _copyBytesSlice(out, 1 + lenBytes.length, data, 0, len);
        return out;
    }

    function _encodeListPrefix(uint256 len) internal pure returns (bytes memory out) {
        if (len <= 55) {
            out = new bytes(1);
            // forge-lint: disable-next-line(unsafe-typecast)
            out[0] = bytes1(uint8(0xc0 + len));
            return out;
        }
        bytes memory lenBytes = _toBinary(len);
        out = new bytes(1 + lenBytes.length);
        out[0] = bytes1(uint8(0xf7 + lenBytes.length));
        _copyBytesSlice(out, 1, lenBytes, 0, lenBytes.length);
    }

    function _toBinary(uint256 value) internal pure returns (bytes memory out) {
        if (value == 0) {
            return new bytes(0);
        }
        uint256 temp = value;
        uint256 len = 0;
        while (temp != 0) {
            len++;
            temp >>= 8;
        }
        out = new bytes(len);
        for (uint256 i = len; i > 0; i--) {
            // forge-lint: disable-next-line(unsafe-typecast)
            out[i - 1] = bytes1(uint8(value));
            value >>= 8;
        }
    }

    function _copyBytes(bytes memory dest, uint256 destOffset, bytes memory src) internal pure returns (uint256) {
        return _copyBytesSlice(dest, destOffset, src, 0, src.length);
    }

    function _copyBytesSlice(bytes memory dest, uint256 destOffset, bytes memory src, uint256 srcOffset, uint256 len)
        internal
        pure
        returns (uint256)
    {
        for (uint256 i = 0; i < len; i++) {
            dest[destOffset + i] = src[srcOffset + i];
        }
        return destOffset + len;
    }
}

contract PunishDoubleSignEvidenceSignerSeparationTest is PunishDoubleSignEvidenceTest {
    uint256 private constant OLD_SIGNER_KEY = 0xBEEF01;
    uint256 private constant NEW_SIGNER_KEY = 0xBEEF02;
    uint256 private constant OTHER_SIGNER_KEY_1 = 0xBEEF03;
    uint256 private constant OTHER_SIGNER_KEY_2 = 0xBEEF04;

    address private validator1;
    address private oldSigner;
    address private newSigner;

    function setUp() public override {
        validator1 = makeAddr("cold-validator-1");
        address validator2 = makeAddr("cold-validator-2");
        address validator3 = makeAddr("cold-validator-3");

        oldSigner = vm.addr(OLD_SIGNER_KEY);
        newSigner = vm.addr(NEW_SIGNER_KEY);
        signer = oldSigner;
        address signer2 = vm.addr(OTHER_SIGNER_KEY_1);
        address signer3 = vm.addr(OTHER_SIGNER_KEY_2);

        address[] memory initVals = new address[](3);
        initVals[0] = validator1;
        initVals[1] = validator2;
        initVals[2] = validator3;

        address[] memory initSigners = new address[](3);
        initSigners[0] = oldSigner;
        initSigners[1] = signer2;
        initSigners[2] = signer3;

        deploySystem(initVals, initSigners, 10);
        vm.deal(STAKING, 200000 ether);
        vm.roll(5);
    }

    function testDoubleSignEvidenceUsesHistoricalSignerAfterRotation() public {
        vm.prank(validator1);
        Validators(VALIDATORS).createOrEditValidator(payable(validator1), newSigner, "", "", "", "", "");

        bytes memory header1 = _buildSignedHeader(9, oldSigner, bytes32(uint256(1)), bytes8(uint64(1)), OLD_SIGNER_KEY);
        bytes memory header2 = _buildSignedHeader(9, oldSigner, bytes32(uint256(2)), bytes8(uint64(2)), OLD_SIGNER_KEY);

        vm.roll(11);

        (uint256 selfStakeBefore,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        uint256 slashAmount = Proposal(PROPOSAL).doubleSignSlashAmount();
        uint256 rewardAmount = Proposal(PROPOSAL).doubleSignRewardAmount();

        address reporter = vm.addr(0xCAFE01);
        uint256 reporterBalanceBefore = reporter.balance;

        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);

        (uint256 selfStakeAfter,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        uint256 actualSlash = slashAmount > selfStakeBefore ? selfStakeBefore : slashAmount;
        uint256 actualReward = rewardAmount > actualSlash ? actualSlash : rewardAmount;

        assertEq(selfStakeAfter, selfStakeBefore - actualSlash, "historical signer should slash validator cold stake");
        assertEq(reporter.balance, reporterBalanceBefore + actualReward, "reporter reward mismatch");
    }

    function testSubmitDoubleSignEvidenceValid() public override {
        uint256 height = block.number - 1;

        bytes memory header1 =
            _buildSignedHeader(height, oldSigner, bytes32(uint256(1)), bytes8(uint64(1)), OLD_SIGNER_KEY);
        bytes memory header2 =
            _buildSignedHeader(height, oldSigner, bytes32(uint256(2)), bytes8(uint64(2)), OLD_SIGNER_KEY);

        (uint256 selfStakeBefore,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        address reporter = vm.addr(0xB0B);
        uint256 reporterBalanceBefore = reporter.balance;

        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);

        (uint256 selfStakeAfter,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        uint256 slashAmount = Proposal(PROPOSAL).doubleSignSlashAmount();
        uint256 rewardAmount = Proposal(PROPOSAL).doubleSignRewardAmount();

        uint256 actualSlash = slashAmount > selfStakeBefore ? selfStakeBefore : slashAmount;
        uint256 actualReward = rewardAmount > actualSlash ? actualSlash : rewardAmount;

        assertEq(selfStakeAfter, selfStakeBefore - actualSlash, "selfStake should be slashed");
        assertEq(reporter.balance, reporterBalanceBefore + actualReward, "reporter should be rewarded");

        vm.expectRevert("Already punished");
        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testPendingSignerCannotBePunishedBeforeActivation() public {
        vm.prank(validator1);
        Validators(VALIDATORS).createOrEditValidator(payable(validator1), newSigner, "", "", "", "", "");

        bytes memory header1 = _buildSignedHeader(6, newSigner, bytes32(uint256(1)), bytes8(uint64(1)), NEW_SIGNER_KEY);
        bytes memory header2 = _buildSignedHeader(6, newSigner, bytes32(uint256(2)), bytes8(uint64(2)), NEW_SIGNER_KEY);

        vm.roll(7);
        vm.expectRevert("Signer not exist");
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testCheckpointBlockStillRejectsNextEpochSignerEvidence() public {
        vm.prank(validator1);
        Validators(VALIDATORS).createOrEditValidator(payable(validator1), newSigner, "", "", "", "", "");

        vm.roll(10);
        vm.coinbase(oldSigner);
        address[] memory newSet = Validators(VALIDATORS).getTopValidators();
        vm.prank(oldSigner);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, 10);

        bytes memory header1 = _buildSignedHeader(10, newSigner, bytes32(uint256(1)), bytes8(uint64(1)), NEW_SIGNER_KEY);
        bytes memory header2 = _buildSignedHeader(10, newSigner, bytes32(uint256(2)), bytes8(uint64(2)), NEW_SIGNER_KEY);

        vm.roll(11);
        vm.expectRevert("Signer not exist");
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }

    function testNextEpochSignerEvidenceSucceedsAfterActivationBeforeSync() public {
        vm.prank(validator1);
        Validators(VALIDATORS).createOrEditValidator(payable(validator1), newSigner, "", "", "", "", "");

        vm.roll(10);
        vm.coinbase(oldSigner);
        address[] memory newSet = Validators(VALIDATORS).getTopValidators();
        vm.prank(oldSigner);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, 10);

        bytes memory header1 = _buildSignedHeader(11, newSigner, bytes32(uint256(1)), bytes8(uint64(1)), NEW_SIGNER_KEY);
        bytes memory header2 = _buildSignedHeader(11, newSigner, bytes32(uint256(2)), bytes8(uint64(2)), NEW_SIGNER_KEY);

        vm.roll(12);

        (uint256 selfStakeBefore,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        uint256 slashAmount = Proposal(PROPOSAL).doubleSignSlashAmount();
        uint256 rewardAmount = Proposal(PROPOSAL).doubleSignRewardAmount();
        address reporter = vm.addr(0xCAFE02);
        uint256 reporterBalanceBefore = reporter.balance;

        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);

        (uint256 selfStakeAfter,,,,,,,,,) = Staking(STAKING).getValidatorInfo(validator1);
        uint256 actualSlash = slashAmount > selfStakeBefore ? selfStakeBefore : slashAmount;
        uint256 actualReward = rewardAmount > actualSlash ? actualSlash : rewardAmount;

        assertEq(selfStakeAfter, selfStakeBefore - actualSlash, "new signer should slash validator after activation");
        assertEq(reporter.balance, reporterBalanceBefore + actualReward, "reporter reward mismatch");
    }

    function testSubmitDoubleSignEvidenceZeroNumberLongExtra() public override {
        vm.roll(1);
        bytes memory extraTrimmed = new bytes(80);
        bytes memory header1 = _buildSignedHeaderWithExtra(
            0, oldSigner, bytes32(uint256(1)), bytes8(uint64(1)), OLD_SIGNER_KEY, extraTrimmed
        );
        bytes memory header2 = _buildSignedHeaderWithExtra(
            0, oldSigner, bytes32(uint256(2)), bytes8(uint64(2)), OLD_SIGNER_KEY, extraTrimmed
        );

        address reporter = vm.addr(0xB0D);
        vm.prank(reporter);
        Punish(PUNISH).submitDoubleSignEvidence(header1, header2);
    }
}
