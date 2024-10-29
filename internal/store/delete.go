package store

func (store *KeyValueStore) Delete(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.data, key)
}
