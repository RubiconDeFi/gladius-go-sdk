// This file contains structures and methods for 'GladiusOrder'
package gladiusOrder

import (
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
	"math/big"
)

type GladiusOrder struct {
	Info                   contracts.OrderInfo
	DecayStartTime         uint64
	DecayEndTime           uint64
	ExclusiveFiller        string
	ExclusivityOverrideBps *big.Int
	Input                  *DutchInput
	Outputs                []*DutchOutput
	FillThreshold          *big.Int
}

// TODO: implement all 'OrderMethods'
//func (g *GladiusOrder) Serialize() string           {}
//func (g *GladiusOrder) GetSigner(sig []byte) string {}
//func (g *GladiusOrder) Hash() string                {}
