// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

contract Params {
    bool public initialized ;

    
    // System contracts (POSA addresses)
    address
        public constant VALIDATOR_CONTRACT_ADDR = 0x000000000000000000000000000000000000F010;
    address
        public constant PUNISH_CONTRACT_ADDR = 0x000000000000000000000000000000000000F011;
    address
        public constant PROPOSAL_ADDR = 0x000000000000000000000000000000000000F012;
    address
        public constant STAKING_CONTRACT_ADDR = 0x000000000000000000000000000000000000F013;

    

    
    modifier onlyMiner() {
        require(
            msg.sender == block.coinbase,
            "Miner only"
        );
        _;
    }

    
    modifier onlyNotInitialized() {
        require(!initialized, "Already initialized");
        _;
    }

    modifier onlyInitialized() {
        require(initialized, "Not init yet");
        _;
    }

    modifier onlyPunishContract() {
        require(msg.sender == PUNISH_CONTRACT_ADDR, "Punish contract only");
        _;
    }

    modifier onlyBlockEpoch(uint256 epoch) {
        require(block.number % epoch == 0, "Block epoch only");
        _;
    }

    modifier onlyValidatorsContract() {
        require(
            msg.sender == VALIDATOR_CONTRACT_ADDR,
            "Validators contract only"
        );
        _;
    }

    modifier onlyProposalContract() {
        require(
            msg.sender == PROPOSAL_ADDR,
            "Proposal contract only"
        );
        _;
    }

    modifier onlyStakingContract() {
        require(
            msg.sender == STAKING_CONTRACT_ADDR,
            "Staking contract only"
        );
        _;
    }

    
}
