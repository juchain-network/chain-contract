# aac-contracts

## Prepare

Install dependency:

```bash
npm install
```

## unit test

Generate test contract files:

```bash
node generate-mock-contracts.js
```

Start ganache:

```bash
ganache-cli -e 20000000000 -a 100 -l 8000000 -g 0
```

Test:

```bash
truffle test
```

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
env ConstructorArguments="proposal id" npx hardhat --network <network> run scripts/add-node/start_vote.js
```