#!/usr/bin/env node

// 提取系统合约字节码并更新 Genesis 文件
const fs = require('fs');
const path = require('path');
const { keccak256 } = require('js-sha3');

console.log('🔧 提取系统合约字节码并更新 Genesis 文件...');

// 合约地址映射
const CONTRACT_ADDRESSES = {
    'Validators': '0x000000000000000000000000000000000000f000',
    'Punish': '0x000000000000000000000000000000000000f001',
    'Proposal': '0x000000000000000000000000000000000000f002',
    'Staking': '0x000000000000000000000000000000000000f003'
};

// 初始验证者信息
const INITIAL_VALIDATORS = [
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
    '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
    '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65'
];

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

// keccak256 计算助手函数
function keccak256Hash(data) {
    if (typeof data === 'string' && data.startsWith('0x')) {
        return '0x' + keccak256(Buffer.from(data.slice(2), 'hex'));
    }
    return '0x' + keccak256(data);
}

// 生成初始验证者的 extraData
// 使用预分配账户作为初始验证者
function generateExtraData() {
    // 构建 extraData 结构:
    // 32字节 vanity + N*20字节 validator addresses + 65字节 signature
    const vanity = '0'.repeat(64); // 32字节的vanity (可以是任意数据)
    
    // 验证者地址列表 (去掉0x前缀，保持20字节)
    const validatorAddresses = INITIAL_VALIDATORS.map(addr => addr.slice(2).toLowerCase()).join('');
    
    // 为了生成正确的签名，我们需要：
    // 1. 构建要签名的数据（不包含签名部分）
    const dataToSign = vanity + validatorAddresses;
    
    // 2. 对数据进行哈希
    const hash = keccak256Hash(Buffer.from(dataToSign, 'hex'));
    
    // 3. 生成一个简单的签名（在实际场景中，这应该是由验证者私钥签名）
    // 这里我们使用一个确定性的方法来生成签名
    const messageHash = Buffer.from(hash.slice(2), 'hex');
    
    // 生成一个固定的签名（65字节：32字节r + 32字节s + 1字节v）
    // 注意：这不是真正的ECDSA签名，只是为了格式正确
    const r = keccak256(messageHash).slice(0, 64);
    const s = keccak256(r + 'salt').slice(0, 64);
    const v = '1c'; // recovery id (通常是1b或1c)
    
    const signature = r + s + v;
    
    return '0x' + vanity + validatorAddresses + signature;
}

// 生成 Staking 合约的存储状态
function generateStakingStorage() {
    const crypto = require('crypto');
    
    // 计算存储槽位的辅助函数（使用正确的 keccak256）
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
    
    // Staking 合约存储槽位布局（考虑从 Params 继承）：
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
    
    // 注意：不设置 initialized，让 Congress 共识引擎调用 initialize() 或 initializeWithValidators() 方法
    
    // 设置 validatorsContract 地址 (slot 8)
    const validatorsAddress = CONTRACT_ADDRESSES.Validators.toLowerCase();
    storage['0x0000000000000000000000000000000000000000000000000000000000000008'] = 
        '0x' + validatorsAddress.slice(2).padStart(64, '0');
    
    // 设置 proposalContract 地址 (slot 9)
    const proposalAddress = CONTRACT_ADDRESSES.Proposal.toLowerCase();
    storage['0x0000000000000000000000000000000000000000000000000000000000000009'] = 
        '0x' + proposalAddress.slice(2).padStart(64, '0');
    
    // 为每个验证者设置质押信息
    const minValidatorStake = BigInt('10000000000000000000000'); // 10,000 ether in wei
    let totalStaked = BigInt(0);
    
    INITIAL_VALIDATORS.forEach((validator, index) => {
        const validatorAddr = validator.toLowerCase();
        
        // validatorStakes[validator] mapping (slot 1)
        const stakingSlot = getStorageSlot(1, validatorAddr);
        
        // ValidatorStake 结构体布局：
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
        
        // allValidators array (slot 4) - 设置数组长度
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

// 更新 Genesis 文件
function updateGenesisFile() {
    const genesisPath = path.join(__dirname, '..', 'chain', 'genesis.json');
    
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
                const contractAlloc = {
                    balance: "0x0",
                    code: bytecode
                };
                
                // 为 Staking 合约添加预设存储状态
                if (contractName === 'Staking') {
                    contractAlloc.storage = generateStakingStorage();
                    console.log(`✅ ${contractName}: ${address} (包含 ${INITIAL_VALIDATORS.length} 个预设验证者)`);
                } else {
                    console.log(`✅ ${contractName}: ${address}`);
                }
                
                genesis.alloc[address] = contractAlloc;
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
        console.log(`📄 文件位置: ${path.relative(process.cwd(), genesisPath)}`);
        
        // 显示摘要
        console.log('\n📋 更新摘要:');
        console.log(`🏗️  共识算法: Congress (POA)`);
        console.log(`⏱️  出块间隔: ${genesis.config.congress.period} 秒`);
        console.log(`🔄 验证者更新周期: ${genesis.config.congress.epoch} 块`);
        console.log(`🏪 系统合约: ${Object.keys(CONTRACT_ADDRESSES).length} 个`);
        console.log(`👥 预设验证者: ${INITIAL_VALIDATORS.length} 个`);
        console.log(`💰 每个验证者质押: 10,000 JU`);
        console.log(`🆔 链 ID: ${genesis.config.chainId}`);
        
        console.log('\n👥 初始验证者列表:');
        INITIAL_VALIDATORS.forEach((validator, index) => {
            console.log(`   ${index + 1}. ${validator}`);
        });
        
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
        console.log('   cd ../chain && ./pm2-init.sh');
        console.log('   或者直接使用: cd ../chain && pm2 start ecosystem.config.js');
        console.log('\n📋 重要提示:');
        console.log('   ✅ 创世区块已包含 5 个预设验证者的质押信息');
        console.log('   ✅ 每个验证者已质押 10,000 JU 代币');
        console.log('   ✅ JPoSA 共识将正常工作，无需手动注册验证者');
        console.log('   ✅ 可以直接进行验证者投票和质押操作');
    }
}

// 运行主函数
main();
