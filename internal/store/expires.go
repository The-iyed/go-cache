package store

import "time"

func (store *KeyValueStore) Expire(key string, seconds int) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	item, exists := store.data[key]
	if !exists {
		return false
	}

	item.expires = time.Now().Add(time.Duration(seconds) * time.Second)
	store.data[key] = item
	return true
}
