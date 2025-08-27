#!/usr/bin/env node

// JPoSA 系统状态综合检查器
// 检查系统合约初始化状态和验证者详细信息
const { Web3 } = require('web3');

async function checkSystemStatus() {
    const web3 = new Web3('http://localhost:8545');
    
    console.log('🔍 JPoSA 系统状态综合检查器\n');
    console.log('='.repeat(80));
    
    // 1. 检查节点连接
    let blockNumber, networkId;
    try {
        blockNumber = await web3.eth.getBlockNumber();
        networkId = await web3.eth.net.getId();
        console.log(`📡 节点连接: ✅ 正常 (网络 ID: ${networkId}, 区块: ${blockNumber})\n`);
    } catch (error) {
        console.log('📡 节点连接: ❌ 失败 - 请检查节点是否启动\n');
        return;
    }

    // 2. 检查系统合约状态
    console.log('🏛️  系统合约状态检查');
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
        console.log(`📋 ${name} 合约 (${address}):`);
        
        // 检查合约代码是否存在
        const code = await web3.eth.getCode(address);
        if (code === '0x') {
            console.log(`   ❌ 合约未部署\n`);
            allInitialized = false;
            continue;
        }
        console.log(`   ✅ 合约已部署 (${Math.floor(code.length / 2)} 字节)`);
        
        // 检查 initialized 状态 (slot 0)
        const initializedSlot = await web3.eth.getStorageAt(address, '0x0');
        // 对于继承了 Params 的合约，initialized 是 bool 类型，只需要检查最低位
        const isInitialized = (BigInt(initializedSlot) & BigInt(1)) === BigInt(1);
        console.log(`   ${isInitialized ? '✅' : '❌'} 初始化状态: ${isInitialized ? '已初始化' : '未初始化'}`);
        
        if (!isInitialized) {
            allInitialized = false;
        }
        
        // 对于 Staking 合约，检查关键数据
        if (name === 'Staking' && isInitialized) {
            console.log('   📊 详细状态:');
            
            // 检查 allValidators 数组长度 (slot 4)
            const arrayLengthSlot = await web3.eth.getStorageAt(address, '0x4');
            validatorCount = parseInt(arrayLengthSlot, 16);
            console.log(`      验证者数量: ${validatorCount}`);
            
            // 检查 totalStaked (slot 6)
            const totalStakedSlot = await web3.eth.getStorageAt(address, '0x6');
            totalStaked = BigInt(totalStakedSlot);
            console.log(`      总质押量: ${totalStaked / BigInt('1000000000000000000')} JU`);
            
            // 检查 MIN_VALIDATOR_STAKE
            try {
                const minStakeData = '0x9c2a2259'; // MIN_VALIDATOR_STAKE()
                const minStakeResult = await web3.eth.call({ to: address, data: minStakeData });
                const minStake = BigInt(minStakeResult);
                console.log(`      最小质押要求: ${minStake / BigInt('1000000000000000000')} JU`);
            } catch (error) {
                console.log(`      最小质押要求: 无法获取`);
            }
            
            // 测试 getTopValidators 方法
            try {
                const getTopValidatorsData = '0x93a5b1b6' + // getTopValidators(uint256)
                    '0000000000000000000000000000000000000000000000000000000000000015'; // 21
                
                const result = await web3.eth.call({
                    to: address,
                    data: getTopValidatorsData
                });
                
                if (result && result !== '0x') {
                    // 解析结果
                    const resultData = result.slice(2);
                    const offset = parseInt(resultData.slice(0, 64), 16);
                    const lengthStart = offset * 2;
                    const topValidatorCount = parseInt(resultData.slice(lengthStart, lengthStart + 64), 16);
                    console.log(`      ✅ getTopValidators 返回 ${topValidatorCount} 个验证者`);
                    
                    if (topValidatorCount !== validatorCount) {
                        console.log(`      ⚠️  返回数量与存储数量不匹配 (${topValidatorCount} vs ${validatorCount})`);
                    }
                } else {
                    console.log(`      ❌ getTopValidators 调用失败`);
                    allInitialized = false;
                }
            } catch (error) {
                console.log(`      ❌ getTopValidators 执行错误: ${error.message}`);
                allInitialized = false;
            }
        }
        
        console.log();
    }

    // 3. 检查验证者详细状态
    if (allInitialized && validatorCount > 0) {
        console.log('👥 验证者详细状态检查');
        console.log('-'.repeat(40));
        
        const validators = [
            '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
            '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
            '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
            '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
            '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
            '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc', // 验证者6
            '0x50C554aC9c134491818fa6f21d504f2AE5BD9c26'  // 验证者7
        ];
        
        let qualifiedValidators = 0;
        
        for (let i = 0; i < Math.min(validators.length, validatorCount); i++) {
            const validator = validators[i];
            console.log(`👤 验证者 ${i + 1}: ${validator}`);
            
            try {
                // 调用 getValidatorInfo 方法
                const getValidatorInfoData = '0xaa735578' + // getValidatorInfo(address)
                    validator.slice(2).padStart(64, '0');
                
                const result = await web3.eth.call({
                    to: stakingAddress,
                    data: getValidatorInfoData
                });
                
                if (result && result !== '0x') {
                    const resultData = result.slice(2);
                    
                    // 解析返回数据 (5个返回值)
                    const selfStake = BigInt('0x' + resultData.slice(0, 64));
                    const totalDelegated = BigInt('0x' + resultData.slice(64, 128));
                    const commissionRate = BigInt('0x' + resultData.slice(128, 192));
                    const isJailed = BigInt('0x' + resultData.slice(192, 256)) !== BigInt(0);
                    const jailUntilBlock = BigInt('0x' + resultData.slice(256, 320));
                    
                    console.log(`   质押信息:`);
                    console.log(`     selfStake: ${selfStake / BigInt('1000000000000000000')} JU`);
                    console.log(`     totalDelegated: ${totalDelegated / BigInt('1000000000000000000')} JU`);
                    console.log(`     commissionRate: ${commissionRate} (${Number(commissionRate) / 100}%)`);
                    console.log(`   状态信息:`);
                    console.log(`     isJailed: ${isJailed}`);
                    console.log(`     jailUntilBlock: ${jailUntilBlock}`);
                    
                    // 检查是否满足条件
                    const MIN_VALIDATOR_STAKE = BigInt('10000000000000000000000'); // 10,000 JU
                    const meetsStakeRequirement = selfStake >= MIN_VALIDATOR_STAKE;
                    const meetsJailRequirement = !isJailed || BigInt(blockNumber) >= jailUntilBlock;
                    
                    console.log(`   合规性检查:`);
                    console.log(`     ${meetsStakeRequirement ? '✅' : '❌'} 质押要求: ${selfStake / BigInt('1000000000000000000')} >= ${MIN_VALIDATOR_STAKE / BigInt('1000000000000000000')} JU`);
                    console.log(`     ${meetsJailRequirement ? '✅' : '❌'} 监禁状态: ${!isJailed ? '未监禁' : `监禁至区块 ${jailUntilBlock}`}`);
                    
                    const isQualified = meetsStakeRequirement && meetsJailRequirement;
                    console.log(`   📊 综合状态: ${isQualified ? '✅ 符合条件' : '❌ 不符合条件'}`);
                    
                    if (isQualified) {
                        qualifiedValidators++;
                    }
                    
                } else {
                    console.log(`   ❌ 无法获取验证者信息`);
                }
                
            } catch (error) {
                console.log(`   ❌ 查询错误: ${error.message}`);
            }
            
            console.log();
        }
        
        // 4. 系统状态总结
        console.log('='.repeat(80));
        console.log(`🎯 系统状态总结:`);
        console.log(`   📋 系统合约: ${allInitialized ? '✅ 全部正常' : '❌ 存在问题'}`);
        console.log(`   👥 注册验证者: ${validatorCount} 个`);
        console.log(`   ✅ 合规验证者: ${qualifiedValidators} 个`);
        console.log(`   💰 总质押量: ${totalStaked / BigInt('1000000000000000000')} JU`);
        console.log(`   📊 当前区块: ${blockNumber}`);
        
        if (allInitialized && qualifiedValidators >= 1) {
            console.log(`   🚀 JPoSA 共识状态: ✅ 系统正常运行`);
        } else if (allInitialized && qualifiedValidators === 0) {
            console.log(`   ⚠️  JPoSA 共识状态: ❌ 无合规验证者`);
        } else {
            console.log(`   ⚠️  JPoSA 共识状态: ❌ 系统存在问题`);
        }
        
    } else {
        console.log('='.repeat(80));
        console.log(`🎯 系统状态总结:`);
        console.log(`   📋 系统合约: ${allInitialized ? '✅ 全部正常' : '❌ 存在问题'}`);
        console.log(`   👥 验证者数量: ${validatorCount}`);
        
        if (!allInitialized) {
            console.log(`   ⚠️  系统存在问题，无法检查验证者状态`);
        } else if (validatorCount === 0) {
            console.log(`   ⚠️  系统正常但无验证者注册`);
        }
    }
    
    console.log('='.repeat(80));
    
    // 5. 执行交易测试
    console.log('\n💸 交易功能测试');
    console.log('-'.repeat(40));
    
    try {
        // 检查发送方余额
        const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
        const toAddress = '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc';
        
        const fromBalanceBefore = await web3.eth.getBalance(fromAddress);
        const toBalanceBefore = await web3.eth.getBalance(toAddress);
        
        console.log(`📊 转账前余额:`);
        console.log(`   发送方 (${fromAddress}): ${web3.utils.fromWei(fromBalanceBefore, 'ether')} ETH`);
        console.log(`   接收方 (${toAddress}): ${web3.utils.fromWei(toBalanceBefore, 'ether')} ETH`);
        
        // 执行转账交易
        const transferAmount = web3.utils.toWei('100', 'ether'); // 转账100 ETH
        const gasPrice = web3.utils.toWei('20', 'gwei');
        
        console.log(`\n🚀 执行转账交易:`);
        console.log(`   转账金额: 100 ETH`);
        console.log(`   Gas 价格: 20 Gwei`);
        
        const txHash = await web3.eth.sendTransaction({
            from: fromAddress,
            to: toAddress,
            value: transferAmount,
            gas: 21000,
            gasPrice: gasPrice
        });
        
        console.log(`   ✅ 交易发送成功`);
        console.log(`   📋 交易哈希: ${typeof txHash === 'object' ? txHash.transactionHash || JSON.stringify(txHash) : txHash}`);
        
        // 等待交易确认并获取收据
        let receipt = null;
        let attempts = 0;
        const maxAttempts = 10;
        
        console.log(`   ⏳ 等待交易确认...`);
        
        const actualTxHash = typeof txHash === 'object' ? txHash.transactionHash : txHash;
        
        while (!receipt && attempts < maxAttempts) {
            try {
                receipt = await web3.eth.getTransactionReceipt(actualTxHash);
                if (!receipt) {
                    await new Promise(resolve => setTimeout(resolve, 2000));
                    attempts++;
                    console.log(`   ⏳ 等待中... (${attempts}/${maxAttempts})`);
                }
            } catch (error) {
                await new Promise(resolve => setTimeout(resolve, 2000));
                attempts++;
                console.log(`   ⏳ 等待中... (${attempts}/${maxAttempts})`);
            }
        }
        
        if (receipt) {
            console.log(`   ✅ 交易已确认`);
            console.log(`   📦 区块号: ${receipt.blockNumber}`);
            console.log(`   ⛽ Gas 使用: ${receipt.gasUsed}`);

            const txStatus = receipt.status;
            // Web3.js 返回的 status 是 BigInt 类型
            const isSuccess = txStatus === BigInt(1) || txStatus === 1 || txStatus === '0x1' || txStatus === true;
            console.log(`   📊 状态: ${isSuccess ? '成功' : '失败'} (原始值: ${txStatus}, 类型: ${typeof txStatus})`);
            
            // 检查转账后余额
            const fromBalanceAfter = await web3.eth.getBalance(fromAddress);
            const toBalanceAfter = await web3.eth.getBalance(toAddress);
            
            console.log(`\n📊 转账后余额:`);
            console.log(`   发送方 (${fromAddress}): ${web3.utils.fromWei(fromBalanceAfter, 'ether')} ETH`);
            console.log(`   接收方 (${toAddress}): ${web3.utils.fromWei(toBalanceAfter, 'ether')} ETH`);
            
            // 计算实际变化
            const fromChange = BigInt(fromBalanceAfter) - BigInt(fromBalanceBefore);
            const toChange = BigInt(toBalanceAfter) - BigInt(toBalanceBefore);
            const gasCost = BigInt(receipt.gasUsed) * BigInt(gasPrice);
            
            console.log(`\n📈 余额变化:`);
            console.log(`   发送方减少: ${web3.utils.fromWei((-fromChange).toString(), 'ether')} ETH`);
            console.log(`   接收方增加: ${web3.utils.fromWei(toChange.toString(), 'ether')} ETH`);
            console.log(`   Gas 费用: ${web3.utils.fromWei(gasCost.toString(), 'ether')} ETH`);
            
            // 验证余额变化是否正确
            const expectedFromChange = -(BigInt(transferAmount) + gasCost);
            const expectedToChange = BigInt(transferAmount);
            
            const fromCorrect = fromChange === expectedFromChange;
            const toCorrect = toChange === expectedToChange;
            
            console.log(`\n✅ 验证结果:`);
            console.log(`   发送方余额变化: ${fromCorrect ? '✅ 正确' : '❌ 错误'}`);
            console.log(`   接收方余额变化: ${toCorrect ? '✅ 正确' : '❌ 错误'}`);
            console.log(`   💸 交易测试: ${isSuccess ? '✅ 成功' : '❌ 失败'}`);
            
        } else {
            console.log(`   ❌ 交易确认超时`);
        }
        
    } catch (error) {
        console.log(`   ❌ 交易测试失败: ${error.message}`);
    }
    
    console.log('='.repeat(80));
}

checkSystemStatus().catch(error => {
    console.error('❌ 检查过程中发生错误:', error.message);
    process.exit(1);
});
