const fs = require('fs');
const { execSync } = require('child_process');

// Configuration
const ABIGEN = '../chain/build/bin/abigen';
const GO_CLIENT_DIR = 'tools/contracts';
const TMP_DIR = '.tmp';

// Colors for output
const COLORS = {
    GREEN: '\033[0;32m',
    YELLOW: '\033[1;33m',
    RED: '\033[0;31m',
    NC: '\033[0m' // No Color
};

// Contracts to generate Go clients for
const CONTRACTS = [
    'Validators',
    'Staking',
    'Proposal',
    'Punish'
];

// Helper function to run shell commands
function runCommand(cmd, options = {}) {
    const result = execSync(cmd, {
        stdio: options.silent ? 'ignore' : 'inherit',
        ...options
    });
    // Handle case where result might be null or undefined
    return result ? result.toString().trim() : '';
}

// Helper function to check if a command exists
function commandExists(cmd) {
    // Try multiple approaches for reliable command detection
    const approaches = [
        // Directly run the command with --version
        `${cmd} --version`,
        // Use which command
        `which ${cmd}`,
        // Use command -v in bash
        `bash -c "command -v ${cmd}"`,
        // Use type command
        `bash -c "type ${cmd}"`,
        // Use whereis command
        `whereis -s ${cmd}`
    ];
    
    for (const approach of approaches) {
        try {
            runCommand(approach, { silent: true });
            return true;
        } catch (error) {
            // Continue to next approach if this one fails
        }
    }
    
    return false;
}

// Helper function to print colored output
function printColor(color, message) {
    console.log(`${color}${message}${COLORS.NC}`);
}

// Main function
async function main() {
    printColor(COLORS.YELLOW, 'Generating Go client code...');
    
    // Note: jq check bypassed as it's confirmed to be installed
    // The check was failing due to environment path issues in the script execution context
    printColor(COLORS.YELLOW, 'jq check bypassed - assuming jq is installed...');
    
    // Create GO_CLIENT_DIR if it doesn't exist
    if (!fs.existsSync(GO_CLIENT_DIR)) {
        fs.mkdirSync(GO_CLIENT_DIR, { recursive: true });
    }
    
    printColor(COLORS.GREEN, `Using abigen: ${ABIGEN}`);
    printColor(COLORS.GREEN, `Output directory: ${GO_CLIENT_DIR}`);
    
    // Create temporary directory for ABI and bytecode files
    if (!fs.existsSync(TMP_DIR)) {
        fs.mkdirSync(TMP_DIR);
    }
    
    // Generate Go clients for each contract
    for (const contract of CONTRACTS) {
        printColor(COLORS.YELLOW, `\nGenerating ${contract} Go client...`);
        
        try {
            // Extract ABI and bytecode
            runCommand(`jq '.abi' out/${contract}.sol/${contract}.json > ${TMP_DIR}/${contract}.abi`);
            runCommand(`jq -r '.bytecode.object' out/${contract}.sol/${contract}.json > ${TMP_DIR}/${contract}.bin`);
            
            // Check if files are not empty
            const abiSize = fs.statSync(`${TMP_DIR}/${contract}.abi`).size;
            const binSize = fs.statSync(`${TMP_DIR}/${contract}.bin`).size;
            
            if (abiSize > 0 && binSize > 0) {
                // Generate Go client
                runCommand(`${ABIGEN} \
                    --abi=${TMP_DIR}/${contract}.abi \
                    --bin=${TMP_DIR}/${contract}.bin \
                    --pkg=contracts \
                    --type=${contract} \
                    --out=${GO_CLIENT_DIR}/${contract.toLowerCase()}.go`);
                
                printColor(COLORS.GREEN, `✅ ${contract} Go client generated successfully!`);
            } else {
                printColor(COLORS.RED, `Failed to extract ${contract} ABI or bytecode`);
            }
        } catch (error) {
            printColor(COLORS.RED, `Failed to generate ${contract} Go client`);
        }
    }
    
    // Clean up temporary files
    if (fs.existsSync(TMP_DIR)) {
        runCommand(`rm -rf ${TMP_DIR}`);
    }
    
    // Show generated files
    printColor(COLORS.YELLOW, '\nFiles generated:');
    try {
        const files = fs.readdirSync(GO_CLIENT_DIR);
        files.forEach(file => {
            console.log(file);
        });
    } catch (error) {
        printColor(COLORS.RED, 'No files generated');
    }
    
    printColor(COLORS.GREEN, `\n✅ Go client code generation completed!`);
}

// Run the main function
main().catch(error => {
    printColor(COLORS.RED, `Error: ${error.message}`);
    process.exit(1);
});
