package test

import (
	"log"
	"testing"
	"time"

	"github.com/go-redis-v1/internal/store"
)

func TestPersist(t *testing.T) {
	kvStore := store.NewKeyValueStore()

	kvStore.Set("key1", "value1", 5*time.Second)
	kvStore.Persist("key1")
	ttl := kvStore.TTL("key1")
	if ttl != -1 {
		t.Errorf("TestPersist: Expected TTL of -1 for key1, got %d", ttl)
	} else {
		log.Println("TestPersist: Passed")
	}
}
