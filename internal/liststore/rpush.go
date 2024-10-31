package liststore

func (ls *ListStore) RPUSH(key string, values ...string) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.lists[key] = append(ls.lists[key], values...)
}
