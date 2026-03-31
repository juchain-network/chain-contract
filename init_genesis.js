#!/usr/bin/env node

// Extract system contract bytecode and update Genesis file
const fs = require("fs");
const path = require("path");
const { keccak256 } = require("js-sha3");

console.log(
  "🔧 Extracting system contract bytecode and updating Genesis file...",
);

// Contract address mappings
const CONTRACT_ADDRESSES = {
  Validators: "0x000000000000000000000000000000000000f010",
  Punish: "0x000000000000000000000000000000000000f011",
  Proposal: "0x000000000000000000000000000000000000f012",
  Staking: "0x000000000000000000000000000000000000f013",
};

// Initial validator information for the local genesis template.
const INITIAL_VALIDATORS = ["0x70997970C51812dc3A010C7d01b50e0d17dc79C8"];

function normalizeAddress(addr) {
  return `0x${addr.slice(2).toLowerCase()}`;
}

function ensureBootstrapConfig(genesis) {
  if (!genesis.config) {
    genesis.config = {};
  }
  if (!genesis.config.congress) {
    genesis.config.congress = {};
  }

  const normalizedValidators = INITIAL_VALIDATORS.map(normalizeAddress);
  genesis.config.congress.initialValidators = normalizedValidators;
  genesis.config.congress.initialSigners = normalizedValidators;
}

function ensureBootstrapBalances(genesis, minValidatorStake) {
  const normalizedValidators = INITIAL_VALIDATORS.map(normalizeAddress);

  for (const validator of normalizedValidators) {
    const currentBalanceHex = genesis.alloc[validator]?.balance || "0x0";
    const currentBalance = BigInt(currentBalanceHex);

    if (currentBalance < minValidatorStake) {
      genesis.alloc[validator] = {
        ...(genesis.alloc[validator] || {}),
        balance: `0x${minValidatorStake.toString(16)}`,
      };
      console.log(
        `✅ Ensured bootstrap balance for ${validator}: ${formatJu(minValidatorStake)}`,
      );
    }
  }
}

function readDefaultMinValidatorStake() {
  const proposalSourcePath = path.join(__dirname, "contracts", "Proposal.sol");
  const proposalSource = fs.readFileSync(proposalSourcePath, "utf8");
  const match = proposalSource.match(
    /DEFAULT_MIN_VALIDATOR_STAKE\s*=\s*([0-9_]+)\s*ether/,
  );
  if (!match) {
    throw new Error("Unable to parse DEFAULT_MIN_VALIDATOR_STAKE from Proposal.sol");
  }
  return BigInt(match[1].replaceAll("_", "")) * 10n ** 18n;
}

function formatJu(wei) {
  return `${wei / 10n ** 18n} JU`;
}

// Read contract bytecode
function getContractBytecode(contractName) {
  // Try Foundry first (out directory)
  const foundryPath = path.join(
    __dirname,
    "out",
    `${contractName}.sol`,
    `${contractName}.json`,
  );
  try {
    const artifact = JSON.parse(fs.readFileSync(foundryPath, "utf8"));
    return artifact.deployedBytecode?.object || artifact.bytecode?.object;
  } catch (foundryError) {
    // Fallback to Hardhat artifacts
    const artifactPath = path.join(
      __dirname,
      "..",
      "artifacts",
      "contracts",
      `${contractName}.sol`,
      `${contractName}.json`,
    );
    try {
      const artifact = JSON.parse(fs.readFileSync(artifactPath, "utf8"));
      return artifact.deployedBytecode;
    } catch (error) {
      console.error(
        `❌ Failed to read ${contractName} contract bytecode:`,
        error.message,
      );
      return null;
    }
  }
}

// keccak256 hash helper function
function keccak256Hash(data) {
  if (typeof data === "string" && data.startsWith("0x")) {
    return "0x" + keccak256(Buffer.from(data.slice(2), "hex"));
  }
  return "0x" + keccak256(data);
}

// Generate initial validators' extraData
// Using pre-allocated accounts as initial validators
function generateExtraData() {
  // Build extraData structure:
  // 32 bytes vanity + N*20 bytes validator addresses + 65 bytes signature
  const vanity = "0".repeat(64); // 32 bytes vanity (can be arbitrary data)

  // Validator address list (remove 0x prefix, keep 20 bytes)
  const validatorAddresses = INITIAL_VALIDATORS.map((addr) =>
    addr.slice(2).toLowerCase(),
  ).join("");

  // To generate the correct signature, we need:
  // 1. Build data to sign (excluding signature part)
  const dataToSign = vanity + validatorAddresses;

  // 2. Hash the data
  const hash = keccak256Hash(Buffer.from(dataToSign, "hex"));

  // 3. Generate a simple signature (in real scenarios, this should be signed by validator private keys)
  // Here we use a deterministic method to generate the signature
  const messageHash = Buffer.from(hash.slice(2), "hex");

  // Generate a fixed signature (65 bytes: 32 bytes r + 32 bytes s + 1 byte v)
  // Note: This is not a real ECDSA signature, just for correct format
  const r = keccak256(messageHash).slice(0, 64);
  const s = keccak256(r + "salt").slice(0, 64);
  const v = "1c"; // recovery id (usually 1b or 1c)

  const signature = r + s + v;

  return "0x" + vanity + validatorAddresses + signature;
}


// Update Genesis file
function updateGenesisFile() {
  // const genesisPath = path.join(__dirname, '..', 'chain', 'genesis.json');
  const genesisPath = path.join(__dirname, "genesis.json");
  const minValidatorStake = readDefaultMinValidatorStake();

  try {
    // Read existing Genesis file
    const genesis = JSON.parse(fs.readFileSync(genesisPath, "utf8"));

    // Ensure alloc field exists
    if (!genesis.alloc) {
      genesis.alloc = {};
    }

    ensureBootstrapConfig(genesis);
    ensureBootstrapBalances(genesis, minValidatorStake);

    // Add system contracts
    console.log("📋 Adding system contracts to Genesis file...");

    for (const [contractName, address] of Object.entries(CONTRACT_ADDRESSES)) {
      const bytecode = getContractBytecode(contractName);
      if (bytecode) {
        const contractAlloc = {
          balance: "0x0",
          code: bytecode,
        };

        // Add preset storage state for Staking contract
        if (contractName === "Staking") {
          // contractAlloc.storage = generateStakingStorage();
          console.log(
            `✅ ${contractName}: ${address} (includes ${INITIAL_VALIDATORS.length} preset validators)`,
          );
        } else {
          console.log(`✅ ${contractName}: ${address}`);
        }

        genesis.alloc[address] = contractAlloc;
      } else {
        console.log(`❌ ${contractName}: Failed to get bytecode`);
      }
    }

    // Update extraData to include initial validators
    genesis.extraData = generateExtraData();
    console.log("✅ Updated extraData to include initial validators");
    console.log("✅ Updated config.congress.initialValidators / initialSigners");

    // Write back to Genesis file
    fs.writeFileSync(genesisPath, JSON.stringify(genesis, null, 2));
    console.log("✅ Genesis file updated successfully!");
    console.log(
      `📄 File location: ${path.relative(process.cwd(), genesisPath)}`,
    );

    // Display summary
    console.log("\n📋 Update Summary:");
    console.log(`🏗️  Consensus Algorithm: Congress (POA)`);
    console.log(
      `⏱️  Block Interval: ${genesis.config.congress.period} seconds`,
    );
    console.log(
      `🔄 Validator Update Cycle: ${genesis.config.congress.epoch} blocks`,
    );
    console.log(
      `🏪 System Contracts: ${Object.keys(CONTRACT_ADDRESSES).length} contracts`,
    );
    console.log(
      `👥 Preset Validators: ${INITIAL_VALIDATORS.length} validators`,
    );
    console.log(`💰 Bootstrap Stake Floor: ${formatJu(minValidatorStake)}`);
    console.log(`🆔 Chain ID: ${genesis.config.chainId}`);

    console.log("\n👥 Initial Validator List:");
    INITIAL_VALIDATORS.forEach((validator, index) => {
      console.log(`   ${index + 1}. ${validator}`);
    });
  } catch (error) {
    console.error("❌ Failed to update Genesis file:", error.message);
    process.exit(1);
  }
}

// Verify contract compilation status
function verifyContracts() {
  console.log("🔍 Verifying contract compilation status...");

  let allContractsReady = true;
  for (const contractName of Object.keys(CONTRACT_ADDRESSES)) {
    const bytecode = getContractBytecode(contractName);
    if (!bytecode || bytecode === "0x") {
      console.log(`❌ ${contractName}: Not compiled or bytecode is empty`);
      allContractsReady = false;
    } else {
      console.log(
        `✅ ${contractName}: Compilation successful (${bytecode.length} characters)`,
      );
    }
  }

  if (!allContractsReady) {
    console.log(
      "\n❌ Please compile contracts first: forge build or npx hardhat compile",
    );
    process.exit(1);
  }

  return true;
}

// Main function
function main() {
  console.log("🚀 Congress Consensus Configuration Tool\n");
  const minValidatorStake = readDefaultMinValidatorStake();

  // Verify contract compilation status
  if (verifyContracts()) {
    // Update Genesis file
    updateGenesisFile();

    console.log("\n🎉 Congress Consensus Configuration Complete!");
    console.log("💡 Next steps to start private chain:");
    console.log("   cd ../chain && ./pm2-init.sh");
    console.log(
      "   Or directly use: cd ../chain && pm2 start ecosystem.config.js",
    );
    console.log("\n📋 Important Notes:");
    console.log(
      `   ✅ extraData includes ${INITIAL_VALIDATORS.length} preset validator(s) for local genesis`,
    );
    console.log(
      `   ✅ Each genesis validator cold address should hold at least ${formatJu(minValidatorStake)} so bootstrap can move that amount into Staking at initialization`,
    );
    console.log(
      "   ✅ JPoSA consensus will work normally, no manual validator registration needed",
    );
    console.log(
      "   ✅ Validator voting and staking operations can be performed directly",
    );
    console.log(
      "   ✅ Contract-side validator validity is gated by minValidatorStake",
    );
  }
}

// Run main function
main();
