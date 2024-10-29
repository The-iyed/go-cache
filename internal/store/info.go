package store

import "fmt"

func (store *KeyValueStore) Info() string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var totalSize int
	for _, item := range store.data {
		totalSize += len(item.value)
	}

	info := fmt.Sprintf(
		"ICache Server\n"+
			"Number of Keys: %d\n"+
			"Total Size: %d bytes\n"+
			"Memory Usage: %.2f KB\n",
		len(store.data),
		totalSize,
		float64(totalSize)/1024,
	)

	return info
}
