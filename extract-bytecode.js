#!/usr/bin/env node

/**
 * Extract bytecode from compiled contracts and write to consensus bytecode directory
 * 
 * Usage:
 *   node extract-bytecode.js
 *   npm run extract-bytecode
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

// Contract mapping: contract name -> output filename
const CONTRACTS = {
  'Validators': 'Validators.hex',
  'Proposal': 'Proposal.hex',
  'Punish': 'Punish.hex',
  'Staking': 'Staking.hex',
};

// Paths
const CONTRACT_DIR = __dirname;
const OUT_DIR = path.join(CONTRACT_DIR, 'out');
const BYTECODE_TARGET_DIR = path.join(CONTRACT_DIR, '..', 'chain.gitlab', 'consensus', 'congress', 'bytecode');

/**
 * Extract bytecode from compiled contract JSON
 */
function extractBytecode(contractName) {
  const jsonPath = path.join(OUT_DIR, `${contractName}.sol`, `${contractName}.json`);
  
  if (!fs.existsSync(jsonPath)) {
    throw new Error(`Contract JSON not found: ${jsonPath}`);
  }

  const json = JSON.parse(fs.readFileSync(jsonPath, 'utf8'));
  
  // Try deployedBytecode first (runtime bytecode), fallback to bytecode (creation bytecode)
  let bytecode = json.deployedBytecode?.object || json.bytecode?.object;
  
  if (!bytecode) {
    // Fallback to direct fields
    bytecode = json.deployedBytecode || json.bytecode;
  }

  if (!bytecode || typeof bytecode !== 'string') {
    throw new Error(`Invalid bytecode format for ${contractName}`);
  }

  // Remove 0x prefix if present
  if (bytecode.startsWith('0x')) {
    bytecode = bytecode.substring(2);
  }

  return bytecode;
}

/**
 * Write bytecode to hex file
 */
function writeBytecodeFile(filename, bytecode) {
  const targetPath = path.join(BYTECODE_TARGET_DIR, filename);
  
  // Create directory if it doesn't exist
  if (!fs.existsSync(BYTECODE_TARGET_DIR)) {
    fs.mkdirSync(BYTECODE_TARGET_DIR, { recursive: true });
  }

  // Write bytecode (hex string without 0x prefix)
  fs.writeFileSync(targetPath, bytecode, 'utf8');
  
  console.log(`✓ Wrote ${filename} (${bytecode.length / 2} bytes)`);
}

/**
 * Get git version (commit hash)
 */
function getGitVersion() {
  try {
    const gitHash = execSync('git rev-parse HEAD', { 
      cwd: CONTRACT_DIR,
      encoding: 'utf8',
      stdio: ['ignore', 'pipe', 'ignore']
    }).trim();
    return gitHash;
  } catch (error) {
    console.warn('⚠️  Warning: Failed to get git version:', error.message);
    return 'unknown';
  }
}

/**
 * Get current timestamp
 */
function getCurrentTimestamp() {
  return new Date().toISOString();
}

/**
 * Write version file
 */
function writeVersionFile() {
  const versionPath = path.join(BYTECODE_TARGET_DIR, 'version.txt');
  
  const gitVersion = getGitVersion();
  const timestamp = getCurrentTimestamp();
  
  const versionContent = `Git Version: ${gitVersion}
Compile Time: ${timestamp}
`;
  
  fs.writeFileSync(versionPath, versionContent, 'utf8');
  console.log(`✓ Wrote version.txt (git: ${gitVersion.substring(0, 8)}..., time: ${timestamp})`);
}

/**
 * Main function
 */
function main() {
  console.log('🔨 Extracting bytecode from compiled contracts...\n');

  // Check if out directory exists
  if (!fs.existsSync(OUT_DIR)) {
    console.error('❌ Error: out/ directory not found. Please run "forge build" first.');
    process.exit(1);
  }

  let successCount = 0;
  let errorCount = 0;

  // Process each contract
  for (const [contractName, filename] of Object.entries(CONTRACTS)) {
    try {
      console.log(`Extracting ${contractName}...`);
      const bytecode = extractBytecode(contractName);
      writeBytecodeFile(filename, bytecode);
      successCount++;
    } catch (error) {
      console.error(`❌ Error extracting ${contractName}: ${error.message}`);
      errorCount++;
    }
  }

  console.log(`\n✅ Successfully extracted ${successCount} contract(s)`);
  if (errorCount > 0) {
    console.error(`❌ Failed to extract ${errorCount} contract(s)`);
    process.exit(1);
  }

  // Write version file
  console.log('\nWriting version information...');
  writeVersionFile();

  console.log(`\n📁 Bytecode files written to: ${BYTECODE_TARGET_DIR}`);
}

// Run main function
if (require.main === module) {
  main();
}

module.exports = { extractBytecode, writeBytecodeFile, CONTRACTS };

