package liststore

func (ls *ListStore) LPOP(key string) (string, bool) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	list, exists := ls.lists[key]
	if !exists || len(list) == 0 {
		return "", false
	}
	value := list[0]
	ls.lists[key] = list[1:]
	return value, true
}
