package contracts

import (
	"math/big"
	"testing"
	"time"
)

var (
	ro = ResolvedOrder{Info: OrderInfo{}}
)

func TestValidateDeadline(t *testing.T) {
	secs := int64(5)
	deadline := time.Now().Unix() + secs
	ro.Info.Deadline = big.NewInt(deadline)

	// Success cases
	if ro.ValidateDeadlineSec(4) != nil {
		t.Error("2) ValidateDeadlineSec(4)")
	}
	if ro.ValidateDeadline() != nil {
		t.Error("2) ValidateDeadline()")
	}
	if ro.ValidateDeadlineUnix(deadline-1) != nil {
		t.Error("3) ValidateDeadlineUnix(deadline-1)")
	}

	// Fail cases
	if ro.ValidateDeadlineSec(secs*2) == nil {
		t.Error("4) ValidateDeadlineSec(secs*2)")
	}
	if ro.ValidateDeadlineUnix(deadline+(secs*2)) == nil {
		t.Error("5) ValidateDeadlineUnix(deadline+(secs*2))")
	}
}
