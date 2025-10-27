const { Web3 } = require('web3');

// 连接到本地私链
const web3 = new Web3('http://localhost:8545');

// 转账函数
async function transferFunds() {
    try {
        console.log('�� 检查网络连接...');
        const chainId = await web3.eth.getChainId();
        const blockNumber = await web3.eth.getBlockNumber();
        console.log(`✅ 连接成功！链ID: ${chainId}, 区块高度: ${blockNumber}`);

        // 发送方（已有资金的账户）
        const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
        
        // 接收方（validator6和validator7）
        const validator6 = '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc';
        const validator7 = '0x50c554ac9c134491818fa6f21d504f2ae5bd9c26';
        
        // 转账金额（20000 ETH）
        const amount = web3.utils.toWei('20000', 'ether');
        const gasPrice = web3.utils.toWei('20', 'gwei');
        
        // 检查发送方余额
        const senderBalance = await web3.eth.getBalance(fromAddress);
        console.log(`📊 发送方余额: ${web3.utils.fromWei(senderBalance, 'ether')} ETH`);
        
        // 转账给validator6
        console.log('🚀 转账20000 ETH给validator6...');
        const tx1 = await web3.eth.sendTransaction({
            from: fromAddress,
            to: validator6,
            value: amount,
            gas: 21000,
            gasPrice: gasPrice
        });
        console.log(`✅ 转账成功！交易哈希: ${tx1.transactionHash}`);
        
        // 转账给validator7
        console.log('🚀 转账20000 ETH给validator7...');
        const tx2 = await web3.eth.sendTransaction({
            from: fromAddress,
            to: validator7,
            value: amount,
            gas: 21000,
            gasPrice: gasPrice
        });
        console.log(`✅ 转账成功！交易哈希: ${tx2.transactionHash}`);
        
        // 检查转账后的余额
        const validator6Balance = await web3.eth.getBalance(validator6);
        const validator7Balance = await web3.eth.getBalance(validator7);
        
        console.log('📊 转账后余额:');
        console.log(`   Validator6: ${web3.utils.fromWei(validator6Balance, 'ether')} ETH`);
        console.log(`   Validator7: ${web3.utils.fromWei(validator7Balance, 'ether')} ETH`);
        
    } catch (error) {
        console.error('❌ 转账失败:', error.message);
        if (error.message.includes('insufficient funds')) {
            console.log('💡 建议: 确保发送方账户有足够的余额');
        }
        if (error.message.includes('connection')) {
            console.log('💡 建议: 确保私链正在运行 (geth 或 anvil)');
        }
    }
}

// 执行转账
transferFunds();
