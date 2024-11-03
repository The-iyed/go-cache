package transaction

import (
	"errors"
)

func (t *Transaction) AbortTransaction() error {

	if !t.IsActive {
		return errors.New("transaction abort failed: no active transaction")
	}

	t.IsActive = false
	t.Commands = nil

	return nil
}
