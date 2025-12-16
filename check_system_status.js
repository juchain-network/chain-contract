#!/usr/bin/env node

// JPoSA System Status Comprehensive Checker
// Check system contract initialization status and validator details
const { Web3 } = require('web3');

async function checkSystemStatus() {
    const web3 = new Web3('http://localhost:8545');
    
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
    
    const contracts = {
        'Validators': '0x000000000000000000000000000000000000f000',
        'Punish': '0x000000000000000000000000000000000000f001', 
        'Proposal': '0x000000000000000000000000000000000000f002',
        'Staking': '0x000000000000000000000000000000000000f003'
    };
    
    let allInitialized = true;
    let stakingAddress = contracts['Staking'];
    let validatorCount = 0;
    let totalStaked = BigInt(0);
    
    for (const [name, address] of Object.entries(contracts)) {
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
            
            // Check allValidators array length (slot 4)
            const arrayLengthSlot = await web3.eth.getStorageAt(address, '0x4');
            validatorCount = parseInt(arrayLengthSlot, 16);
            console.log(`      Validator Count: ${validatorCount}`);
            
            // Check totalStaked (slot 6)
            const totalStakedSlot = await web3.eth.getStorageAt(address, '0x6');
            totalStaked = BigInt(totalStakedSlot);
            console.log(`      Total Staked: ${totalStaked / BigInt('1000000000000000000')} JU`);
            
            // Check MIN_VALIDATOR_STAKE
            try {
                const minStakeData = '0x9c2a2259'; // MIN_VALIDATOR_STAKE()
                const minStakeResult = await web3.eth.call({ to: address, data: minStakeData });
                const minStake = BigInt(minStakeResult);
                console.log(`      Minimum Stake Requirement: ${minStake / BigInt('1000000000000000000')} JU`);
            } catch (error) {
                console.log(`      Minimum Stake Requirement: Unable to fetch`);
            }
            
            // Test getTopValidators method
            try {
                const getTopValidatorsData = '0x93a5b1b6' + // getTopValidators(uint256)
                    '0000000000000000000000000000000000000000000000000000000000000015'; // 21
                
                const result = await web3.eth.call({
                    to: address,
                    data: getTopValidatorsData
                });
                
                if (result && result !== '0x') {
                    // Parse result
                    const resultData = result.slice(2);
                    const offset = parseInt(resultData.slice(0, 64), 16);
                    const lengthStart = offset * 2;
                    const topValidatorCount = parseInt(resultData.slice(lengthStart, lengthStart + 64), 16);
                    console.log(`      ✅ getTopValidators returned ${topValidatorCount} validators`);
                    
                    if (topValidatorCount !== validatorCount) {
                        console.log(`      ⚠️  Return count mismatch with storage count (${topValidatorCount} vs ${validatorCount})`);
                    }
                } else {
                    console.log(`      ❌ getTopValidators call failed`);
                    allInitialized = false;
                }
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
        
        const validators = [
            '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
            '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
            '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
            '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
            '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
            '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc', // Validator 6
            '0x50C554aC9c134491818fa6f21d504f2AE5BD9c26'  // Validator 7
        ];
        
        let qualifiedValidators = 0;
        
        for (let i = 0; i < Math.min(validators.length, validatorCount); i++) {
            const validator = validators[i];
            console.log(`👤 Validator ${i + 1}: ${validator}`);
            
            try {
                // Call getValidatorInfo method
                const getValidatorInfoData = '0xaa735578' + // getValidatorInfo(address)
                    validator.slice(2).padStart(64, '0');
                
                const result = await web3.eth.call({
                    to: stakingAddress,
                    data: getValidatorInfoData
                });
                
                if (result && result !== '0x') {
                    const resultData = result.slice(2);
                    
                    // Parse return data (5 return values)
                    const selfStake = BigInt('0x' + resultData.slice(0, 64));
                    const totalDelegated = BigInt('0x' + resultData.slice(64, 128));
                    const commissionRate = BigInt('0x' + resultData.slice(128, 192));
                    const isJailed = BigInt('0x' + resultData.slice(192, 256)) !== BigInt(0);
                    const jailUntilBlock = BigInt('0x' + resultData.slice(256, 320));
                    
                    console.log(`   Stake Info:`);
                    console.log(`     selfStake: ${selfStake / BigInt('1000000000000000000')} JU`);
                    console.log(`     totalDelegated: ${totalDelegated / BigInt('1000000000000000000')} JU`);
                    console.log(`     commissionRate: ${commissionRate} (${Number(commissionRate) / 100}%)`);
                    console.log(`   Status Info:`);
                    console.log(`     isJailed: ${isJailed}`);
                    console.log(`     jailUntilBlock: ${jailUntilBlock}`);
                    
                    // Check if requirements are met
                    const MIN_VALIDATOR_STAKE = BigInt('10000000000000000000000'); // 10,000 JU
                    const meetsStakeRequirement = selfStake >= MIN_VALIDATOR_STAKE;
                    const meetsJailRequirement = !isJailed || BigInt(blockNumber) >= jailUntilBlock;
                    
                    console.log(`   Compliance Check:`);
                    console.log(`     ${meetsStakeRequirement ? '✅' : '❌'} Stake Requirement: ${selfStake / BigInt('1000000000000000000')} >= ${MIN_VALIDATOR_STAKE / BigInt('1000000000000000000')} JU`);
                    console.log(`     ${meetsJailRequirement ? '✅' : '❌'} Jail Status: ${!isJailed ? 'Not jailed' : `Jailed until block ${jailUntilBlock}`}`);
                    
                    const isQualified = meetsStakeRequirement && meetsJailRequirement;
                    console.log(`   📊 Overall Status: ${isQualified ? '✅ Qualified' : '❌ Not Qualified'}`);
                    
                    if (isQualified) {
                        qualifiedValidators++;
                    }
                    
                } else {
                    console.log(`   ❌ Unable to get validator info`);
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
        console.log(`   💰 Total Staked: ${totalStaked / BigInt('1000000000000000000')} JU`);
        console.log(`   📊 Current Block: ${blockNumber}`);
        
        // MIN_VALIDATORS is now 3 (changed from 5)
        const MIN_VALIDATORS = 3;
        if (allInitialized && qualifiedValidators >= MIN_VALIDATORS) {
            console.log(`   🚀 JPoSA Consensus Status: ✅ System Running Normally (${qualifiedValidators} >= ${MIN_VALIDATORS})`);
        } else if (allInitialized && qualifiedValidators > 0 && qualifiedValidators < MIN_VALIDATORS) {
            console.log(`   ⚠️  JPoSA Consensus Status: ⚠️  Insufficient Validators (${qualifiedValidators} < ${MIN_VALIDATORS})`);
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