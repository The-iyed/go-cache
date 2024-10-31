package liststore

func (ls *ListStore) LRANGE(key string, start, stop int) []string {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	list, exists := ls.lists[key]
	if !exists {
		return nil
	}

	if start < 0 {
		start = len(list) + start
	}
	if stop < 0 {
		stop = len(list) + stop
	}
	if start < 0 {
		start = 0
	}
	if stop > len(list) {
		stop = len(list)
	}
	if start > stop {
		return nil
	}
	return list[start:stop]
}
