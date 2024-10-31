package liststore

func (ls *ListStore) LPUSH(key string, values ...string) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.lists[key] = append(values, ls.lists[key]...)
}
