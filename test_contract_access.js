#!/usr/bin/env node

// 实际访问和测试合约功能
const fs = require('fs');
const { spawn } = require('child_process');

console.log('🚀 启动合约功能测试...\n');

// 启动一个轻量级测试节点
function startTestNode() {
    return new Promise((resolve, reject) => {
        console.log('📡 启动测试节点...');
        
        const geth = spawn('geth', [
            '--datadir', './private-chain/data-test',
            'init', './genesis.json'
        ], {
            cwd: '/Users/enty/ju-chain-work/chain',
            stdio: ['pipe', 'pipe', 'pipe']
        });

        let initComplete = false;
        
        geth.on('close', (code) => {
            if (code === 0) {
                console.log('✅ 创世区块初始化完成');
                initComplete = true;
                
                // 启动节点
                const gethNode = spawn('geth', [
                    '--datadir', './private-chain/data-test',
                    '--networkid', '202599',
                    '--port', '30399',
                    '--http', '--http.port', '8599',
                    '--http.api', 'eth,net,web3,personal,admin,miner',
                    '--http.corsdomain', '*',
                    '--allow-insecure-unlock',
                    '--nodiscover',
                    '--maxpeers', '0',
                    '--verbosity', '3'
                ], {
                    cwd: '/Users/enty/ju-chain-work/chain',
                    stdio: ['pipe', 'pipe', 'pipe']
                });

                let nodeReady = false;
                
                gethNode.stdout.on('data', (data) => {
                    const output = data.toString();
                    console.log('节点输出:', output.slice(0, 200) + (output.length > 200 ? '...' : ''));
                    
                    if (output.includes('HTTP server started') || output.includes('IPC endpoint opened')) {
                        if (!nodeReady) {
                            nodeReady = true;
                            setTimeout(() => resolve(gethNode), 3000);
                        }
                    }
                });

                gethNode.stderr.on('data', (data) => {
                    console.log('节点错误:', data.toString());
                });

                gethNode.on('close', (code) => {
                    console.log(`节点进程退出，代码: ${code}`);
                    if (!nodeReady) {
                        reject(new Error(`节点启动失败，退出代码: ${code}`));
                    }
                });

                // 15 秒超时
                setTimeout(() => {
                    if (!nodeReady) {
                        gethNode.kill();
                        reject(new Error('节点启动超时'));
                    }
                }, 15000);
                
            } else {
                reject(new Error(`创世区块初始化失败，退出代码: ${code}`));
            }
        });
    });
}

// 通过 RPC 调用合约方法
async function callContract(method, params = []) {
    const Web3 = require('web3');
    const web3 = new Web3('http://localhost:8599');
    
    try {
        const result = await web3.eth.call({
            to: '0x000000000000000000000000000000000000f003', // Staking 合约地址
            data: web3.eth.abi.encodeFunctionCall({
                name: method,
                type: 'function',
                inputs: params.map(p => ({ type: p.type, name: p.name }))
            }, params.map(p => p.value))
        });
        
        return result;
    } catch (error) {
        console.error(`调用 ${method} 失败:`, error.message);
        return null;
    }
}

// 测试 Staking 合约的关键方法
async function testStakingContract() {
    console.log('\n🧪 测试 Staking 合约方法...\n');
    
    const Web3 = require('web3');
    const web3 = new Web3('http://localhost:8599');
    
    try {
        // 测试基本信息
        console.log('1. 📊 获取基本信息:');
        
        // totalStaked
        const totalStakedCall = {
            to: '0x000000000000000000000000000000000000f003',
            data: '0x817b1cd2' // totalStaked() 方法签名
        };
        
        const totalStaked = await web3.eth.call(totalStakedCall);
        console.log('   总质押量:', web3.utils.fromWei(web3.utils.hexToNumberString(totalStaked)), 'JU');
        
        // 测试验证者信息
        console.log('\n2. 👥 验证者信息:');
        const validators = [
            '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
            '0x70997970C51812dc3A010C7d01b50e0d17dc79C8', 
            '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC'
        ];
        
        for (let i = 0; i < Math.min(validators.length, 3); i++) {
            const validator = validators[i];
            console.log(`\n   验证者 ${i + 1}: ${validator}`);
            
            // getValidatorInfo(address)
            const validatorInfoCall = {
                to: '0x000000000000000000000000000000000000f003',
                data: '0x8a11d7c9' + validator.slice(2).padStart(64, '0') // getValidatorInfo(address)
            };
            
            try {
                const result = await web3.eth.call(validatorInfoCall);
                if (result && result !== '0x') {
                    // 解析返回的数据 (selfStake, totalDelegated, commissionRate, isJailed, jailUntilBlock)
                    const selfStake = web3.utils.hexToNumberString('0x' + result.slice(2, 66));
                    const totalDelegated = web3.utils.hexToNumberString('0x' + result.slice(66, 130));
                    const commissionRate = web3.utils.hexToNumberString('0x' + result.slice(130, 194));
                    
                    console.log(`     ✅ 自质押: ${web3.utils.fromWei(selfStake)} JU`);
                    console.log(`     📊 总委托: ${web3.utils.fromWei(totalDelegated)} JU`);
                    console.log(`     💵 佣金率: ${commissionRate / 100}%`);
                } else {
                    console.log('     ❌ 未找到验证者信息');
                }
            } catch (error) {
                console.log(`     ❌ 获取信息失败: ${error.message}`);
            }
        }
        
        // 测试 getTopValidators
        console.log('\n3. 🏆 获取顶级验证者:');
        try {
            const topValidatorsCall = {
                to: '0x000000000000000000000000000000000000f003',
                data: '0xaa7355780000000000000000000000000000000000000000000000000000000000000015' // getTopValidators(21)
            };
            
            const topValidators = await web3.eth.call(topValidatorsCall);
            console.log('   getTopValidators 调用结果:', topValidators);
            
            if (topValidators && topValidators !== '0x' && topValidators.length > 2) {
                // 解析地址数组
                console.log('   ✅ 成功获取顶级验证者列表');
            } else {
                console.log('   ❌ getTopValidators 返回空结果');
            }
        } catch (error) {
            console.log(`   ❌ getTopValidators 调用失败: ${error.message}`);
        }
        
    } catch (error) {
        console.error('测试过程中发生错误:', error);
    }
}

// 主测试流程
async function runTests() {
    let gethProcess = null;
    
    try {
        // 清理旧数据
        console.log('🧹 清理测试数据...');
        if (fs.existsSync('/Users/enty/ju-chain-work/chain/private-chain/data-test')) {
            fs.rmSync('/Users/enty/ju-chain-work/chain/private-chain/data-test', { recursive: true });
        }
        
        // 启动测试节点
        gethProcess = await startTestNode();
        console.log('✅ 测试节点启动成功\n');
        
        // 等待节点稳定
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // 运行合约测试
        await testStakingContract();
        
    } catch (error) {
        console.error('❌ 测试失败:', error.message);
    } finally {
        // 清理
        if (gethProcess) {
            console.log('\n🧹 清理测试节点...');
            gethProcess.kill();
            
            // 等待进程完全结束
            await new Promise(resolve => setTimeout(resolve, 2000));
        }
        
        console.log('\n🎯 测试完成!');
    }
}

// 检查依赖
function checkDependencies() {
    try {
        require('web3');
        return true;
    } catch (error) {
        console.log('❌ 缺少 web3 依赖，正在安装...');
        return false;
    }
}

// 主入口
async function main() {
    if (!checkDependencies()) {
        const { spawn } = require('child_process');
        
        console.log('📦 安装 web3...');
        const npm = spawn('npm', ['install', 'web3'], {
            cwd: '/Users/enty/ju-chain-work/sys-contract',
            stdio: 'inherit'
        });
        
        npm.on('close', (code) => {
            if (code === 0) {
                console.log('✅ web3 安装成功，重新启动测试...');
                delete require.cache[require.resolve('web3')];
                runTests();
            } else {
                console.error('❌ web3 安装失败');
            }
        });
    } else {
        await runTests();
    }
}

main().catch(console.error);
