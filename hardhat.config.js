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