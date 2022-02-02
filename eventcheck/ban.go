package eventcheck

import (
	base "github.com/TechPay-io/sirius-base/eventcheck"

	"github.com/TechPay-io/go-photon/eventcheck/epochcheck"
)

var (
	ErrAlreadyConnectedEvent = base.ErrAlreadyConnectedEvent
	ErrSpilledEvent          = base.ErrSpilledEvent
	ErrDuplicateEvent        = base.ErrDuplicateEvent
)

func IsBan(err error) bool {
	if err == epochcheck.ErrNotRelevant ||
		err == ErrAlreadyConnectedEvent ||
		err == ErrSpilledEvent ||
		err == ErrDuplicateEvent {
		return false
	}
	return err != nil
}
