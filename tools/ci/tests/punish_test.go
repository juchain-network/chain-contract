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
