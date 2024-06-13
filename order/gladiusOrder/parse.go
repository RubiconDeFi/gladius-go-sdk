// This file contains parsing methods to get
// order-related values from 'EncodedOrder'
package gladiusOrder

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Bytes offset in the 'EncodedOrder'
type Offset struct {
	Start, End int
}

var (
	// Offset that contains pointer to the
	// data part of 'OrderInfo'.
	// 64:128 (+ 2 for '0x' in a string)
	ORDER_INFO_PTR = Offset{66, 130}

	/*Offsets in 'OrderInfo' (wihtout '0x'!)*/

	REACTOR  = Offset{64, 128}
	SWAPPER  = Offset{128, 192}
	NONCE    = Offset{192, 256}
	DEADLINE = Offset{256, 320}
	AVC      = Offset{320, 384}
	// Offset, that contains pointer
	// to the data part (size, value).
	AVD_PTR = Offset{384, 448}

	// Decay goes after 'OrderInfo' head part.
	// + 2 for '0x'
	DECAY_START = Offset{130, 194}
	DECAY_END   = Offset{194, 258}

	// The last parameter in the head part
	// of the encoded order (string includes '0x').
	FILL_THRESHOLD = Offset{642, 706}
)

func ParseReactor(encodedOrder string) common.Address {
	offset := infoOffset(encodedOrder)
	reactor := encodedOrder[offset+REACTOR.Start : offset+REACTOR.End]

	return common.HexToAddress(reactor)
}

func ParseSwapper(encodedOrder string) common.Address {
	offset := infoOffset(encodedOrder)
	swapper := encodedOrder[offset+SWAPPER.Start : offset+SWAPPER.End]

	return common.HexToAddress(swapper)
}

func ParseNonce(encodedOrder string) *big.Int {
	offset := infoOffset(encodedOrder)
	nonce := encodedOrder[offset+NONCE.Start : offset+NONCE.End]

	n := new(big.Int)
	n.SetString(nonce, 16)

	return n
}

func ParseDeadline(encodedOrder string) *big.Int {
	offset := infoOffset(encodedOrder)
	deadline := encodedOrder[offset+DEADLINE.Start : offset+DEADLINE.End]

	d := new(big.Int)
	d.SetString(deadline, 16)

	return d
}

func ParseAdditionalValidationContract(encodedOrder string) common.Address {
	offset := infoOffset(encodedOrder)
	avc := encodedOrder[offset+AVC.Start : offset+AVC.End]

	return common.HexToAddress(avc)
}

func ParseAdditionalValidationData(encodedOrder string) []byte {
	/*TODO: implement*/
	return []byte{}
}

func ParseDecayStartTime(encodedOrder string) *big.Int {
	return extractNumber(encodedOrder, DECAY_START)
}

func ParseDecayEndTime(encodedOrder string) *big.Int {
	return extractNumber(encodedOrder, DECAY_END)
}

func ParseFillThreshold(encodedOrder string) *big.Int {
	return extractNumber(encodedOrder, FILL_THRESHOLD)
}

func infoOffset(encodedOrder string) int {
	return getOffset(extractNumber(encodedOrder, ORDER_INFO_PTR))
}

func extractNumber(enco string, offset Offset) *big.Int {
	ptr := enco[offset.Start:offset.End]

	// hex -> dec
	num := new(big.Int)
	num.SetString(ptr, 16)

	return num
}

func getOffset(ptr *big.Int) int {
	return 2 + (int(ptr.Uint64())/32)*64
}
