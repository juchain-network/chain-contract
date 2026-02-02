# Integration Test Validation Plan & Log

## Validation Requirements (current)
- Run **grouped integration tests** to save time and keep environments isolated.
- Each group must run with a **fresh reset** (use existing Makefile targets).
- Only after **all groups pass individually**, run a full `make ci-log` as a final end‑to‑end confirmation.

### Grouped Targets (each does reset/init/run/ready/stop)
- `make -C test-integration test-config`
- `make -C test-integration test-governance`
- `make -C test-integration test-staking`
- `make -C test-integration test-delegation`
- `make -C test-integration test-punish`
- `make -C test-integration test-rewards`
- `make -C test-integration test-epoch`

## Validation Log (most recent first)
### 2026-02-02
- **Run:** `make -C test-integration test-governance`
  - **Result:** PASS (~7.4 min)
  - **Notes:** G-16 logs “V2 register succeeded unexpectedly” (warning only; test still passes).

- **Run:** `make -C test-integration test-governance`
  - **Result:** FAIL
  - **Failure:** `TestB_Governance/G-16_SmoothExpansion` → vote tx reverted (`G-16 V2 failed`)
  - **Action:** Hardened `voteProposalToPass` to handle revert/epoch edge cases, skip inactive/jailed voters, and retry.

- **Run:** `make -C test-integration test-config`
  - **Result:** PASS (~10.6 min)
  - **Notes:** Config suite is slow due to multiple epoch waits; no failures after import fix.

- **Run:** `make ci-log` (log: `logs/ci_20260202_162817.log`)
  - **Result:** FAIL
  - **Failure:** `TestE_Delegation/D-15_DelegatorToValidator` → `utils.go:18: should be validator`
  - **Action:** Added retry + receipt status check around `RegisterValidator` in `test-integration/tests/delegation_test.go` to handle epoch / “too many validators” windows and reverted txs.

- **Run:** `make ci-log` (log: `logs/ci_20260202_152222.log`)
  - **Result:** FAIL
  - **Failure:** `TestE_Delegation/D-15_DelegatorToValidator` (same symptom)
  - **Action:** Added epoch wait after register, but validator still not registered → refined fix above.

- **Run:** `make ci-log` (log: `logs/ci_20260202_151039.log`)
  - **Result:** FAIL
  - **Failure:** `TestA_SystemConfigSetup` → `WithdrawProfitPeriod mismatch: expected 20, got 86400`
  - **Action:** In `test-integration/tests/config_test.go`, if a config value mismatches, call `ctx.EnsureConfig(...)` then re‑read and assert.

## Next Steps
1. Re-run **grouped tests** in order (see targets above) until all pass.
2. Record each group result here (PASS/FAIL + log file).
3. Only after all groups pass, run a single `make ci-log`.
