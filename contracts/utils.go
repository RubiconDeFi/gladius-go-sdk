package contracts

import (
	"math/big"
	"time"
)

// Validates deadline of an order, by checking
// if it has expired or not.
// Returns 'nil' if deadline is valid.
func (o *ResolvedOrder) ValidateDeadline() error {
	return o.ValidateDeadlineSec(0)
}

// Validates deadline of an order, making sure that it
// will be valid for a specified number of seconds - 'sec'.
// To note, remaining time should be only greater than 'sec',
// if it equals to it, error will be returned.
func (o *ResolvedOrder) ValidateDeadlineSec(sec int64) (err error) {
	dd := o.Info.Deadline
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
// should be greater than provided `timestamp`,
// otherwise an error will be returned.
func (o *ResolvedOrder) ValidateDeadlineUnix(timestamp int64) (err error) {
	dd := o.Info.Deadline
	ts := big.NewInt(timestamp)

	if dd.Cmp(ts) != 1 {
		err = InvalidDeadline
	}
	return
}
