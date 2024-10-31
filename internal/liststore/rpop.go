package liststore

func (ls *ListStore) RPOP(key string) (string, bool) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	list, exists := ls.lists[key]
	if !exists || len(list) == 0 {
		return "", false
	}
	value := list[len(list)-1]
	ls.lists[key] = list[:len(list)-1]
	return value, true
}
