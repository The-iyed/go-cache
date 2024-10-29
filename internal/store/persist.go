package store

import "time"

func (store *KeyValueStore) Persist(key string) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	item, exists := store.data[key]
	if !exists {
		return false
	}

	item.expires = time.Time{}
	store.data[key] = item
	return true
}
