package txrules_test

import (
	"testing"

	"github.com/abcsuite/abcd/chaincfg"
	"github.com/abcsuite/abcutil"
	. "github.com/abcsuite/abcwallet/wallet/txrules"
)

func TestStakePoolTicketFee(t *testing.T) {
	params := &chaincfg.MainNetParams
	tests := []struct {
		StakeDiff abcutil.Amount
		Fee       abcutil.Amount
		Height    int32
		PoolFee   float64
		Expected  abcutil.Amount
	}{
		0: {10 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.09231899 * 1e8},
		1: {20 * 1e8, 0.01 * 1e8, 25000, 1.00, 0.17123526 * 1e8},
		2: {5 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.1253002 * 1e8},
		3: {15 * 1e8, 0.05 * 1e8, 50000, 2.59, 0.34447601 * 1e8},
	}
	for i, test := range tests {
		poolFeeAmt := StakePoolTicketFee(test.StakeDiff, test.Fee, test.Height,
			test.PoolFee, params)
		if poolFeeAmt != test.Expected {
			t.Errorf("Test %d: Got %v: Want %v", i, poolFeeAmt, test.Expected)
		}
	}
}
