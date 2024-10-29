package store

func (kvs *KeyValueStore) Update(key string, value string) string {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	if item, exists := kvs.data[key]; exists {
		kvs.data[key] = Item{value: value, expires: item.expires}
		return item.value
	}
	return ""
}
