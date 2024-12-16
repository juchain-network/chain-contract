// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

async function main() {
    // Hardhat always runs the compile task when running scripts with its command
    // line interface.
    //
    // If this script is run directly using `node` you may want to call compile
    // manually to make sure everything is compiled
    // await hre.run('compile');

    // We get the contract to deploy
    let signers = await ethers.getSigners();
    let miner1 = signers[0];
    let miner2 = signers[1];
    let miner3 = signers[2];
    // 新加入矿工的地址
    let toAdd = "0x50f98e9e9dd4725e17b2f93ae7abdda0bc5718aa";
    const Proposal = await hre.ethers.getContractFactory("Proposal");
    const proposal = await Proposal.attach("0x000000000000000000000000000000000000F002");

    // 矿工创建提案
    let tx = await proposal.connect(miner1).createProposal(toAdd, true, "test proposal");
    let receipt = await tx.wait();
    console.log("miner1 create proposal tx:", tx.hash);
    let ev = receipt.events.find(event => event.event == "LogCreateProposal");
    const [id, proposer, dst, flag, t] = ev.args;
    console.log("proposal-id:", id, "proposer:", proposer, "dst:", dst, "flag", flag, "time:", t)

    // 矿工1、2、3对提案进行投票
    tx = await proposal.connect(miner1).voteProposal(id, true);
    console.log("miner1 vote tx:", tx.hash);
    tx = await proposal.connect(miner2).voteProposal(id, true);
    console.log("miner2 vote tx:", tx.hash);
    tx = await proposal.connect(miner3).voteProposal(id, true);
    console.log("miner3 vote tx:", tx.hash);
    await tx.wait();

    // get current top validators
    const Validators = await hre.ethers.getContractFactory("Validators");
    const validators = await Validators.attach("0x000000000000000000000000000000000000f000");
    let top = await validators.getTopValidators();
    console.log("current top validators:", top);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });