package store

import (
    "sync"
    "time"
)

type Item struct {
    value   string
    expires time.Time
}

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

func (store *KeyValueStore) Delete(key string) {
    store.mu.Lock()
    defer store.mu.Unlock()
    delete(store.data, key)
}

