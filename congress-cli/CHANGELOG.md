# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
