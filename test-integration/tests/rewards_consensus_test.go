package tests

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
)

func TestI_ConsensusRewards(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	t.Run("V-03_DistributeBlockReward", func(t *testing.T) {
		t.Skip("Blocked by node (forbidden system transaction)")
	})

	t.Run("S-22_DistributeRewardsAndCooldown", func(t *testing.T) {
		t.Skip("Blocked by node (forbidden system transaction)")
	})
	
	t.Run("QueryRewards", func(t *testing.T) {
		valAddr := ctx.Config.Validators[0].Address
		_, _, _, _, rew, _ := ctx.Validators.GetValidatorInfo(nil, common.HexToAddress(valAddr))
		t.Logf("Validator %s rewards: %s", valAddr, rew.String())
	})
}
