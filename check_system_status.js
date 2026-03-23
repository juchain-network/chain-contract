#!/usr/bin/env node

// JPoSA System Status Comprehensive Checker
// Check system contract initialization status and validator details
const { Web3 } = require('web3');

const JU = 10n ** 18n;

const CONTRACTS = {
    Validators: '0x000000000000000000000000000000000000f010',
    Punish: '0x000000000000000000000000000000000000f011',
    Proposal: '0x000000000000000000000000000000000000f012',
    Staking: '0x000000000000000000000000000000000000f013'
};

const proposalAbi = [
    {
        inputs: [],
        name: 'minValidatorStake',
        outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
        stateMutability: 'view',
        type: 'function'
    }
];

const stakingAbi = [
    {
        inputs: [],
        name: 'getValidatorCount',
        outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
        stateMutability: 'view',
        type: 'function'
    },
    {
        inputs: [],
        name: 'totalStaked',
        outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
        stateMutability: 'view',
        type: 'function'
    },
    {
        inputs: [{ internalType: 'address', name: 'validator', type: 'address' }],
        name: 'getValidatorInfo',
        outputs: [
            { internalType: 'uint256', name: 'selfStake', type: 'uint256' },
            { internalType: 'uint256', name: 'totalDelegated', type: 'uint256' },
            { internalType: 'uint256', name: 'commissionRate', type: 'uint256' },
            { internalType: 'uint256', name: 'accumulatedRewards', type: 'uint256' },
            { internalType: 'bool', name: 'isJailed', type: 'bool' },
            { internalType: 'uint256', name: 'jailUntilBlock', type: 'uint256' },
            { internalType: 'uint256', name: 'totalClaimedRewards', type: 'uint256' },
            { internalType: 'uint256', name: 'lastClaimBlock', type: 'uint256' },
            { internalType: 'bool', name: 'isRegistered', type: 'bool' },
            { internalType: 'uint256', name: 'totalRewards', type: 'uint256' }
        ],
        stateMutability: 'view',
        type: 'function'
    }
];

const validatorsAbi = [
    {
        inputs: [],
        name: 'getTopValidators',
        outputs: [{ internalType: 'address[]', name: '', type: 'address[]' }],
        stateMutability: 'view',
        type: 'function'
    },
    {
        inputs: [],
        name: 'getActiveValidators',
        outputs: [{ internalType: 'address[]', name: '', type: 'address[]' }],
        stateMutability: 'view',
        type: 'function'
    }
];

async function checkSystemStatus() {
    const web3 = new Web3('http://localhost:8545');
    const proposal = new web3.eth.Contract(proposalAbi, CONTRACTS.Proposal);
    const staking = new web3.eth.Contract(stakingAbi, CONTRACTS.Staking);
    const validatorsContract = new web3.eth.Contract(validatorsAbi, CONTRACTS.Validators);
    
    console.log('🔍 JPoSA System Status Comprehensive Checker\n');
    console.log('='.repeat(80));
    
    // 1. Check node connection
    let blockNumber, networkId;
    try {
        blockNumber = await web3.eth.getBlockNumber();
        networkId = await web3.eth.net.getId();
        console.log(`📡 Node Connection: ✅ Normal (Network ID: ${networkId}, Block: ${blockNumber})\n`);
    } catch (error) {
        console.log('📡 Node Connection: ❌ Failed - Please check if node is running\n');
        return;
    }

    // 2. Check system contract status
    console.log('🏛️  System Contract Status Check');
    console.log('-'.repeat(40));
    
    let allInitialized = true;
    let stakingAddress = CONTRACTS['Staking'];
    let validatorCount = 0;
    let totalStaked = BigInt(0);
    let minValidatorStake = 0n;
    
    for (const [name, address] of Object.entries(CONTRACTS)) {
        console.log(`📋 ${name} Contract (${address}):`);
        
        // Check if contract code exists
        const code = await web3.eth.getCode(address);
        if (code === '0x') {
            console.log(`   ❌ Contract not deployed\n`);
            allInitialized = false;
            continue;
        }
        console.log(`   ✅ Contract deployed (${Math.floor(code.length / 2)} bytes)`);
        
        // Check initialized status (slot 0)
        const initializedSlot = await web3.eth.getStorageAt(address, '0x0');
        // For contracts inheriting Params, initialized is bool type, only need to check the least significant bit
        const isInitialized = (BigInt(initializedSlot) & BigInt(1)) === BigInt(1);
        console.log(`   ${isInitialized ? '✅' : '❌'} Initialization Status: ${isInitialized ? 'Initialized' : 'Not Initialized'}`);
        
        if (!isInitialized) {
            allInitialized = false;
        }
        
        // For Staking contract, check key data
        if (name === 'Staking' && isInitialized) {
            console.log('   📊 Detailed Status:');
            
            validatorCount = Number(await staking.methods.getValidatorCount().call());
            console.log(`      Validator Count: ${validatorCount}`);
            
            totalStaked = BigInt(await staking.methods.totalStaked().call());
            console.log(`      Total Staked: ${totalStaked / JU} JU`);
            
            minValidatorStake = BigInt(await proposal.methods.minValidatorStake().call());
            console.log(`      Minimum Stake Requirement: ${minValidatorStake / JU} JU`);
            
            // Test getTopValidators method
            try {
                const topValidators = await validatorsContract.methods.getTopValidators().call();
                console.log(`      ✅ getTopValidators returned ${topValidators.length} validators`);
            } catch (error) {
                console.log(`      ❌ getTopValidators execution error: ${error.message}`);
                allInitialized = false;
            }
        }
        
        console.log();
    }

    // 3. Check validator detailed status
    if (allInitialized && validatorCount > 0) {
        console.log('👥 Validator Detailed Status Check');
        console.log('-'.repeat(40));
        
        const validators = await validatorsContract.methods.getActiveValidators().call();
        
        let qualifiedValidators = 0;
        
        for (let i = 0; i < validators.length; i++) {
            const validator = validators[i];
            console.log(`👤 Validator ${i + 1}: ${validator}`);
            
            try {
                const info = await staking.methods.getValidatorInfo(validator).call();
                const selfStake = BigInt(info.selfStake);
                const totalDelegated = BigInt(info.totalDelegated);
                const commissionRate = BigInt(info.commissionRate);
                const isJailed = info.isJailed;
                const jailUntilBlock = BigInt(info.jailUntilBlock);
                const isRegistered = info.isRegistered;

                console.log(`   Stake Info:`);
                console.log(`     selfStake: ${selfStake / JU} JU`);
                console.log(`     totalDelegated: ${totalDelegated / JU} JU`);
                console.log(`     commissionRate: ${commissionRate} (${Number(commissionRate) / 100}%)`);
                console.log(`   Status Info:`);
                console.log(`     isRegistered: ${isRegistered}`);
                console.log(`     isJailed: ${isJailed}`);
                console.log(`     jailUntilBlock: ${jailUntilBlock}`);
                
                const meetsStakeRequirement = selfStake >= minValidatorStake;
                const meetsJailRequirement = !isJailed || BigInt(blockNumber) >= jailUntilBlock;
                
                console.log(`   Compliance Check:`);
                console.log(`     ${meetsStakeRequirement ? '✅' : '❌'} Stake Requirement: ${selfStake / JU} >= ${minValidatorStake / JU} JU`);
                console.log(`     ${meetsJailRequirement ? '✅' : '❌'} Jail Status: ${!isJailed ? 'Not jailed' : `Jailed until block ${jailUntilBlock}`}`);
                
                const isQualified = isRegistered && meetsStakeRequirement && meetsJailRequirement;
                console.log(`   📊 Overall Status: ${isQualified ? '✅ Qualified' : '❌ Not Qualified'}`);
                
                if (isQualified) {
                    qualifiedValidators++;
                }
                
            } catch (error) {
                console.log(`   ❌ Query error: ${error.message}`);
            }
            
            console.log();
        }
        
        // 4. System status summary
        console.log('='.repeat(80));
        console.log(`🎯 System Status Summary:`);
        console.log(`   📋 System Contracts: ${allInitialized ? '✅ All Normal' : '❌ Issues Found'}`);
        console.log(`   👥 Registered Validators: ${validatorCount} validators`);
        console.log(`   ✅ Qualified Validators: ${qualifiedValidators} validators`);
        console.log(`   💰 Total Staked: ${totalStaked / JU} JU`);
        console.log(`   📊 Current Block: ${blockNumber}`);
        
        if (allInitialized && qualifiedValidators > 0) {
            console.log(`   🚀 JPoSA Consensus Status: ✅ At least one qualified validator is active (${qualifiedValidators})`);
        } else if (allInitialized && qualifiedValidators === 0) {
            console.log(`   ⚠️  JPoSA Consensus Status: ❌ No Qualified Validators`);
        } else {
            console.log(`   ⚠️  JPoSA Consensus Status: ❌ System Issues Found`);
        }
        
    } else {
        console.log('='.repeat(80));
        console.log(`🎯 System Status Summary:`);
        console.log(`   📋 System Contracts: ${allInitialized ? '✅ All Normal' : '❌ Issues Found'}`);
        console.log(`   👥 Validator Count: ${validatorCount}`);
        
        if (!allInitialized) {
            console.log(`   ⚠️  System has issues, unable to check validator status`);
        } else if (validatorCount === 0) {
            console.log(`   ⚠️  System normal but no validators registered`);
        }
    }
    
    console.log('='.repeat(80));
    
    // 5. Execute transaction test
    console.log('\n💸 Transaction Function Test');
    console.log('-'.repeat(40));
    
    try {
        // Check sender balance
        const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
        const toAddress = '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc';
        
        const fromBalanceBefore = await web3.eth.getBalance(fromAddress);
        const toBalanceBefore = await web3.eth.getBalance(toAddress);
        
        console.log(`📊 Balance Before Transfer:`);
        console.log(`   Sender (${fromAddress}): ${web3.utils.fromWei(fromBalanceBefore, 'ether')} ETH`);
        console.log(`   Receiver (${toAddress}): ${web3.utils.fromWei(toBalanceBefore, 'ether')} ETH`);
        
        // Execute transfer transaction
        const transferAmount = web3.utils.toWei('100', 'ether'); // Transfer 100 ETH
        const gasPrice = web3.utils.toWei('20', 'gwei');
        
        console.log(`\n🚀 Execute Transfer Transaction:`);
        console.log(`   Transfer Amount: 100 ETH`);
        console.log(`   Gas Price: 20 Gwei`);
        
        // Get correct nonce
        const nonce = await web3.eth.getTransactionCount(fromAddress, 'pending');
        console.log(`   Nonce: ${nonce}`);
        
        const txHash = await web3.eth.sendTransaction({
            from: fromAddress,
            to: toAddress,
            value: transferAmount,
            gas: 21000,
            gasPrice: gasPrice,
            nonce: nonce
        });
        
        console.log(`   ✅ Transaction Sent Successfully`);
        console.log(`   📋 Transaction Hash: ${typeof txHash === 'object' ? txHash.transactionHash || JSON.stringify(txHash) : txHash}`);
        
        // Wait for transaction confirmation and get receipt
        let receipt = null;
        let attempts = 0;
        const maxAttempts = 10;
        
        console.log(`   ⏳ Waiting for transaction confirmation...`);
        
        const actualTxHash = typeof txHash === 'object' ? txHash.transactionHash : txHash;
        
        while (!receipt && attempts < maxAttempts) {
            try {
                receipt = await web3.eth.getTransactionReceipt(actualTxHash);
                if (!receipt) {
                    await new Promise(resolve => setTimeout(resolve, 2000));
                    attempts++;
                    console.log(`   ⏳ Waiting... (${attempts}/${maxAttempts})`);
                }
            } catch (error) {
                await new Promise(resolve => setTimeout(resolve, 2000));
                attempts++;
                console.log(`   ⏳ Waiting... (${attempts}/${maxAttempts})`);
            }
        }
        
        if (receipt) {
            console.log(`   ✅ Transaction Confirmed`);
            console.log(`   📦 Block Number: ${receipt.blockNumber}`);
            console.log(`   ⛽ Gas Used: ${receipt.gasUsed}`);

            const txStatus = receipt.status;
            // Web3.js returns status as BigInt type
            const isSuccess = txStatus === BigInt(1) || txStatus === 1 || txStatus === '0x1' || txStatus === true;
            console.log(`   📊 Status: ${isSuccess ? 'Success' : 'Failed'} (Raw value: ${txStatus}, Type: ${typeof txStatus})`);
            
            // Check balance after transfer
            const fromBalanceAfter = await web3.eth.getBalance(fromAddress);
            const toBalanceAfter = await web3.eth.getBalance(toAddress);
            
            console.log(`\n📊 Balance After Transfer:`);
            console.log(`   Sender (${fromAddress}): ${web3.utils.fromWei(fromBalanceAfter, 'ether')} ETH`);
            console.log(`   Receiver (${toAddress}): ${web3.utils.fromWei(toBalanceAfter, 'ether')} ETH`);
            
            // Calculate actual changes
            const fromChange = BigInt(fromBalanceAfter) - BigInt(fromBalanceBefore);
            const toChange = BigInt(toBalanceAfter) - BigInt(toBalanceBefore);
            const gasCost = BigInt(receipt.gasUsed) * BigInt(gasPrice);
            
            console.log(`\n📈 Balance Changes:`);
            console.log(`   Sender Decrease: ${web3.utils.fromWei((-fromChange).toString(), 'ether')} ETH`);
            console.log(`   Receiver Increase: ${web3.utils.fromWei(toChange.toString(), 'ether')} ETH`);
            console.log(`   Gas Cost: ${web3.utils.fromWei(gasCost.toString(), 'ether')} ETH`);
            
            // Verify if balance changes are correct
            const expectedFromChange = -(BigInt(transferAmount) + gasCost);
            const expectedToChange = BigInt(transferAmount);
            
            const fromCorrect = fromChange === expectedFromChange;
            const toCorrect = toChange === expectedToChange;
            
            console.log(`\n✅ Verification Results:`);
            console.log(`   Sender Balance Change: ${fromCorrect ? '✅ Correct' : '❌ Incorrect'}`);
            console.log(`   Receiver Balance Change: ${toCorrect ? '✅ Correct' : '❌ Incorrect'}`);
            console.log(`   💸 Transaction Test: ${isSuccess ? '✅ Success' : '❌ Failed'}`);
            
        } else {
            console.log(`   ❌ Transaction Confirmation Timeout`);
        }
        
    } catch (error) {
        console.log(`   ❌ Transaction Test Failed: ${error.message}`);
    }
    
    console.log('='.repeat(80));
}

checkSystemStatus().catch(error => {
    console.error('❌ Error occurred during check:', error.message);
    process.exit(1);
});
