package liststore

import (
	"errors"

	error_message "github.com/go-redis-v1/error"
)

func (ls *ListStore) LTRIM(key string, start, stop int) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	list, exists := ls.lists[key]
	if !exists {
		return errors.New(error_message.NOT_FOUND)
	}
	if start < 0 {
		start += len(list)
	}
	if stop < 0 {
		stop += len(list)
	}
	if start < 0 || stop >= len(list) || stop > start {
		return errors.New(error_message.OUT_OF_RANGE)
	}

	ls.lists[key] = list[start : stop+1]
	return nil
}
