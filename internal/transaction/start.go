package transaction

import (
	"errors"
)

func (t *Transaction) StartTransaction() error {

	if t.IsActive {
		return errors.New("transaction start failed: active transactio already running")
	}

	t.IsActive = true

	return nil
}
