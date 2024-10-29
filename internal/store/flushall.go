package store

func (store *KeyValueStore) FlushAll() {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data = make(map[string]Item)
}
