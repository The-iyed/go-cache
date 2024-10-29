package store

import (
	"time"
)

func (store *KeyValueStore) Get(key string) (string, bool) {
	store.mu.RLock()
	item, exist := store.data[key]
	store.mu.RUnlock()

	if !exist {
		return "", false
	}

	if !item.expires.IsZero() && time.Now().After(item.expires) {
		store.mu.Lock()
		delete(store.data, key)
		store.mu.Unlock()
		return "", false
	}

	return item.value, true
}
