const { Web3 } = require('web3');
const web3 = new Web3('http://localhost:8545');

async function transferFunds() {
    const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const toAddress = '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc';
    const amount = web3.utils.toWei('10000', 'ether');
    
    console.log('🚀 Attempting to transfer 10000 ETH to validator6...');
    console.log('From:', fromAddress);
    console.log('To:', toAddress);
    console.log('Amount:', web3.utils.fromWei(amount, 'ether'), 'ETH');
    
    // Check balance
    const balance = await web3.eth.getBalance(fromAddress);
    console.log('Sender balance:', web3.utils.fromWei(balance, 'ether'), 'ETH');
    
    try {
        // First check if account is unlocked
        const accounts = await web3.eth.getAccounts();
        console.log('Available accounts count:', accounts.length);
        
        const txHash = await web3.eth.sendTransaction({
            from: fromAddress,
            to: toAddress,
            value: amount,
            gas: 21000,
            gasPrice: web3.utils.toWei('20', 'gwei')
        });
        
        console.log('✅ Transfer successful!');
        console.log('Transaction hash:', typeof txHash === 'object' ? txHash.transactionHash : txHash);
        
        const newBalance = await web3.eth.getBalance(toAddress);
        console.log('Validator6 new balance:', web3.utils.fromWei(newBalance, 'ether'), 'ETH');
        
    } catch (error) {
        console.error('❌ Transfer failed details:');
        console.error('Error message:', error.message);
        console.error('Error code:', error.code);
        if (error.data) {
            console.error('Error data:', error.data);
        }
        if (error.receipt) {
            console.error('Transaction receipt:', error.receipt);
        }
    }
}

transferFunds();