package store

import (
	"fmt"
	"regexp"
	"strings"
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

func (store *KeyValueStore) Keys(pattern string) []string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var keys []string

	useRegex := strings.Contains(pattern, "*")
	var re *regexp.Regexp
	if useRegex {
		re = regexp.MustCompile("^" + strings.ReplaceAll(regexp.QuoteMeta(pattern), "\\*", ".*") + "$")
	}

	for key := range store.data {
		if !useRegex && pattern == "*" || useRegex && re.MatchString(key) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (store *KeyValueStore) Exist(key string) bool {
	store.mu.RLock()
	defer store.mu.RUnlock()

	_, exist := store.data[key]

	return exist

}

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

func (store *KeyValueStore) FlushAll() {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data = make(map[string]Item)
}

func (store *KeyValueStore) Info() string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var totalSize int
	for _, item := range store.data {
		totalSize += len(item.value)
	}

	info := fmt.Sprintf(
		"ICache Server\n"+
			"Number of Keys: %d\n"+
			"Total Size: %d bytes\n"+
			"Memory Usage: %.2f KB\n",
		len(store.data),
		totalSize,
		float64(totalSize)/1024,
	)

	return info
}

func (store *KeyValueStore) Ping() string {
	return "PONG"
}

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

func (kv *KeyValueStore) MSET(pairs ...string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if len(pairs)%2 != 0 {
		return
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		kv.data[key] = Item{
			value:   value,
			expires: time.Time{},
		}
	}
}

func (store *KeyValueStore) MGET(keys ...string) []string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	results := make([]string, len(keys))

	for i, key := range keys {
		item, exists := store.data[key]
		if !exists || item.IsExpired() {
			results[i] = "(nil)"
		} else {
			results[i] = item.value
		}
	}
	return results
}

func (store *KeyValueStore) GETSET(key string, value string) (string, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()

	item, exists := store.data[key]
	if exists && !item.IsExpired() {
		oldValue := item.value
		store.data[key] = Item{value: value, expires: item.expires}
		return oldValue, true
	}
	store.data[key] = Item{value: value, expires: time.Time{}}
	return "", false
}

func (kvs *KeyValueStore) Update(key string, value string) string {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	if item, exists := kvs.data[key]; exists {
		kvs.data[key] = Item{value: value, expires: item.expires}
		return item.value
	}
	return ""
}


func (kvs *KeyValueStore) GetSet(key string, value string) string {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	oldValue := ""
	if item, exists := kvs.data[key]; exists {
		oldValue = item.value
	}
	kvs.data[key] = Item{value: value, expires: time.Time{}}
	return oldValue
}

