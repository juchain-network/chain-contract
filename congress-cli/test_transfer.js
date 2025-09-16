const { Web3 } = require('web3');
const web3 = new Web3('http://localhost:8545');

async function transferFunds() {
    const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const toAddress = '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc';
    const amount = web3.utils.toWei('10000', 'ether');
    
    console.log('🚀 尝试转账10000 ETH给验证者6...');
    console.log('From:', fromAddress);
    console.log('To:', toAddress);
    console.log('Amount:', web3.utils.fromWei(amount, 'ether'), 'ETH');
    
    // 检查余额
    const balance = await web3.eth.getBalance(fromAddress);
    console.log('发送方余额:', web3.utils.fromWei(balance, 'ether'), 'ETH');
    
    try {
        // 先检查账户是否解锁
        const accounts = await web3.eth.getAccounts();
        console.log('可用账户数量:', accounts.length);
        
        const txHash = await web3.eth.sendTransaction({
            from: fromAddress,
            to: toAddress,
            value: amount,
            gas: 21000,
            gasPrice: web3.utils.toWei('20', 'gwei')
        });
        
        console.log('✅ 转账成功!');
        console.log('交易哈希:', typeof txHash === 'object' ? txHash.transactionHash : txHash);
        
        const newBalance = await web3.eth.getBalance(toAddress);
        console.log('验证者6新余额:', web3.utils.fromWei(newBalance, 'ether'), 'ETH');
        
    } catch (error) {
        console.error('❌ 转账失败详细信息:');
        console.error('错误消息:', error.message);
        console.error('错误代码:', error.code);
        if (error.data) {
            console.error('错误数据:', error.data);
        }
        if (error.receipt) {
            console.error('交易收据:', error.receipt);
        }
    }
}

transferFunds();