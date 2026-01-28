package tests

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestG_DoubleSign(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	// [P-07] Submit Double Sign Evidence
	t.Run("P-07_DoubleSignEvidence", func(t *testing.T) {
		// Use Genesis Validator 0 as the double signer
		valKey := ctx.GenesisValidators[0]
		valAddr := crypto.PubkeyToAddress(valKey.PublicKey)
		
		// Use a clean account to submit evidence
		reporterKey, reporterAddr, err := ctx.CreateAndFundAccount(utils.ToWei(10))
		utils.AssertNoError(t, err, "failed to setup reporter")

		// 1. Prepare Block Height
		// Must be in the past but within doubleSignWindow
		// And usually header timestamp must be valid?
		// Punish.sol checks: block.number >= number1
		// So we use current block number - 1
		header, _ := ctx.Clients[0].HeaderByNumber(nil, nil)
		targetHeight := new(big.Int).Sub(header.Number, big.NewInt(1))
		if targetHeight.Cmp(big.NewInt(0)) <= 0 {
			targetHeight = big.NewInt(1)
		}

		t.Logf("Constructing double sign evidence for validator %s at height %s", valAddr.Hex(), targetHeight)

		// 2. Construct Two Headers
		// Header 1
		h1 := &types.Header{
			ParentHash:  common.Hash{}, // Dummy
			UncleHash:   types.EmptyUncleHash,
			Coinbase:    valAddr,
			Root:        common.Hash{},
			TxHash:      types.EmptyRootHash,
			ReceiptHash: types.EmptyRootHash,
			Bloom:       types.Bloom{},
			Difficulty:  big.NewInt(1),
			Number:      targetHeight,
			GasLimit:    30000000,
			GasUsed:     0,
			Time:        uint64(100000), // Dummy time
			Extra:       make([]byte, 32+65), // Vanity + Signature
			MixDigest:   common.Hash{},
			Nonce:       types.BlockNonce{},
		}
		// Header 2 (Different StateRoot to ensure different Hash)
		h2 := &types.Header{
			ParentHash:  common.Hash{},
			UncleHash:   types.EmptyUncleHash,
			Coinbase:    valAddr,
			Root:        common.Hash{0x01}, // Different
			TxHash:      types.EmptyRootHash,
			ReceiptHash: types.EmptyRootHash,
			Bloom:       types.Bloom{},
			Difficulty:  big.NewInt(1),
			Number:      targetHeight,
			GasLimit:    30000000,
			GasUsed:     0,
			Time:        uint64(100000),
			Extra:       make([]byte, 32+65),
			MixDigest:   common.Hash{},
			Nonce:       types.BlockNonce{},
		}

		// 3. Sign Headers (Clique Style)
		rlp1, err := signHeaderClique(h1, valKey)
		utils.AssertNoError(t, err, "failed to sign h1")
		
		rlp2, err := signHeaderClique(h2, valKey)
		utils.AssertNoError(t, err, "failed to sign h2")

		// 4. Record State Before
		infoBefore, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		reporterBalBefore, _ := ctx.Clients[0].BalanceAt(nil, reporterAddr, nil)
		
		t.Logf("Validator Stake Before: %s", infoBefore.SelfStake)
		t.Logf("Reporter Balance Before: %s", reporterBalBefore)

		// 5. Submit Evidence
		opts, _ := ctx.GetTransactor(reporterKey)
		tx, err := ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1, rlp2)
		utils.AssertNoError(t, err, "failed to submit evidence")
		ctx.WaitMined(tx.Hash())

		// 6. Verify Slash
		infoAfter, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		t.Logf("Validator Stake After: %s", infoAfter.SelfStake)
		
		// Stake should decrease
		utils.AssertTrue(t, infoAfter.SelfStake.Cmp(infoBefore.SelfStake) < 0, "Validator should be slashed")
		
		// Validator should be jailed
		utils.AssertTrue(t, infoAfter.IsJailed, "Validator should be jailed")
		
		// Reporter should get reward
		reporterBalAfter, _ := ctx.Clients[0].BalanceAt(nil, reporterAddr, nil)
		t.Logf("Reporter Balance After: %s", reporterBalAfter)
		
		// Check if reporter balance increased (ignoring gas cost approx)
		// Since gas cost is small compared to reward (typically), this check usually holds if reward is large enough.
		// DoubleSignRewardAmount defaults to 10000 ETH? Or similar large number.
		// Let's check exact params.
		reward, _ := ctx.Proposal.DoubleSignRewardAmount(nil)
		t.Logf("Expected Reward: %s", reward)
		
		// Use Approximate check
		utils.AssertTrue(t, reporterBalAfter.Cmp(reporterBalBefore) > 0, "Reporter should receive reward")

		// Duplicate evidence should fail
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1, rlp2)
		if err == nil {
			t.Fatal("Duplicate evidence should fail")
		}
	})

	// [P-21] Resign + Double Sign (Should work)
	t.Run("P-21_ResignThenDoubleSign", func(t *testing.T) {
		key, addr, _ := createAndRegisterValidator(t, "P-21 ResignDS")
		opts, _ := ctx.GetTransactor(key)
		
		// 1. Resign
		txR, _ := ctx.Staking.ResignValidator(opts)
		ctx.WaitMined(txR.Hash())
		
		// 2. Submit Double Sign Evidence
		header, _ := ctx.Clients[0].HeaderByNumber(nil, nil)
		targetHeight := new(big.Int).Sub(header.Number, big.NewInt(1))
		
		h1 := &types.Header{Coinbase: addr, Number: targetHeight, Extra: make([]byte, 32+65), Root: common.Hash{0x21}}
		h2 := &types.Header{Coinbase: addr, Number: targetHeight, Extra: make([]byte, 32+65), Root: common.Hash{0x22}}
		
		rlp1, _ := signHeaderClique(h1, key)
		rlp2, _ := signHeaderClique(h2, key)
		
		txDS, err := ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1, rlp2)
		utils.AssertNoError(t, err, "Should allow double sign punishment after resign")
		ctx.WaitMined(txDS.Hash())
	})

	// [P-22] Exit + Double Sign (Should fail)
	t.Run("P-22_ExitThenDoubleSign", func(t *testing.T) {
		key, addr, _ := createAndRegisterValidator(t, "P-22 ExitDS")
		opts, _ := ctx.GetTransactor(key)
		
		// 1. Resign -> Wait -> Exit
		ctx.Staking.ResignValidator(opts)
		waitBlocks(t, 55)
		txE, err := ctx.Staking.ExitValidator(opts)
		utils.AssertNoError(t, err, "Exit failed")
		ctx.WaitMined(txE.Hash())
		
		// 2. Submit Double Sign Evidence
		header, _ := ctx.Clients[0].HeaderByNumber(nil, nil)
		h1 := &types.Header{Coinbase: addr, Number: header.Number, Extra: make([]byte, 32+65), Root: common.Hash{0x31}}
		h2 := &types.Header{Coinbase: addr, Number: header.Number, Extra: make([]byte, 32+65), Root: common.Hash{0x32}}
		rlp1, _ := signHeaderClique(h1, key)
		rlp2, _ := signHeaderClique(h2, key)
		
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1, rlp2)
		if err == nil {
			t.Fatal("Should fail double sign punishment after exit (validator not exist)")
		}
	})

	// [P-10~P-14] Double Sign Exceptions
	t.Run("P-10-14_DoubleSignExceptions", func(t *testing.T) {
		key, addr, err := createAndRegisterValidator(t, "DS Exceptions")
		if err != nil {
			t.Skipf("create validator failed: %v", err)
		}
		opts, _ := ctx.GetTransactor(key)
		header, _ := ctx.Clients[0].HeaderByNumber(nil, nil)
		
		hBase := &types.Header{Coinbase: addr, Number: header.Number, Extra: make([]byte, 32+65)}
		
		// P-11: Same Header
		h1_same, _ := signHeaderClique(hBase, key)
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, h1_same, h1_same)
		if err == nil { t.Fatal("Should fail with 'Same header'") }
		
		// P-12: Height Mismatch
		h1_h1 := *hBase
		h2_h2 := *hBase
		h2_h2.Number = new(big.Int).Add(hBase.Number, big.NewInt(1))
		h2_h2.Root = common.Hash{0x01}
		rlp1, _ := signHeaderClique(&h1_h1, key)
		rlp2, _ := signHeaderClique(&h2_h2, key)
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1, rlp2)
		if err == nil { t.Fatal("Should fail with 'Height mismatch'") }
		
		// P-14: Signer != Coinbase (Using different key to sign)
		otherKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		h1_wrong := *hBase
		h1_wrong.Root = common.Hash{0x05}
		h2_wrong := *hBase
		h2_wrong.Root = common.Hash{0x06}
		rlp1_wrong, _ := signHeaderClique(&h1_wrong, otherKey)
		rlp2_wrong, _ := signHeaderClique(&h2_wrong, otherKey)
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlp1_wrong, rlp2_wrong)
		if err == nil { t.Fatal("Should fail with 'Signer != coinbase'") }

		// P-10: Future block
		hFuture := *hBase
		hFuture.Number = new(big.Int).Add(hBase.Number, big.NewInt(1))
		hFuture.Root = common.Hash{0x07}
		rlpFuture1, _ := signHeaderClique(&hFuture, key)
		hFuture2 := hFuture
		hFuture2.Root = common.Hash{0x08}
		rlpFuture2, _ := signHeaderClique(&hFuture2, key)
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlpFuture1, rlpFuture2)
		if err == nil { t.Fatal("Should fail with 'Future block'") }

		// P-13: Non-validator signer
		nonValKey, nonValAddr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		hNonVal := &types.Header{Coinbase: nonValAddr, Number: hBase.Number, Extra: make([]byte, 32+65), Root: common.Hash{0x09}}
		hNonVal2 := &types.Header{Coinbase: nonValAddr, Number: hBase.Number, Extra: make([]byte, 32+65), Root: common.Hash{0x0a}}
		rlpNon1, _ := signHeaderClique(hNonVal, nonValKey)
		rlpNon2, _ := signHeaderClique(hNonVal2, nonValKey)
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlpNon1, rlpNon2)
		if err == nil { t.Fatal("Should fail with 'Signer not exist'") }

		// P-10 (Malformed): invalid header bytes
		_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, []byte{0x01, 0x02}, []byte{0x03})
		if err == nil { t.Fatal("Should fail with malformed header") }

		// P-11: Evidence expired (if chain height is high enough)
		window, _ := ctx.Proposal.DoubleSignWindow(nil)
		if hBase.Number.Cmp(window) > 0 {
			expiredHeight := new(big.Int).Sub(hBase.Number, new(big.Int).Add(window, big.NewInt(1)))
			hExp1 := &types.Header{Coinbase: addr, Number: expiredHeight, Extra: make([]byte, 32+65), Root: common.Hash{0x0b}}
			hExp2 := &types.Header{Coinbase: addr, Number: expiredHeight, Extra: make([]byte, 32+65), Root: common.Hash{0x0c}}
			rlpExp1, _ := signHeaderClique(hExp1, key)
			rlpExp2, _ := signHeaderClique(hExp2, key)
			_, err = ctx.Punish.SubmitDoubleSignEvidence(opts, rlpExp1, rlpExp2)
			if err == nil { t.Fatal("Should fail with expired evidence") }
		}
	})
}

// signHeaderClique calculates the signature of the header and returns RLP encoded header with signature
func signHeaderClique(h *types.Header, key *ecdsa.PrivateKey) ([]byte, error) {
	// 1. Encode header without signature (extra data suffix)
	// Clique: The signature is the last 65 bytes of Extra
	// We need to hash the header with Extra trimmed by 65 bytes
	
	origExtra := h.Extra
	if len(origExtra) < 65 {
		// Should be at least 65 bytes
		h.Extra = make([]byte, 65)
	}
	
	// Create a copy for hashing with truncated signature
	headerForHash := *h
	// Punish.sol implementation strips the last 65 bytes completely
	
	extraCopy := make([]byte, len(h.Extra)-65)
	copy(extraCopy, h.Extra[:len(h.Extra)-65])
	headerForHash.Extra = extraCopy
	
	// Hash
	// We need to RLP encode this headerForHash and then Keccak256
	encodedForHash, err := rlp.EncodeToBytes(&headerForHash)
	if err != nil { return nil, err }
	hash := crypto.Keccak256(encodedForHash)
	
	// 2. Sign
	sig, err := crypto.Sign(hash, key)
	if err != nil { return nil, err }
	
	// 3. Put signature back into original header
	// Note: headerForHash was a shallow copy struct, but Extra was a new slice.
	// We need to modify the ORIGINAL h's Extra.
	// Ensure h.Extra has enough space
	if len(h.Extra) < 65 {
		h.Extra = make([]byte, 32+65)
	}
	// Copy sig into the last 65 bytes
	copy(h.Extra[len(h.Extra)-65:], sig)
	
	// 4. Return RLP of the fully signed header
	return rlp.EncodeToBytes(h)
}
