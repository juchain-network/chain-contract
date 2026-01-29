# Integration Testing Findings

## Potential Contract Bugs
1. **Proposal.sol: initialize()**: The `initialize` function does not seem to prevent multiple calls, or the storage was not correctly marked as initialized in the genesis template. Tests show it doesn't revert on second call.
2. **Proposal.sol: createUpdateConfigProposal()**: Does not validate `cid` or `value` during proposal creation. Invalid IDs or zero values are accepted, potentially failing only at the execution stage (which is too late/risky).
3. **Staking.sol: self-delegation**: The contract seems to allow a validator to delegate to themselves beyond the `RegisterValidator` call in some circumstances? Need to verify.

## System/Environment Constraints
1. **POA System Transactions**: The node (juchain) blocks direct calls to `onlyMiner` functions from user accounts with `forbidden system transaction`. These cannot be tested via RPC.
2. **Proposer Cooldown**: `Proposal.sol` enforces a 1-block cooldown per proposer. Tests must rotate genesis validators or wait.
3. **Consensus Health**: Sequential removal of genesis validators without starting new nodes will stall the network.
4. **Gas Pricing**: POA nodes require careful Legacy vs EIP-1559 handling. Forcing Legacy with 1 Gwei is currently the most stable approach.

## Test Suite Status
- **Staking Management**: PASSED
- **Delegation Flow**: PASSED
- **Governance Proposals**: PASSED
- **Punishment/DoubleSign**: PARTIALLY PASSED (Environment sensitive)
- **Config Updates**: PARTIALLY PASSED (Positive cases pass, boundary checks fail due to contract logic)
