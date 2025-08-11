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
    version: '0.8.20',
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
                "0xca881281fb10b53a87d00cbfae29f7cf8cfe8ac7c8389b3d20b24fc6bc3f3ff9",
                "0x38addee35ab7fbecc26792602517b5de270938d6933581686f58d5baa0c0cd7e",
                "0x85c864aa0828b2b37273b1e080068dd2e33f3e54e233252b829ceff23ea6091a"
            ],
            timeout: 50000,
            gas: 7000000
        },
        hardhat: {
            allowUnlimitedContractSize: true,
        },
        localhost: {
            url: 'http://localhost:8545',
            accounts: [
                "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
                "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
            ],
            timeout: 50000,
            gas: 7000000
        },
        aacTest: {
            url: 'http://3.129.253.183:10212',
            accounts: process.env.MINER ? [process.env.MINER] : [],
            timeout: 50000,
            gas: 7000000
        }
    },
}