const hre = require("hardhat");

async function main() {
    // We get the contract to deploy
    let signers = await ethers.getSigners();
    let miner = signers[0];
    // new miner address
    let toAdd = "0x50f98e9e9dd4725e17b2f93ae7abdda0bc5718aa";
    //if user input address used this
    let address = process.env.ConstructorArguments;
    if(address){
        toAdd = address;
    }
    let reg = /^0x[A-Z|a-z|0-9]{40}$/;
    let valid_address = reg.test(toAdd);
    if(!valid_address){
        console.log("you input address is error");
        return;
    }
    const Proposal = await hre.ethers.getContractFactory("Proposal");
    const proposal = await Proposal.attach("0x000000000000000000000000000000000000F002");

    // 矿工创建提案
    let tx = await proposal.connect(miner).createProposal(toAdd, true, "create a proposal");
    let receipt = await tx.wait();
    console.log("miner create proposal tx:", tx.hash);
    let ev = receipt.events.find(event => event.event == "LogCreateProposal");
    const [id, proposer, dst, flag, t] = ev.args;
    console.log("proposal-id:", id, "proposer:", proposer, "dst:", dst, "flag", flag, "time:", t)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });