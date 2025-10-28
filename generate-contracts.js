const fs = require("fs");
const nunjucks = require("nunjucks");

// Get configuration from command line arguments or environment variables
const isMock = process.argv.includes('--mock') || process.env.MOCK === 'true';

const config = {
    mock: isMock,
}

var list = [
    { src: "contracts/Params.template", dst: "contracts/Params.sol" },
];

// Display current mode
console.log(`Generating contracts in ${isMock ? 'MOCK' : 'PRODUCTION'} mode...`);

for (let i = 0; i < list.length; i++) {
    const templateStr = fs.readFileSync(list[i].src).toString();
    const contractStr = nunjucks.renderString(templateStr, config);
    fs.writeFileSync(list[i].dst, contractStr);
    console.log(`Generated: ${list[i].dst} (${isMock ? 'mock' : 'production'} version)`);
}

console.log(`Generate ${isMock ? 'mock' : 'system'} contracts success`);

// Usage instructions
if (process.argv.includes('--help') || process.argv.includes('-h')) {
    console.log(`
Usage:
  node generate-contracts.js [--mock] [--help]

Options:
  --mock    Generate mock contracts for testing (default: production contracts)
  --help    Show this help message

Environment Variables:
  MOCK=true    Same as --mock flag

Examples:
  node generate-contracts.js              # Generate production contracts
  node generate-contracts.js --mock       # Generate mock contracts
  MOCK=true node generate-contracts.js    # Generate mock contracts via env var
`);
}