package store

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
