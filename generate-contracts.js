const fs = require("fs");
const nunjucks = require("nunjucks");

// Get configuration from command line arguments or environment variables
const isMock = process.argv.includes('--mock') || process.env.MOCK === 'true';
const isClean = process.argv.includes('--clean');

const config = {
    mock: isMock,
}

// List of contracts to generate
var list = [
    { src: "contracts/Params.template", dst: "contracts/Params.sol" },
    // Add other contracts that need to be generated from templates here
];

// Display current mode
console.log(`Generating contracts in ${isMock ? 'MOCK' : 'PRODUCTION'} mode...`);

// Cleanup generated contracts if --clean flag is provided
if (isClean) {
    console.log("Cleaning up generated contracts...");
    // For now, we'll just regenerate production contracts
    // In future, we could add more sophisticated cleanup logic
    config.mock = false;
}

for (let i = 0; i < list.length; i++) {
    const templateStr = fs.readFileSync(list[i].src).toString();
    const contractStr = nunjucks.renderString(templateStr, config);
    fs.writeFileSync(list[i].dst, contractStr);
    console.log(`Generated: ${list[i].dst} (${config.mock ? 'mock' : 'production'} version)`);
}

console.log(`Generate ${config.mock ? 'mock' : 'system'} contracts success`);

// Usage instructions
if (process.argv.includes('--help') || process.argv.includes('-h')) {
    console.log(`
Usage:
  node generate-contracts.js [--mock] [--clean] [--help]

Options:
  --mock    Generate mock contracts for testing (default: production contracts)
  --clean   Cleanup generated contracts (regenerate production contracts)
  --help    Show this help message

Environment Variables:
  MOCK=true    Same as --mock flag

Examples:
  node generate-contracts.js              # Generate production contracts
  node generate-contracts.js --mock       # Generate mock contracts
  node generate-contracts.js --clean      # Regenerate production contracts
  MOCK=true node generate-contracts.js    # Generate mock contracts via env var
`);
}