// This file contains validation functions for 'GladiusOrder'
// and its JSON analogue.
package gladiusOrder

import (
	"math/big"
	"time"
)

// Validates deadline of an order, by checking
// if it has expired or not.
// Returns 'nil' if deadline is valid.
func (g *GladiusOrderJSON) ValidateDeadline() error {
	return g.ValidateDeadlineSec(0)
}

// Validates deadline of an order, making sure that it
// will be valid for a specified number of seconds - 'sec'.
// To note, remaining time should be only greater than 'sec',
// if it equals to it, error will be returned.
func (g *GladiusOrderJSON) ValidateDeadlineSec(sec int64) (err error) {
	dd := ParseDeadline(g.EncodedOrder)
	now := big.NewInt(time.Now().Unix())
	seconds := big.NewInt(sec)

	diff := big.NewInt(0).Sub(dd, now)

	// Diff should be greater than
	// the number of seconds.
	if diff.Cmp(seconds) != 1 {
		err = InvalidDeadline
	}
	return
}

// Validate deadline of an order against unix timestamp.
// For a deadline to be considered valid here, it
// should be greater than provided 'timest',
// otherwise an error will be returned.
func (g *GladiusOrderJSON) ValidateDeadlineUnix(timest int64) (err error) {
	dd := ParseDeadline(g.EncodedOrder)
	ts := big.NewInt(timest)

	if dd.Cmp(ts) != 1 {
		err = InvalidDeadline
	}
	return
}
