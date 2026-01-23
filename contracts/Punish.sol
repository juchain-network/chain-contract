// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from "./Params.sol";
import {IStaking} from "./IStaking.sol";
import {IValidators} from "./IValidators.sol";
import {IProposal} from "./IProposal.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract Punish is Params, ReentrancyGuard {
    struct PunishRecord {
        uint256 missedBlocksCounter;
        uint256 index;
        bool exist;
    }

    IValidators validators;
    IProposal proposal;
    IStaking staking;

    mapping(address => PunishRecord) punishRecords;
    address[] public punishValidators;

    mapping(uint256 => bool) punished;
    mapping(uint256 => bool) decreased;
    mapping(address => bool) pendingRemove;
    mapping(address => bool) pendingRemoveIncoming;
    address[] public pendingValidators;
    mapping(address => uint256) pendingIndex;
    mapping(uint256 => mapping(address => bool)) doubleSigned;

    uint256 public constant DOUBLE_SIGN_WINDOW = 86400;
    uint256 private constant SECP256K1N_HALF =
        0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0;

    event LogDecreaseMissedBlocksCounter();
    event LogPunishValidator(address indexed val, uint256 time);
    event LogDoubleSignPunish(
        address indexed val,
        address indexed reporter,
        uint256 height,
        uint256 slashAmount,
        uint256 rewardAmount,
        uint256 time
    );

    modifier onlyNotPunished() {
        _onlyNotPunished();
        _;
    }

    modifier onlyNotDecreased() {
        _onlyNotDecreased();
        _;
    }

    function _onlyNotPunished() internal view {
        require(!punished[block.number], "Already punished");
    }

    function _onlyNotDecreased() internal view {
        require(!decreased[block.number], "Already decreased");
    }

    /**
     * @dev Initializes the Punish contract with required dependencies.
     * @param validators_ Address of the Validators contract.
     * @param proposal_ Address of the Proposal contract.
     * @param staking_ Address of the Staking contract.
     */
    function initialize(
        address validators_,
        address proposal_,
        address staking_
    ) external onlyNotInitialized {
        require(validators_ != address(0), "Invalid validators address");
        require(proposal_ != address(0), "Invalid proposal address");
        
        validators = IValidators(validators_);
        proposal = IProposal(proposal_);
        staking = IStaking(staking_);
        _initializeEpoch(proposal.epoch());

        initialized = true;
    }

    /**
     * @dev Punishes a validator for missing blocks.
     * @param val Address of the validator to punish.
     * @notice Only the miner can call this function.
     * @notice Punishment is applied based on missed blocks threshold.
     */
    function punish(address val) external onlyMiner onlyInitialized onlyNotPunished nonReentrant {
        punished[block.number] = true;
        require(epoch > 0, "Epoch not set");
        bool isEpochBlock = block.number % epoch == 0;

        if (!isEpochBlock) {
            if (pendingRemove[val]) {
                pendingRemove[val] = false;
                pendingRemoveIncoming[val] = false;
                punishRecords[val].missedBlocksCounter = 0;
                staking.jailValidator(val, proposal.validatorUnjailPeriod());
                validators.removeValidator(val);
                emit LogPunishValidator(val, block.timestamp);
                return;
            }
            if (pendingRemoveIncoming[val]) {
                pendingRemoveIncoming[val] = false;
                validators.removeValidatorIncoming(val);
                emit LogPunishValidator(val, block.timestamp);
                return;
            }
        }

        if (!punishRecords[val].exist) {
            punishRecords[val].index = punishValidators.length;
            punishValidators.push(val);
            punishRecords[val].exist = true;
        }
        punishRecords[val].missedBlocksCounter++;

        if (isEpochBlock) {
            if (punishRecords[val].missedBlocksCounter % proposal.removeThreshold() == 0) {
                pendingRemove[val] = true;
                pendingRemoveIncoming[val] = false;
                _enqueuePending(val);
            } else if (punishRecords[val].missedBlocksCounter % proposal.punishThreshold() == 0) {
                pendingRemoveIncoming[val] = true;
                _enqueuePending(val);
            }
            emit LogPunishValidator(val, block.timestamp);
            return;
        }

        if (punishRecords[val].missedBlocksCounter % proposal.removeThreshold() == 0) {
            // reset validator's missed blocks counter
            punishRecords[val].missedBlocksCounter = 0;
            // jail validator first (sets isJailed in Staking contract)
            staking.jailValidator(val, proposal.validatorUnjailPeriod());
            // then remove validator (which will check isJailed status)
            validators.removeValidator(val);
        } else if (punishRecords[val].missedBlocksCounter % proposal.punishThreshold() == 0) {
            validators.removeValidatorIncoming(val);
        }

        emit LogPunishValidator(val, block.timestamp);
    }

    function executePending(uint256 limit) external onlyInitialized onlyNotEpoch nonReentrant {
        if (limit == 0) {
            return;
        }

        uint256 processed = 0;
        while (processed < limit && pendingValidators.length > 0) {
            address val = pendingValidators[pendingValidators.length - 1];
            pendingValidators.pop();
            pendingIndex[val] = 0;

            if (pendingRemove[val]) {
                pendingRemove[val] = false;
                pendingRemoveIncoming[val] = false;
                punishRecords[val].missedBlocksCounter = 0;
                staking.jailValidator(val, proposal.validatorUnjailPeriod());
                validators.removeValidator(val);
            } else if (pendingRemoveIncoming[val]) {
                pendingRemoveIncoming[val] = false;
                validators.removeValidatorIncoming(val);
            }

            processed++;
        }
    }

    /**
     * @dev Submit double-sign evidence for punishment.
     * @param header1 RLP-encoded block header (variant 1)
     * @param header2 RLP-encoded block header (variant 2)
     * @notice Anyone can submit evidence, only allowed on non-epoch blocks.
     */
    function submitDoubleSignEvidence(bytes calldata header1, bytes calldata header2)
        external
        onlyInitialized
        onlyNotEpoch
        nonReentrant
    {
        (uint256 number1, address signer1, bytes32 hash1) = _recoverSignerAndNumber(header1);
        (uint256 number2, address signer2, bytes32 hash2) = _recoverSignerAndNumber(header2);

        require(number1 == number2, "Height mismatch");
        require(hash1 != hash2, "Same header");
        require(signer1 == signer2, "Different signer");
        require(validators.isValidatorExist(signer1), "Signer not exist");
        require(!doubleSigned[number1][signer1], "Already punished");
        require(block.number >= number1, "Future block");
        require(block.number - number1 <= DOUBLE_SIGN_WINDOW, "Evidence expired");

        doubleSigned[number1][signer1] = true;

        staking.jailValidator(signer1, proposal.validatorUnjailPeriod());
        (uint256 actualSlash, uint256 actualReward) = staking.slashValidator(
            signer1,
            proposal.doubleSignSlashAmount(),
            msg.sender,
            proposal.doubleSignRewardAmount(),
            proposal.burnAddress()
        );

        emit LogDoubleSignPunish(signer1, msg.sender, number1, actualSlash, actualReward, block.timestamp);
    }

    /**
     * @dev Decreases the missed blocks counter for all validators at the end of an epoch.
     * @param epoch The epoch number for which to decrease the counter.
     * @notice Only the miner can call this function.
     * @notice This function is called once per epoch to reduce punishment counters.
     */
    function decreaseMissedBlocksCounter(uint256 epoch)
        external
        onlyMiner
        onlyNotDecreased
        onlyInitialized
        onlyBlockEpoch(epoch)
    {
        decreased[block.number] = true;
        if (punishValidators.length == 0) {
            return;
        }

        // Cache external results outside the loop
        uint256 removeThreshold = proposal.removeThreshold();
        uint256 decreaseRate = proposal.decreaseRate();
        uint256 decreaseAmount = removeThreshold / decreaseRate;

        uint256 punishValidatorsLength = punishValidators.length;
        for (uint256 i = 0; i < punishValidatorsLength; i++) {
            address validator = punishValidators[i];
            if (punishRecords[validator].missedBlocksCounter > decreaseAmount) {
                punishRecords[validator].missedBlocksCounter -= decreaseAmount;
            } else {
                punishRecords[validator].missedBlocksCounter = 0;
            }
        }

        emit LogDecreaseMissedBlocksCounter();
    }

    function _enqueuePending(address val) private {
        if (pendingIndex[val] != 0) {
            return;
        }
        pendingValidators.push(val);
        pendingIndex[val] = pendingValidators.length;
    }

    /**
     * @dev Cleans the punishment record for a validator.
     * @param val Address of the validator whose record to clean.
     * @return bool Returns true if the operation was successful.
     * @notice This function is called when a validator restakes.
     */
    function cleanPunishRecord(address val) public onlyInitialized onlyValidatorsContract returns (bool) {
        if (punishRecords[val].missedBlocksCounter != 0) {
            punishRecords[val].missedBlocksCounter = 0;
        }

        // remove it out of array if exist
        if (punishRecords[val].exist && punishValidators.length > 0) {
            if (punishRecords[val].index != punishValidators.length - 1) {
                address uval = punishValidators[punishValidators.length - 1];
                punishValidators[punishRecords[val].index] = uval;

                punishRecords[uval].index = punishRecords[val].index;
            }
            punishValidators.pop();
            punishRecords[val].index = 0;
            punishRecords[val].exist = false;
        }

        return true;
    }

    /**
     * @dev Gets the number of validators currently being punished.
     * @return uint256 The number of validators in the punishment list.
     */
    function getPunishValidatorsLen() public view returns (uint256) {
        return punishValidators.length;
    }

    /**
     * @dev Gets the punishment record for a specific validator.
     * @param val Address of the validator to check.
     * @return uint256 The number of missed blocks for the validator.
     */
    function getPunishRecord(address val) public view returns (uint256) {
        return punishRecords[val].missedBlocksCounter;
    }

    struct RLPItem {
        uint256 offset;
        uint256 len;
    }

    function _recoverSignerAndNumber(bytes calldata header) private pure returns (uint256 number, address signer, bytes32 headerHash) {
        bytes memory headerMem = header;
        headerHash = keccak256(headerMem);

        (RLPItem[] memory items, uint256 listOffset, uint256 listLen) = _decodeListItems(headerMem);
        require(items.length >= 15 && items.length <= 20, "Invalid header length");
        require(listOffset + listLen == headerMem.length, "Invalid header length");

        bytes memory sig;
        bytes memory encodedExtra;
        address coinbase;
        (number, coinbase, sig, encodedExtra) = _extractHeaderEvidence(headerMem, items);
        bytes32 sealHash = _computeSealHash(headerMem, items, encodedExtra);
        signer = _recoverSigner(sealHash, sig);
        require(signer == coinbase, "Signer != coinbase");
    }

    function _extractHeaderEvidence(bytes memory headerMem, RLPItem[] memory items)
        private
        pure
        returns (uint256 number, address coinbase, bytes memory sig, bytes memory encodedExtra)
    {
        number = _decodeUint(headerMem, items[8]);
        coinbase = _decodeAddress(headerMem, items[2]);
        (uint256 extraOffset, uint256 extraLen) = _payloadOffsetLen(headerMem, items[12]);
        require(extraLen >= 97, "Invalid extra");
        sig = _slice(headerMem, extraOffset + extraLen - 65, 65);
        bytes memory extraTrimmed = _slice(headerMem, extraOffset, extraLen - 65);
        encodedExtra = _encodeBytes(extraTrimmed);
    }

    function _computeSealHash(bytes memory headerMem, RLPItem[] memory items, bytes memory encodedExtra)
        private
        pure
        returns (bytes32)
    {
        uint256 payloadLen = 0;
        for (uint256 i = 0; i < items.length; i++) {
            if (i == 12) {
                payloadLen += encodedExtra.length;
            } else {
                payloadLen += items[i].len;
            }
        }

        bytes memory listPrefix = _encodeListPrefix(payloadLen);
        bytes memory out = new bytes(listPrefix.length + payloadLen);
        uint256 dest = _copyBytes(out, 0, listPrefix);

        for (uint256 i = 0; i < items.length; i++) {
            if (i == 12) {
                dest = _copyBytes(out, dest, encodedExtra);
            } else {
                dest = _copyBytesSlice(out, dest, headerMem, items[i].offset, items[i].len);
            }
        }

        return keccak256(out);
    }

    function _recoverSigner(bytes32 sealHash, bytes memory sig) private pure returns (address) {
        require(sig.length == 65, "Invalid signature length");
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
        if (v < 27) {
            v += 27;
        }
        require(v == 27 || v == 28, "Invalid signature v");
        require(uint256(s) <= SECP256K1N_HALF, "Invalid signature s");
        address signer = ecrecover(sealHash, v, r, s);
        require(signer != address(0), "Invalid signature");
        return signer;
    }

    function _decodeListItems(bytes memory data)
        private
        pure
        returns (RLPItem[] memory items, uint256 listOffset, uint256 listLen)
    {
        (uint256 itemLen, uint256 payloadOffset, uint256 payloadLen, bool isList) = _decodeItem(data, 0);
        require(isList, "Not a list");
        listOffset = payloadOffset;
        listLen = payloadLen;
        require(itemLen == data.length, "List length mismatch");

        uint256 offset = listOffset;
        uint256 end = listOffset + listLen;
        uint256 count = 0;
        while (offset < end) {
            (uint256 len,, ,) = _decodeItem(data, offset);
            offset += len;
            count++;
        }
        require(offset == end, "Invalid list");

        items = new RLPItem[](count);
        offset = listOffset;
        for (uint256 i = 0; i < count; i++) {
            (uint256 len,, ,) = _decodeItem(data, offset);
            items[i] = RLPItem(offset, len);
            offset += len;
        }
    }

    function _payloadOffsetLen(bytes memory data, RLPItem memory item) private pure returns (uint256 offset, uint256 len) {
        (, offset, len,) = _decodeItem(data, item.offset);
    }

    function _decodeUint(bytes memory data, RLPItem memory item) private pure returns (uint256 value) {
        (uint256 offset, uint256 len) = _payloadOffsetLen(data, item);
        require(len <= 32, "RLP uint too large");
        if (len == 0) {
            return 0;
        }
        for (uint256 i = 0; i < len; i++) {
            value = (value << 8) | uint8(data[offset + i]);
        }
    }

    function _decodeAddress(bytes memory data, RLPItem memory item) private pure returns (address) {
        (uint256 offset, uint256 len) = _payloadOffsetLen(data, item);
        require(len == 20, "Invalid address length");
        uint256 value = 0;
        for (uint256 i = 0; i < len; i++) {
            value = (value << 8) | uint8(data[offset + i]);
        }
        return address(uint160(value));
    }

    function _decodeItem(bytes memory data, uint256 offset)
        private
        pure
        returns (uint256 itemLen, uint256 payloadOffset, uint256 payloadLen, bool isList)
    {
        require(offset < data.length, "RLP offset out of bounds");
        uint8 byte0 = uint8(data[offset]);
        if (byte0 <= 0x7f) {
            return (1, offset, 1, false);
        } else if (byte0 <= 0xb7) {
            payloadLen = byte0 - 0x80;
            payloadOffset = offset + 1;
            itemLen = 1 + payloadLen;
            require(payloadOffset + payloadLen <= data.length, "RLP item out of bounds");
            return (itemLen, payloadOffset, payloadLen, false);
        } else if (byte0 <= 0xbf) {
            uint256 lenOfLen = byte0 - 0xb7;
            payloadLen = _readBigEndian(data, offset + 1, lenOfLen);
            payloadOffset = offset + 1 + lenOfLen;
            itemLen = 1 + lenOfLen + payloadLen;
            require(payloadOffset + payloadLen <= data.length, "RLP item out of bounds");
            return (itemLen, payloadOffset, payloadLen, false);
        } else if (byte0 <= 0xf7) {
            payloadLen = byte0 - 0xc0;
            payloadOffset = offset + 1;
            itemLen = 1 + payloadLen;
            require(payloadOffset + payloadLen <= data.length, "RLP item out of bounds");
            return (itemLen, payloadOffset, payloadLen, true);
        } else {
            uint256 lenOfLen = byte0 - 0xf7;
            payloadLen = _readBigEndian(data, offset + 1, lenOfLen);
            payloadOffset = offset + 1 + lenOfLen;
            itemLen = 1 + lenOfLen + payloadLen;
            require(payloadOffset + payloadLen <= data.length, "RLP item out of bounds");
            return (itemLen, payloadOffset, payloadLen, true);
        }
    }

    function _readBigEndian(bytes memory data, uint256 offset, uint256 len) private pure returns (uint256 value) {
        require(len > 0 && len <= 32, "Invalid length");
        require(offset + len <= data.length, "RLP length out of bounds");
        for (uint256 i = 0; i < len; i++) {
            value = (value << 8) | uint8(data[offset + i]);
        }
    }

    function _encodeBytes(bytes memory data) private pure returns (bytes memory out) {
        uint256 len = data.length;
        if (len == 1 && uint8(data[0]) <= 0x7f) {
            return data;
        }
        if (len <= 55) {
            out = new bytes(1 + len);
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

    function _encodeListPrefix(uint256 len) private pure returns (bytes memory out) {
        if (len <= 55) {
            out = new bytes(1);
            out[0] = bytes1(uint8(0xc0 + len));
            return out;
        }
        bytes memory lenBytes = _toBinary(len);
        out = new bytes(1 + lenBytes.length);
        out[0] = bytes1(uint8(0xf7 + lenBytes.length));
        _copyBytesSlice(out, 1, lenBytes, 0, lenBytes.length);
    }

    function _toBinary(uint256 value) private pure returns (bytes memory out) {
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
            out[i - 1] = bytes1(uint8(value));
            value >>= 8;
        }
    }

    function _slice(bytes memory data, uint256 start, uint256 len) private pure returns (bytes memory out) {
        require(start + len <= data.length, "Slice out of bounds");
        out = new bytes(len);
        for (uint256 i = 0; i < len; i++) {
            out[i] = data[start + i];
        }
    }

    function _copyBytes(bytes memory dest, uint256 destOffset, bytes memory src) private pure returns (uint256) {
        return _copyBytesSlice(dest, destOffset, src, 0, src.length);
    }

    function _copyBytesSlice(
        bytes memory dest,
        uint256 destOffset,
        bytes memory src,
        uint256 srcOffset,
        uint256 len
    ) private pure returns (uint256) {
        require(destOffset + len <= dest.length, "Copy overflow");
        require(srcOffset + len <= src.length, "Copy out of bounds");
        for (uint256 i = 0; i < len; i++) {
            dest[destOffset + i] = src[srcOffset + i];
        }
        return destOffset + len;
    }
}
