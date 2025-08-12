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
            console.error(`❌ 无法读取 ${contractName} 合约字节码:`, error.message);
            return null;
        }
    }
}

// 生成初始验证者的 extraData
// 使用预分配账户作为初始验证者
function generateExtraData() {
    const crypto = require('crypto');
    const { keccak256 } = require('js-sha3');
    
    // 初始验证者地址
    const initialValidator = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    
    // 构建 extraData 结构:
    // 32字节 vanity + N*20字节 validator addresses + 65字节 signature
    const vanity = '0'.repeat(64); // 32字节的vanity (可以是任意数据)
    
    // 验证者地址列表 (去掉0x前缀，保持20字节)
    const validatorHex = initialValidator.slice(2).toLowerCase();
    
    // 为了生成正确的签名，我们需要：
    // 1. 构建要签名的数据（不包含签名部分）
    const dataToSign = vanity + validatorHex;
    
    // 2. 对数据进行keccak256哈希
    const hash = keccak256(Buffer.from(dataToSign, 'hex'));
    
    // 3. 生成一个简单的签名（在实际场景中，这应该是由验证者私钥签名）
    // 这里我们使用一个确定性的方法来生成签名
    const messageHash = Buffer.from(hash, 'hex');
    
    // 生成一个固定的签名（65字节：32字节r + 32字节s + 1字节v）
    // 注意：这不是真正的ECDSA签名，只是为了格式正确
    const r = crypto.createHash('sha256').update(messageHash).digest('hex').slice(0, 64);
    const s = crypto.createHash('sha256').update(r + 'salt').digest('hex').slice(0, 64);
    const v = '1c'; // recovery id (通常是1b或1c)
    
    const signature = r + s + v;
    
    return '0x' + vanity + validatorHex + signature;
}

// 更新 Genesis 文件
function updateGenesisFile() {
    const genesisPath = path.join(__dirname, 'genesis.json');
    
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
        console.log('\n❌ 请先编译合约: forge build 或 npx hardhat compile');
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
