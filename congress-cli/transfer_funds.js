const { Web3 } = require('web3');

// Connect to local private chain
const web3 = new Web3('http://localhost:8545');

// Transfer function
async function transferFunds() {
    try {
        console.log('🔍 Checking network connection...');
        const chainId = await web3.eth.getChainId();
        const blockNumber = await web3.eth.getBlockNumber();
        console.log(`✅ Connection successful! Chain ID: ${chainId}, Block height: ${blockNumber}`);

        // Sender (account with existing funds)
        const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
        
        // Receiver (validator6 and validator7)
        const validator6 = '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc';
        const validator7 = '0x50c554ac9c134491818fa6f21d504f2ae5bd9c26';
        
        // Transfer amount (20000 ETH)
        const amount = web3.utils.toWei('20000', 'ether');
        const gasPrice = web3.utils.toWei('20', 'gwei');
        
        // Check sender balance
        const senderBalance = await web3.eth.getBalance(fromAddress);
        console.log(`📊 Sender balance: ${web3.utils.fromWei(senderBalance, 'ether')} ETH`);
        
        // Transfer to validator6
        console.log('🚀 Transferring 20000 ETH to validator6...');
        const tx1 = await web3.eth.sendTransaction({
            from: fromAddress,
            to: validator6,
            value: amount,
            gas: 21000,
            gasPrice: gasPrice
        });
        console.log(`✅ Transfer successful! Transaction hash: ${tx1.transactionHash}`);
        
        // Transfer to validator7
        console.log('🚀 Transferring 20000 ETH to validator7...');
        const tx2 = await web3.eth.sendTransaction({
            from: fromAddress,
            to: validator7,
            value: amount,
            gas: 21000,
            gasPrice: gasPrice
        });
        console.log(`✅ Transfer successful! Transaction hash: ${tx2.transactionHash}`);
        
        // Check balances after transfer
        const validator6Balance = await web3.eth.getBalance(validator6);
        const validator7Balance = await web3.eth.getBalance(validator7);
        
        console.log('📊 Balances after transfer:');
        console.log(`   Validator6: ${web3.utils.fromWei(validator6Balance, 'ether')} ETH`);
        console.log(`   Validator7: ${web3.utils.fromWei(validator7Balance, 'ether')} ETH`);
        
    } catch (error) {
        console.error('❌ Transfer failed:', error.message);
        if (error.message.includes('insufficient funds')) {
            console.log('💡 Suggestion: Ensure sender account has sufficient balance');
        }
        if (error.message.includes('connection')) {
            console.log('💡 Suggestion: Ensure private chain is running (geth or anvil)');
        }
    }
}

// Execute transfer
transferFunds();
