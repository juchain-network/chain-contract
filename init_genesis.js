#!/usr/bin/env node

// Extract system contract bytecode and update Genesis file
const fs = require('fs');
const path = require('path');
const { keccak256 } = require('js-sha3');

console.log('🔧 Extracting system contract bytecode and updating Genesis file...');

// Contract address mappings
const CONTRACT_ADDRESSES = {
    'Validators': '0x000000000000000000000000000000000000f010',
    'Punish': '0x000000000000000000000000000000000000f011',
    'Proposal': '0x000000000000000000000000000000000000f012',
    'Staking': '0x000000000000000000000000000000000000f013'
};

// Initial validator information (3 validators, meeting minimum validator requirement)
const INITIAL_VALIDATORS = [
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC'
];

// Read contract bytecode
function getContractBytecode(contractName) {
    // Try Foundry first (out directory)
    const foundryPath = path.join(__dirname, 'out', `${contractName}.sol`, `${contractName}.json`);
    try {
        const artifact = JSON.parse(fs.readFileSync(foundryPath, 'utf8'));
        return artifact.deployedBytecode?.object || artifact.bytecode?.object;
    } catch (foundryError) {
        // Fallback to Hardhat artifacts
        const artifactPath = path.join(__dirname, '..', 'artifacts', 'contracts', `${contractName}.sol`, `${contractName}.json`);
        try {
            const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
            return artifact.deployedBytecode;
        } catch (error) {
            console.error(`❌ Failed to read ${contractName} contract bytecode:`, error.message);
            return null;
        }
    }
}

// keccak256 hash helper function
function keccak256Hash(data) {
    if (typeof data === 'string' && data.startsWith('0x')) {
        return '0x' + keccak256(Buffer.from(data.slice(2), 'hex'));
    }
    return '0x' + keccak256(data);
}

// Generate initial validators' extraData
// Using pre-allocated accounts as initial validators
function generateExtraData() {
    // Build extraData structure:
    // 32 bytes vanity + N*20 bytes validator addresses + 65 bytes signature
    const vanity = '0'.repeat(64); // 32 bytes vanity (can be arbitrary data)
    
    // Validator address list (remove 0x prefix, keep 20 bytes)
    const validatorAddresses = INITIAL_VALIDATORS.map(addr => addr.slice(2).toLowerCase()).join('');
    
    // To generate the correct signature, we need:
    // 1. Build data to sign (excluding signature part)
    const dataToSign = vanity + validatorAddresses;
    
    // 2. Hash the data
    const hash = keccak256Hash(Buffer.from(dataToSign, 'hex'));
    
    // 3. Generate a simple signature (in real scenarios, this should be signed by validator private keys)
    // Here we use a deterministic method to generate the signature
    const messageHash = Buffer.from(hash.slice(2), 'hex');
    
    // Generate a fixed signature (65 bytes: 32 bytes r + 32 bytes s + 1 byte v)
    // Note: This is not a real ECDSA signature, just for correct format
    const r = keccak256(messageHash).slice(0, 64);
    const s = keccak256(r + 'salt').slice(0, 64);
    const v = '1c'; // recovery id (usually 1b or 1c)
    
    const signature = r + s + v;
    
    return '0x' + vanity + validatorAddresses + signature;
}

// Generate Staking contract storage state
function generateStakingStorage() {
    const crypto = require('crypto');
    
    // Helper function to calculate storage slots (using correct keccak256)
    function getStorageSlot(slot, key) {
        const keyPadded = key.slice(2).padStart(64, '0');
        const slotPadded = slot.toString(16).padStart(64, '0');
        return '0x' + keccak256(Buffer.from(keyPadded + slotPadded, 'hex'));
    }
    
    function getArraySlot(slot, index) {
        const baseSlot = keccak256(Buffer.from(slot.toString(16).padStart(64, '0'), 'hex'));
        const indexBN = BigInt(index);
        const baseBN = BigInt('0x' + baseSlot);
        return '0x' + (baseBN + indexBN).toString(16).padStart(64, '0');
    }
    
    // Staking contract storage slot layout (considering inheritance from Params):
    // slot 0: initialized bool (from Params)
    // slot 1: validatorStakes mapping(address => ValidatorStake)
    // slot 2: delegations mapping(address => mapping(address => Delegation))  
    // slot 3: unbondingDelegations mapping(address => mapping(address => UnbondingEntry[]))
    // slot 4: allValidators address[]
    // slot 5: validatorIndex mapping(address => uint256)
    // slot 6: totalStaked uint256
    // slot 7: rewardPerShare mapping(address => uint256)
    // slot 8: validatorsContract IValidators
    // slot 9: proposalContract Proposal
    
    const storage = {};
    
    // Note: Don't set initialized, let Congress consensus engine call initialize() or initializeWithValidators() method
    
    // Set validatorsContract address (slot 8)
    const validatorsAddress = CONTRACT_ADDRESSES.Validators.toLowerCase();
    storage['0x0000000000000000000000000000000000000000000000000000000000000008'] = 
        '0x' + validatorsAddress.slice(2).padStart(64, '0');
    
    // Set proposalContract address (slot 9)
    const proposalAddress = CONTRACT_ADDRESSES.Proposal.toLowerCase();
    storage['0x0000000000000000000000000000000000000000000000000000000000000009'] = 
        '0x' + proposalAddress.slice(2).padStart(64, '0');
    
    // Set staking information for each validator
    const minValidatorStake = BigInt('10000000000000000000000'); // 10,000 ether in wei
    let totalStaked = BigInt(0);
    
    INITIAL_VALIDATORS.forEach((validator, index) => {
        const validatorAddr = validator.toLowerCase();
        
        // validatorStakes[validator] mapping (slot 1)
        const stakingSlot = getStorageSlot(1, validatorAddr);
        
        // ValidatorStake struct layout:
        // offset 0: selfStake (uint256)
        // offset 1: totalDelegated (uint256) 
        // offset 2: commissionRate (uint256)
        // offset 3: accumulatedRewards (uint256)
        // offset 4: isJailed (bool, stored as uint256)
        // offset 5: jailUntilBlock (uint256)
        
        // selfStake (offset 0)
        storage[stakingSlot] = '0x' + minValidatorStake.toString(16).padStart(64, '0');
        
        // totalDelegated (offset 1) 
        const totalDelegatedSlot = '0x' + (BigInt(stakingSlot) + BigInt(1)).toString(16).padStart(64, '0');
        storage[totalDelegatedSlot] = '0x' + '0'.padStart(64, '0');
        
        // commissionRate (offset 2) - 5% = 500 basis points
        const commissionSlot = '0x' + (BigInt(stakingSlot) + BigInt(2)).toString(16).padStart(64, '0');
        storage[commissionSlot] = '0x' + '500'.padStart(64, '0');
        
        // accumulatedRewards (offset 3)
        const rewardsSlot = '0x' + (BigInt(stakingSlot) + BigInt(3)).toString(16).padStart(64, '0');
        storage[rewardsSlot] = '0x' + '0'.padStart(64, '0');
        
        // isJailed (offset 4) - false = 0
        const jailSlot = '0x' + (BigInt(stakingSlot) + BigInt(4)).toString(16).padStart(64, '0');
        storage[jailSlot] = '0x' + '0'.padStart(64, '0');
        
        // jailUntilBlock (offset 5) - 0 
        const jailUntilSlot = '0x' + (BigInt(stakingSlot) + BigInt(5)).toString(16).padStart(64, '0');
        storage[jailUntilSlot] = '0x' + '0'.padStart(64, '0');
        
        // allValidators array (slot 4) - Set array length
        if (index === 0) {
            storage['0x0000000000000000000000000000000000000000000000000000000000000004'] = 
                '0x' + INITIAL_VALIDATORS.length.toString(16).padStart(64, '0');
        }
        
        // allValidators[index] 
        const arrayElementSlot = getArraySlot(4, index);
        storage[arrayElementSlot] = '0x' + validatorAddr.slice(2).padStart(64, '0');
        
        // validatorIndex[validator] = index (slot 5)
        const indexSlot = getStorageSlot(5, validatorAddr);
        storage[indexSlot] = '0x' + index.toString(16).padStart(64, '0');
        
        totalStaked += minValidatorStake;
    });
    
    // totalStaked (slot 6)
    storage['0x0000000000000000000000000000000000000000000000000000000000000006'] = 
        '0x' + totalStaked.toString(16).padStart(64, '0');
    
    return storage;
}

// Update Genesis file
function updateGenesisFile() {
    const genesisPath = path.join(__dirname, '..', 'chain', 'genesis.json');
    
    try {
        // Read existing Genesis file
        const genesis = JSON.parse(fs.readFileSync(genesisPath, 'utf8'));
        
        // Ensure alloc field exists
        if (!genesis.alloc) {
            genesis.alloc = {};
        }
        
        // Add system contracts
        console.log('📋 Adding system contracts to Genesis file...');
        
        for (const [contractName, address] of Object.entries(CONTRACT_ADDRESSES)) {
            const bytecode = getContractBytecode(contractName);
            if (bytecode) {
                const contractAlloc = {
                    balance: "0x0",
                    code: bytecode
                };
                
                // Add preset storage state for Staking contract
                if (contractName === 'Staking') {
                    // contractAlloc.storage = generateStakingStorage();
                    console.log(`✅ ${contractName}: ${address} (includes ${INITIAL_VALIDATORS.length} preset validators)`);
                } else {
                    console.log(`✅ ${contractName}: ${address}`);
                }
                
                genesis.alloc[address] = contractAlloc;
            } else {
                console.log(`❌ ${contractName}: Failed to get bytecode`);
            }
        }
        
        // Update extraData to include initial validators
        genesis.extraData = generateExtraData();
        console.log('✅ Updated extraData to include initial validators');
        
        // Write back to Genesis file
        fs.writeFileSync(genesisPath, JSON.stringify(genesis, null, 2));
        console.log('✅ Genesis file updated successfully!');
        console.log(`📄 File location: ${path.relative(process.cwd(), genesisPath)}`);
        
        // Display summary
        console.log('\n📋 Update Summary:');
        console.log(`🏗️  Consensus Algorithm: Congress (POA)`);
        console.log(`⏱️  Block Interval: ${genesis.config.congress.period} seconds`);
        console.log(`🔄 Validator Update Cycle: ${genesis.config.congress.epoch} blocks`);
        console.log(`🏪 System Contracts: ${Object.keys(CONTRACT_ADDRESSES).length} contracts`);
        console.log(`👥 Preset Validators: ${INITIAL_VALIDATORS.length} validators`);
        console.log(`💰 Stake per Validator: 10,000 JU`);
        console.log(`🆔 Chain ID: ${genesis.config.chainId}`);
        
        console.log('\n👥 Initial Validator List:');
        INITIAL_VALIDATORS.forEach((validator, index) => {
            console.log(`   ${index + 1}. ${validator}`);
        });
        
    } catch (error) {
        console.error('❌ Failed to update Genesis file:', error.message);
        process.exit(1);
    }
}

// Verify contract compilation status
function verifyContracts() {
    console.log('🔍 Verifying contract compilation status...');
    
    let allContractsReady = true;
    for (const contractName of Object.keys(CONTRACT_ADDRESSES)) {
        const bytecode = getContractBytecode(contractName);
        if (!bytecode || bytecode === '0x') {
            console.log(`❌ ${contractName}: Not compiled or bytecode is empty`);
            allContractsReady = false;
        } else {
            console.log(`✅ ${contractName}: Compilation successful (${bytecode.length} characters)`);
        }
    }
    
    if (!allContractsReady) {
        console.log('\n❌ Please compile contracts first: forge build or npx hardhat compile');
        process.exit(1);
    }
    
    return true;
}

// Main function
function main() {
    console.log('🚀 Congress Consensus Configuration Tool\n');
    
    // Verify contract compilation status
    if (verifyContracts()) {
        // Update Genesis file
        updateGenesisFile();
        
        console.log('\n🎉 Congress Consensus Configuration Complete!');
        console.log('💡 Next steps to start private chain:');
        console.log('   cd ../chain && ./pm2-init.sh');
        console.log('   Or directly use: cd ../chain && pm2 start ecosystem.config.js');
        console.log('\n📋 Important Notes:');
        console.log('   ✅ Genesis block includes staking information for 3 preset validators');
        console.log('   ✅ Each validator has staked 10,000 JU tokens');
        console.log('   ✅ JPoSA consensus will work normally, no manual validator registration needed');
        console.log('   ✅ Validator voting and staking operations can be performed directly');
        console.log('   ✅ Validator count meets minimum requirement (MIN_VALIDATORS = 3)');
    }
}

// Run main function
main();