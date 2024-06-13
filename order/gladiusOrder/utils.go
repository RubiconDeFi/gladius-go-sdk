package gladiusOrder

import (
	"math/big"
	"time"
)

// Applies decay on 'DutchInput'. Returns Point-In-Time value
// for a decay range, using 'time.Now().Unix()' timestamp
// during the function call.
func (i DutchInput) Decay(dStart, dEnd *big.Int) *big.Int {
	return Decay(dStart, dEnd, i.StartAmount, i.EndAmount)
}

// Applies decay on 'DutchOutput'. Returns Point-In-Time value
// for a decay range, using 'time.Now().Unix()' timestamp
// during the function call.
func (o DutchOutput) Decay(dStart, dEnd *big.Int) *big.Int {
	return Decay(dStart, dEnd, o.StartAmount, o.EndAmount)
}

// Implements 'decay(uint256,uint256,uint256,uint256)' function from:
// https://github.com/RubiconDeFi/gladius-protocol/blob/master/src/lib/DutchDecayLib.sol#L26
func Decay(dstart, dend *big.Int, samt, eamt string) (decayedAmt *big.Int) {
	start := new(big.Int)
	end := new(big.Int)
	start.SetString(samt, 10)
	end.SetString(eamt, 10)

	now := big.NewInt(time.Now().Unix())
	decayedAmt = big.NewInt(0)

	nowCmpStart := now.Cmp(dstart)
	nowCmpEnd := now.Cmp(dend)

	if nowCmpEnd == 1 || nowCmpEnd == 0 {
		decayedAmt = end
	} else if nowCmpStart == -1 || nowCmpStart == 0 {
		decayedAmt = start
	} else {
		elapsed := big.NewInt(0).Sub(now, dstart)
		duration := big.NewInt(0).Sub(dend, dstart)

		if start.Cmp(end) == 1 {
			ssube := big.NewInt(0).Sub(start, end)
			muldiv := big.NewInt(0).Div(
				big.NewInt(0).Mul(ssube, elapsed),
				duration)

			decayedAmt = decayedAmt.Sub(start, muldiv)
		} else {
			esubs := big.NewInt(0).Sub(end, start)
			muldiv := big.NewInt(0).Div(
				big.NewInt(0).Mul(esubs, elapsed), duration)

			decayedAmt = big.NewInt(0).Add(start, muldiv)
		}
	}

	return
}
