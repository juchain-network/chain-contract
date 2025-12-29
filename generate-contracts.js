const fs = require("fs");
const nunjucks = require("nunjucks");

// List of files to process
const filesToProcess = {
    // Production contract
    production: {
        src: "contracts/Params.template",
        dst: "contracts/Params.sol",
        config: { mock: false }
    },
    
    // Test contract
    test: {
        src: "contracts/Params.template",
        dst: "script/integration/generated/Params.sol",
        config: { mock: true }
    },
    
    // Files to copy to test directory
    toCopy: [
        "contracts/Validators.sol",
        "contracts/Punish.sol",
        "contracts/Proposal.sol",
        "contracts/Staking.sol",
        "contracts/IValidators.sol",
        "contracts/IPunish.sol",
        "contracts/IStaking.sol",
        "contracts/IProposal.sol"
    ]
};

// Create test directory if it doesn't exist
const testDir = "script/integration/generated";
if (!fs.existsSync(testDir)) {
    fs.mkdirSync(testDir, { recursive: true });
}

console.log("Starting contract generation...");

// 1. Generate production Params.sol
console.log("\n1. Generating production Params.sol...");
const productionTemplate = fs.readFileSync(filesToProcess.production.src).toString();
const productionContract = nunjucks.renderString(productionTemplate, filesToProcess.production.config);
fs.writeFileSync(filesToProcess.production.dst, productionContract);
console.log(`   ✓ Generated: ${filesToProcess.production.dst}`);

// 2. Generate test Params.sol
console.log("\n2. Generating test Params.sol...");
const testTemplate = fs.readFileSync(filesToProcess.test.src).toString();
const testContract = nunjucks.renderString(testTemplate, filesToProcess.test.config);
fs.writeFileSync(filesToProcess.test.dst, testContract);
console.log(`   ✓ Generated: ${filesToProcess.test.dst}`);

// 3. Copy system contracts and interfaces to test directory
console.log("\n3. Copying system contracts and interfaces...");
for (const file of filesToProcess.toCopy) {
    const fileName = file.split("/").pop();
    const dst = `${testDir}/${fileName}`;
    const content = fs.readFileSync(file).toString();
    fs.writeFileSync(dst, content);
    console.log(`   ✓ Copied: ${dst}`);
}

console.log("\n✓ All contracts generated successfully!");