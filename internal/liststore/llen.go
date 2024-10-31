package liststore

func (ls *ListStore) LLEN(key string) int {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	return len(ls.lists[key])
}
