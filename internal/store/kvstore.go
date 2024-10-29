package store

import (
	"sync"
	"time"
)


type KeyValueStore struct {
	data map[string]Item
	mu   sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	store := &KeyValueStore{
		data: make(map[string]Item),
	}
	go store.Clean()
	return store
}

func (store *KeyValueStore) Clean() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		expired := []string{}
		now := time.Now()

		store.mu.RLock()
		for key, item := range store.data {
			if !item.expires.IsZero() && now.After(item.expires) {
				expired = append(expired, key)
			}
		}
		store.mu.RUnlock()

		if len(expired) > 0 {
			store.mu.Lock()
			for _, key := range expired {
				delete(store.data, key)
			}
			store.mu.Unlock()
		}
	}
}