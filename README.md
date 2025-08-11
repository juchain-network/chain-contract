# aac-contracts

## Prepare

Install dependency:

```bash
npm install
```

## unit test

Generate test contract files:

```bash
node generate-contracts.js --mock
```

Start ganache:

```bash
ganache-cli -e 20000000000 -a 100 -l 8000000 -g 0
```

Test:

```bash
truffle test
```

## Foundry

Foundry tests mirror the Hardhat JS tests and run fully against the same contracts at fixed system addresses.

- Configure: foundry.toml is set to solc 0.8.20, src=contracts, test=forge-tests.
- Run tests:

```bash
forge test
```

### Test Coverage

Foundry tests provide complete parity with Hardhat test suites:

- **ProposalFoundry.t.sol**: Proposal creation, voting, constraints, and config updates
- **ValidatorsFoundry.t.sol**: Basic validator reward distribution and profit withdrawal
- **ValidatorsCompleteFoundry.t.sol**: Complete validator lifecycle (create/edit, proposals, rewards, set updates)
- **PunishFoundry.t.sol**: Punishment thresholds, jailing, and missed block handling
- **RewardFoundry.t.sol**: Advanced reward distribution scenarios (punish redistribution, jailed validator exclusion)
- **Proposal.t.sol**: Basic smoke test for receiver address initialization

Total: **30 tests** covering all contract functionality.

### Foundry scripts (parity with Hardhat scripts)

Comprehensive Solidity scripts that replicate and extend the Hardhat scripts under `forge-scripts/`:

#### Basic Scripts

- `AddNewNode.s.sol`: create an add-validator proposal
- `RemoveNode.s.sol`: create a remove-validator proposal
- `UpdateConfig.s.sol`: create a config update proposal

#### Enhanced Scripts

- `CreateProposal.s.sol`: enhanced proposal creation with validation and ID return
- `VoteProposal.s.sol`: vote on proposals by ID (with convenience functions for yes/no)
- `EndToEndProposal.s.sol`: complete proposal + voting workflow with multiple validators
- `DeploySystem.s.sol`: system initialization and status checking utilities

#### Usage Examples

Basic proposal creation:

```bash
forge script forge-scripts/AddNewNode.s.sol \
  --rpc-url $RPC_URL \
  --broadcast \
  --private-key $VALIDATOR_KEY \
  --sig "run(address)" 0xYourNewValidator
```

Vote on a proposal:

```bash
forge script forge-scripts/VoteProposal.s.sol \
  --rpc-url $RPC_URL \
  --broadcast \
  --private-key $VALIDATOR_KEY \
  --sig "voteYes(bytes32)" 0xProposalId
```

End-to-end workflow (create + vote with multiple validators):

```bash
forge script forge-scripts/EndToEndProposal.s.sol \
  --rpc-url $RPC_URL \
  --broadcast \
  --private-key $PROPOSER_KEY \
  --sig "runAddValidatorFlow(address,string,address[])" \
  0xNewValidator "Add new validator" "[0xValidator1,0xValidator2,0xValidator3]"
```

Config update:

```bash
forge script forge-scripts/UpdateConfig.s.sol \
  --rpc-url $RPC_URL \
  --broadcast \
  --private-key $VALIDATOR_KEY \
  --sig "run(uint256,uint256)" 4 100  # Update withdrawProfitPeriod to 100 blocks
```

Check system status:

```bash
forge script forge-scripts/DeploySystem.s.sol \
  --rpc-url $RPC_URL \
  --sig "checkSystemStatus()" \
  --call
```

Note: The basic scripts only create proposals. Use `VoteProposal.s.sol` or `EndToEndProposal.s.sol` for voting workflows.

## How to join new absenteeism

Modify .Env file Add private key to

```bash
MINER="your node private key"
```

Create a Proposal

```bash
env ConstructorArguments="new address" npx hardhat --network <network> run scripts/add-node/create_proposal.js
```

Start a Vote

```bash
env ConstructorArguments="proposal id" npx hardhat --network network_name run scripts/add-node/start_vote.js
```
env ConstructorArguments="proposal id" npx hardhat --network <network> run scripts/add-node/start_vote.js
```