#!/usr/bin/env node

// 提取系统合约字节码并更新 Genesis 文件
const fs = require('fs');
const path = require('path');

console.log('🔧 提取系统合约字节码并更新 Genesis 文件...');

// 合约地址映射
const CONTRACT_ADDRESSES = {
    'Validators': '0x000000000000000000000000000000000000f000',
    'Punish': '0x000000000000000000000000000000000000f001',
    'Proposal': '0x000000000000000000000000000000000000f002'
};

// 读取合约字节码
function getContractBytecode(contractName) {
    const artifactPath = path.join(__dirname, 'artifacts', 'contracts', `${contractName}.sol`, `${contractName}.json`);
    try {
        const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
        return artifact.deployedBytecode;
    } catch (error) {
        console.error(`❌ 无法读取 ${contractName} 合约字节码:`, error.message);
        return null;
    }
}

// 生成初始验证者的 extraData
// 使用预分配账户作为初始验证者
function generateExtraData() {
    // 这里使用一个简化的 extraData，您可以根据需要调整
    const initialValidator = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    // 移除 '0x' 前缀并填充到64字节
    const validatorHex = initialValidator.slice(2).toLowerCase().padStart(40, '0');
    // 创建简单的 extraData: 32字节的vanity + validator列表 + 65字节的signature
    const vanity = '0'.repeat(64); // 32字节的vanity
    const validatorList = validatorHex.padEnd(40, '0'); // 验证者列表
    const signature = '0'.repeat(130); // 65字节的signature
    
    return '0x' + vanity + validatorList + signature;
}

// 更新 Genesis 文件
function updateGenesisFile() {
    const genesisPath = path.join(__dirname, '..', 'genesis.json');
    
    try {
        // 读取现有的 Genesis 文件
        const genesis = JSON.parse(fs.readFileSync(genesisPath, 'utf8'));
        
        // 确保 alloc 字段存在
        if (!genesis.alloc) {
            genesis.alloc = {};
        }
        
        // 添加系统合约
        console.log('📋 添加系统合约到 Genesis 文件...');
        
        for (const [contractName, address] of Object.entries(CONTRACT_ADDRESSES)) {
            const bytecode = getContractBytecode(contractName);
            if (bytecode) {
                genesis.alloc[address] = {
                    balance: "0x0",
                    code: bytecode
                };
                console.log(`✅ ${contractName}: ${address}`);
            } else {
                console.log(`❌ ${contractName}: 字节码获取失败`);
            }
        }
        
        // 更新 extraData 以包含初始验证者
        genesis.extraData = generateExtraData();
        console.log('✅ 更新 extraData 包含初始验证者');
        
        // 写回 Genesis 文件
        fs.writeFileSync(genesisPath, JSON.stringify(genesis, null, 2));
        console.log('✅ Genesis 文件更新成功!');
        
        // 显示摘要
        console.log('\n📋 更新摘要:');
        console.log(`🏗️  共识算法: Congress (POA)`);
        console.log(`⏱️  出块间隔: ${genesis.config.congress.period} 秒`);
        console.log(`🔄 验证者更新周期: ${genesis.config.congress.epoch} 块`);
        console.log(`🏪 系统合约: ${Object.keys(CONTRACT_ADDRESSES).length} 个`);
        console.log(`🆔 链 ID: ${genesis.config.chainId}`);
        
    } catch (error) {
        console.error('❌ 更新 Genesis 文件失败:', error.message);
        process.exit(1);
    }
}

// 验证合约编译状态
function verifyContracts() {
    console.log('🔍 验证合约编译状态...');
    
    let allContractsReady = true;
    for (const contractName of Object.keys(CONTRACT_ADDRESSES)) {
        const bytecode = getContractBytecode(contractName);
        if (!bytecode || bytecode === '0x') {
            console.log(`❌ ${contractName}: 未编译或字节码为空`);
            allContractsReady = false;
        } else {
            console.log(`✅ ${contractName}: 编译成功 (${bytecode.length} 字符)`);
        }
    }
    
    if (!allContractsReady) {
        console.log('\n❌ 请先编译合约: npx hardhat compile');
        process.exit(1);
    }
    
    return true;
}

// 主函数
function main() {
    console.log('🚀 Congress 共识配置工具\n');
    
    // 验证合约编译状态
    if (verifyContracts()) {
        // 更新 Genesis 文件
        updateGenesisFile();
        
        console.log('\n🎉 Congress 共识配置完成!');
        console.log('💡 接下来可以启动私有链:');
        console.log('   cd ../chain && ./start_private_chain.sh');
    }
}

// 运行主函数
main();
