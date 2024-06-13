// This file includes helper functions to call methods of certain contracts.
package binds

import (
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
)

// Wrapper for 'GetFeeOutputs' function of 'RubiconFeeController'.
func (bopt *BindOpts) GetFee(r *contracts.ResolvedOrder) ([]contracts.OutputToken, error) {
	ctrl, err := bopt.FeeController()
	if err != nil {
		return nil, err
	}

	out, err := ctrl.GetFeeOutputs(nil, *r)
	if err != nil {
		return nil, err
	}

	return out, nil
}
