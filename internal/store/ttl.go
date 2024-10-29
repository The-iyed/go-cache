package store

import "time"

func (store *KeyValueStore) TTL(key string) int64 {
	store.mu.RLock()
	defer store.mu.RUnlock()

	item, exists := store.data[key]
	if !exists {
		return -2
	}

	if item.expires.IsZero() {
		return -1
	}

	remaining := item.expires.Sub(time.Now()).Seconds()
	if remaining <= 0 {
		return 0
	}

	return int64(remaining)
}
