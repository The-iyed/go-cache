package liststore

import (
	"errors"

	error_message "github.com/go-redis-v1/error"
)

func (ls *ListStore) LRANGE(key string, start, stop int) ([]string, error) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	list, exists := ls.lists[key]
	if !exists {
		return nil, errors.New(error_message.NOT_FOUND)
	}

	if start < 0 {
		start = len(list) + start
	}
	if stop < 0 {
		stop = len(list) + stop
	}

	if start < 0 || stop >= len(list) || start > stop {
		return nil, errors.New(error_message.OUT_OF_RANGE)
	}

	return list[start : stop+1], nil
}
