const fs = require('fs');
const path = require('path');

// Contract address mappings (Must match init_genesis.js)
const CONTRACT_ADDRESSES = {
  Validators: "0x000000000000000000000000000000000000f010",
  Punish: "0x000000000000000000000000000000000000f011",
  Proposal: "0x000000000000000000000000000000000000f012",
  Staking: "0x000000000000000000000000000000000000f013",
};

// Project root directory (2 levels up from this script: test-integration/scripts/build_alloc.js)
const PROJECT_ROOT = path.resolve(__dirname, '../..');
const OUT_DIR = path.join(PROJECT_ROOT, 'out');

function getContractBytecode(contractName) {
  const artifactPath = path.join(OUT_DIR, `${contractName}.sol`, `${contractName}.json`);
  try {
    if (!fs.existsSync(artifactPath)) {
      console.error(`Artifact not found: ${artifactPath}`);
      return null;
    }
    const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
    // Use deployedBytecode (runtime bytecode) for genesis alloc
    return artifact.deployedBytecode?.object || artifact.deployedBytecode;
  } catch (error) {
    console.error(`Failed to read bytecode for ${contractName}: ${error.message}`);
    return null;
  }
}

function main() {
  const alloc = {};

  for (const [contractName, address] of Object.entries(CONTRACT_ADDRESSES)) {
    let bytecode = getContractBytecode(contractName);
    
    if (!bytecode) {
      console.error(`Error: Could not find bytecode for ${contractName}. Did you run 'forge build'?`);
      process.exit(1);
    }

    // Ensure hex prefix
    if (!bytecode.startsWith('0x')) {
      bytecode = '0x' + bytecode;
    }

    alloc[address] = {
      balance: "0x0",
      code: bytecode
    };
  }

  // Output JSON string to stdout
  console.log(JSON.stringify(alloc, null, 2));
}

main();
