package test

import (
	"log"
	"testing"
	"time"

	"github.com/go-redis-v1/internal/store"
)



func TestTTL(t *testing.T) {
	store := store.NewKeyValueStore()

	store.Set("key1", "value1", 5*time.Second)
	ttl := store.TTL("key1")
	if ttl <= 0 {
		t.Errorf("Expected positive TTL for key1, got %d", ttl)
	} else {
		store.Set("key2", "value2", 0)
		ttl2 := store.TTL("key2")
		if ttl2 != -1 {
			t.Errorf("Expected TTL of -1 for key2, got %d", ttl2)
		} else {
			ttl3 := store.TTL("nonExistentKey")
			if ttl3 != -2 {
				t.Errorf("Expected TTL of -2 for non-existent key, got %d", ttl3)
			} else {
				store.Set("key3", "value3", 1*time.Second)
				time.Sleep(2 * time.Second)
				ttl4 := store.TTL("key3")
				if ttl4 != 0 {
					t.Errorf("Expected TTL of 0 for expired key3, got %d", ttl4)
				} else {
					log.Println("TestTTL: Passed")
				}
			}
		}
	}
}
