package store

import "time"

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
