// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

contract Params {
    /**
     * @dev Indicates whether the contract has been initialized.
     */
    bool public initialized ;

    uint256 public epoch;
    uint256 public constant CONSENSUS_MAX_VALIDATORS = 21;
    uint256[50] private __gap;

    
    // Production version: constants use SCREAMING_SNAKE_CASE
    // System contracts (POSA addresses)
    /**
     * @dev Address of the Validators contract.
     */
    address
        public constant VALIDATOR_ADDR = 0x000000000000000000000000000000000000F010;
    /**
     * @dev Address of the Punish contract.
     */
    address
        public constant PUNISH_ADDR = 0x000000000000000000000000000000000000F011;
    /**
     * @dev Address of the Proposal contract.
     */
    address
        public constant PROPOSAL_ADDR = 0x000000000000000000000000000000000000F012;
    /**
     * @dev Address of the Staking contract.
     */
    address
        public constant STAKING_ADDR = 0x000000000000000000000000000000000000F013;

    

    modifier onlyMiner() {
        _onlyMiner();
        _;
    }

    modifier onlyNotInitialized() {
        _onlyNotInitialized();
        _;
    }

    modifier onlyInitialized() {
        _onlyInitialized();
        _;
    }

    modifier onlyPunishContract() {
        _onlyPunishContract();
        _;
    }

    modifier onlyBlockEpoch(uint256 epochParam) {
        _onlyBlockEpoch(epochParam);
        _;
    }

    modifier onlyNotEpoch() {
        _onlyNotEpoch();
        _;
    }

    modifier onlyValidatorsContract() {
        _onlyValidatorsContract();
        _;
    }

    modifier onlyProposalContract() {
        _onlyProposalContract();
        _;
    }

    modifier onlyStakingContract() {
        _onlyStakingContract();
        _;
    }

    modifier onlyPunishOrValidatorsContract() {
        _onlyPunishOrValidatorsContract();
        _;
    }

    function _onlyMiner() internal view {
        
        require(
            // For production, strictly use block.coinbase
            msg.sender == block.coinbase,
            "Miner only"
        );
        
    }

    function _onlyNotInitialized() internal view {
        require(!initialized, "Already initialized");
    }

    function _onlyInitialized() internal view {
        require(initialized, "Not initialized yet");
    }

    function _onlyPunishContract() internal view {
        
        require(msg.sender == PUNISH_ADDR, "Punish contract only");
        
    }

    function _onlyBlockEpoch(uint256 epochParam) internal view {
        // Check if block.number is divisible by epoch, no need for >= check since block.number starts at 0
        require(epoch > 0, "Epoch not set");
        require(epochParam == epoch, "Epoch mismatch");
        require(block.number % epoch == 0, "Block epoch only");
    }

    function _onlyNotEpoch() internal view {
        require(epoch > 0, "Epoch not set");
        require(block.number % epoch != 0, "Epoch block forbidden");
    }

    function _onlyValidatorsContract() internal view {
        require(
            
            msg.sender == VALIDATOR_ADDR,
            
            "Validators contract only"
        );
    }

    function _onlyProposalContract() internal view {
        require(
            
            msg.sender == PROPOSAL_ADDR,
            
            "Proposal contract only"
        );
    }

    function _onlyStakingContract() internal view {
        require(
            
            msg.sender == STAKING_ADDR,
            
            "Staking contract only"
        );
    }

    function _onlyPunishOrValidatorsContract() internal view {
        require(
            
            msg.sender == PUNISH_ADDR || msg.sender == VALIDATOR_ADDR,
            
            "Only punish or validators contract can call this function"
        );
    }

    function _initializeEpoch(uint256 epoch_) internal {
        require(epoch_ > 0, "Epoch must be positive");
        require(epoch == 0, "Epoch already set");
        epoch = epoch_;
    }

    
}
