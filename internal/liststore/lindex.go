package liststore

import (
	"errors"

	error_message "github.com/go-redis-v1/error"
)

func (ls *ListStore) LINDEX(key string, index int) (string, error) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	list, exists := ls.lists[key]
	if !exists {
		return "", errors.New(error_message.NOT_FOUND)
	}

	if index < 0 {
		index = len(list) + index
	}

	if index < 0 || index >= len(list) {
		return "", errors.New(error_message.OUT_OF_RANGE)
	}

	return list[index], nil
}
