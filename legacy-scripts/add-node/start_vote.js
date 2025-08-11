const hre = require("hardhat");

async function main() {
    let signers = await ethers.getSigners();
    let miner = signers[0];
  
    //get proposal id from arg
    let proposal_id = process.env.ConstructorArguments;
    if(!proposal_id){
        console.log("please input proposal id");
    }
    let reg = /^0x[A-Z|a-z|0-9]{64}$/;
    let valid_proposal_id = reg.test(proposal_id);
    if(!valid_proposal_id){
        console.log("please input right proposal id");
        return;
    }

    const Proposal = await hre.ethers.getContractFactory("Proposal");
    const proposal = await Proposal.attach("0x000000000000000000000000000000000000F002");

    // start a vote
    tx = await proposal.connect(miner).voteProposal(proposal_id, true);
    console.log("miner1 vote tx:", tx.hash);
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