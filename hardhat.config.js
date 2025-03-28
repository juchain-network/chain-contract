require('@nomiclabs/hardhat-waffle')
require('dotenv').config()

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task('accounts', 'Prints the list of accounts', async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners()

    for (const account of accounts) {
        console.log(account.address)
    }
})

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    solidity: {
        version: '0.6.1',
        settings: {
            optimizer: {
                enabled: true,
                runs: 200,
            },
        },
    },
    networks: {
        juchain: {
            url: 'https://testnet-rpc.juchain.org/',
            accounts: [
                {
                    privateKey: "0xca881281fb10b53a87d00cbfae29f7cf8cfe8ac7c8389b3d20b24fc6bc3f3ff9",
                    balance: "10000000000000000000" // 10000 ETH
                },
                {
                    privateKey: "0x38addee35ab7fbecc26792602517b5de270938d6933581686f58d5baa0c0cd7e",
                    balance: "10000000000000000000" // 10000 ETH
                },
                {
                    privateKey: "0x85c864aa0828b2b37273b1e080068dd2e33f3e54e233252b829ceff23ea6091a",
                    balance: "10000000000000000000" // 10000 ETH
                }
            ],
            timeout: 50000,
            gas: 7000000
        },
        hardhat: {
            allowUnlimitedContractSize: true,
        },
        aacTest: {
            url: 'http://3.129.253.183:10212',
            accounts: [process.env.MINER],
            timeout: 50000,
            gas: 7000000
        },
        testnet: {
            url:'',
            accounts:[process.env.MINER],
            timeout:50000,
            gas:7000000
        },
        devnet: {
            url:'',
            accounts:[process.env.MINER],
            timeout:50000,
            gas:7000000
        },
        mainnet: {
            url:'',
            accounts:[process.env.MINER],
            timeout:50000,
            gas:7000000
        }
    },
}