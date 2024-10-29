package store

import "time"

func (kv *KeyValueStore) MSET(pairs ...string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if len(pairs)%2 != 0 {
		return
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		kv.data[key] = Item{
			value:   value,
			expires: time.Time{},
		}
	}
}
