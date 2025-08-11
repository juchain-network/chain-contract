const Validators = artifacts.require("Validators");
const Punish = artifacts.require("Punish");
const Proposal = artifacts.require("Proposal");

module.exports = function(deployer) {
  // 部署系统合约
  deployer.deploy(Validators)
    .then(() => deployer.deploy(Punish))
    .then(() => deployer.deploy(Proposal))
    .then(() => {
      console.log("系统合约部署完成:");
      console.log("Validators:", Validators.address);
      console.log("Punish:", Punish.address);
      console.log("Proposal:", Proposal.address);
    });
};
