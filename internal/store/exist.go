package store

func (store *KeyValueStore) Exist(key string) bool {
	store.mu.RLock()
	defer store.mu.RUnlock()

	_, exist := store.data[key]

	return exist

}
