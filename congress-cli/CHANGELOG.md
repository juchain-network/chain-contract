# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2025-08-25

### Added

- **JPoSA Staking Support**: Complete integration of Juchain Proof of Stake Authority consensus
  - New `staking` command group with comprehensive validator and delegation management
  - Support for validator registration with self-staking and commission rates
  - Token delegation and undelegation with 7-day unbonding period
  - Staking rewards claiming functionality
  - Real-time validator and delegation status queries

- **Staking Transaction Commands**
  - `register-validator`: Register as a validator with minimum 10,000 JU stake
  - `delegate`: Delegate tokens to existing validators for staking rewards
  - `undelegate`: Start unbonding process to withdraw delegated tokens
  - `claim-rewards`: Claim accumulated staking rewards from validation or delegation

- **Staking Query Commands**
  - `query-validator`: Get detailed validator information (stake, commission, jail status)
  - `query-delegation`: View delegation details and pending rewards
  - `list-top-validators`: List top validators ranked by total stake

- **Enhanced Smart Contract Integration**
  - Staking contract ABI integration for seamless interaction
  - Support for Staking contract at address `0x000000000000000000000000000000000000f003`
  - Automatic wei/JU conversion for user-friendly amount handling
  - Comprehensive error handling for contract calls

- **Improved Transaction Workflow**
  - Consistent transaction file generation for all staking operations
  - Generated files: `registerValidator.json`, `delegate.json`, `undelegate.json`, `claimRewards.json`
  - Full compatibility with existing sign and send workflow
  - Parameter validation for all staking operations

- **Documentation and Testing**
  - Comprehensive `STAKING_USAGE.md` guide with detailed examples
  - Integration test script `test_staking.sh` for command validation
  - Updated architecture documentation in `contracts/README.md`
  - Complete CLI command examples and best practices

### Changed

- **Root Command Integration**
  - Added staking commands to main CLI interface
  - Updated help system to include staking operations
  - Enhanced RPC validation to include staking commands

- **Configuration Updates**
  - Added `StakingContractAddr` constant for contract address management
  - Extended transaction file naming conventions for staking operations
  - Updated global flag validation to support new staking commands

### Technical Improvements

- **Code Architecture**
  - Modular staking command implementation in `cmd/staking.go`
  - Consistent error handling and user feedback patterns
  - Reusable validation functions for addresses and amounts

- **Smart Contract Interaction**
  - Direct contract calls for query operations without transaction fees
  - Proper ABI encoding/decoding for all staking contract methods
  - Support for complex return types (structs, arrays)

- **User Experience**
  - Clear command descriptions and parameter explanations
  - Helpful error messages for common validation failures
  - Consistent output formatting across all staking commands

## [1.1.0] - 2025-08-12

### Added

- **Enhanced Input Validation**: Comprehensive parameter validation for all commands
  - Address format validation using `ValidateAddress()`
  - Chain ID validation with `ValidateChainID()`
  - RPC URL format validation with `ValidateRPCURL()`
  - Proposal ID format validation with `ValidateProposalID()`
  - Operation type validation with `ValidateOperation()`
  - Configuration ID validation with `ValidateConfigID()`

- **Improved User Experience**
  - Colorful output with emoji indicators (✅, ❌, ℹ️, ⚠️)
  - Structured error messages with `PrintError()`, `PrintValidationError()`
  - Success confirmation messages with `PrintSuccess()`
  - Informational messages with `PrintInfo()` and `PrintWarning()`

- **Global Parameter Validation**
  - Automatic validation of RPC URL and Chain ID for relevant commands
  - Pre-execution validation in `validateGlobalFlags()`

- **Configuration Management**
  - Centralized constants in `config.go`
  - Configurable contract addresses
  - Configurable timeout and gas multiplier settings
  - Config ID name mapping for better UX

- **Enhanced Help System**
  - Comprehensive help documentation with detailed descriptions
  - New `examples` command showing practical usage examples
  - Improved version command with build information

- **Better Error Handling**
  - All internal functions now return proper errors
  - Wrapped errors with context using `fmt.Errorf(..., %w, err)`
  - Consistent error propagation throughout the codebase

- **Vote Command Improvements**
  - Simplified voting syntax (use `-a` for approve, omit for reject)
  - Removed required flag for approve parameter
  - Clear indication of vote type in output

- **Enhanced Makefile**
  - New targets: `help`, `test`, `lint`, `fmt`, `tidy`, `install`, `release`
  - Build versioning with Git commit and build date
  - Colored output and progress indicators
  - Cross-platform build support

### Changed

- **Command Interface Updates**
  - All transaction creation functions now return errors
  - Improved function signatures for better error handling
  - Updated command descriptions and help text

- **Code Organization**
  - Extracted validation functions to `utils.go`
  - Centralized configuration in `config.go`
  - Better separation of concerns

- **Transaction File Naming**
  - Using constants for transaction file names
  - Consistent naming across all commands

### Fixed

- **Vote Proposal Command**
  - Fixed boolean flag handling for approve/reject votes
  - Removed incorrect required flag for approve parameter

- **Error Handling**
  - Fixed missing error returns in internal functions
  - Proper error propagation to user interface

- **Address Validation**
  - Improved Ethereum address format validation
  - Better handling of edge cases

### Technical Improvements

- **Test Coverage**
  - Added comprehensive validation function tests
  - Test cases for edge cases and error conditions

- **Code Quality**
  - Consistent error handling patterns
  - Reduced code duplication
  - Better function naming and documentation

- **Build System**
  - Enhanced Makefile with multiple build targets
  - Version injection during build process
  - Support for static linking on Linux

## [1.0.0] - Previous

### Initial Features

- Basic proposal creation and voting
- Validator management commands  
- Transaction signing and broadcasting
- RPC interaction capabilities

---

### Migration Guide from v1.1.0 to v1.2.0

#### New Staking Commands

The v1.2.0 release introduces comprehensive staking functionality for JPoSA consensus:

```bash
# Check available staking commands
./congress-cli staking --help

# Register as a validator (example)
./congress-cli staking register-validator \
  --rpc_laddr http://localhost:8545 \
  --proposer 0x1234567890123456789012345678901234567890 \
  --stake-amount 10000 \
  --commission-rate 500

# Query validator information
./congress-cli staking query-validator \
  --rpc_laddr http://localhost:8545 \
  --address 0x1234567890123456789012345678901234567890
```

#### Breaking Changes

None. All existing commands remain fully compatible.

#### New Requirements

- **Minimum Stake**: Validator registration requires minimum 10,000 JU tokens
- **Commission Rate**: Must be between 0-10,000 basis points (0-100%)
- **Unbonding Period**: Undelegation has a 7-day waiting period

#### Recommended Workflow

1. **Query Network**: Use `list-top-validators` to see active validators
2. **Create Transactions**: Use staking commands to generate transaction files
3. **Sign & Send**: Use existing `sign` and `send` commands for transaction broadcast
4. **Monitor**: Use query commands to track staking status and rewards

---

### Migration Guide from v1.0.0 to v1.1.0

#### Command Changes

- **Vote Command**: No breaking changes, but you can now omit the `-a` flag for reject votes
- **Error Output**: Error messages are now more descriptive and include validation details

#### New Requirements

- Commands now validate parameters before execution
- Invalid parameters will be rejected with clear error messages

#### New Features to Try

```bash
# View examples
./congress-cli examples

# Enhanced version info
./congress-cli version

# Improved help
./congress-cli --help
./congress-cli [command] --help
```
