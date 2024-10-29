package store

import (
	"time"
)



func (store *KeyValueStore) Set(key, value string, ttl time.Duration) {
	store.mu.Lock()
	defer store.mu.Unlock()

	expires := time.Time{}
	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}

	store.data[key] = Item{
		value:   value,
		expires: expires,
	}
}
