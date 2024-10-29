package test

import (
	"log"
	"testing"
	"time"

	"github.com/go-redis-v1/internal/store"
)

func TestExpireKey(t *testing.T) {
	kvStore := store.NewKeyValueStore()

	kvStore.Set("key1", "value1", 1*time.Second)

	value, exist := kvStore.Get("key1")
	if !exist || value != "value1" {
		t.Errorf("TestExpireKey: Expected value1, got %s (exist: %v)", value, exist)
	} else {
		time.Sleep(2 * time.Second)
		_, exist = kvStore.Get("key1")
		if exist {
			t.Errorf("TestExpireKey: Expected key1 to be expired")
		} else {
			log.Println("TestExpireKey: Passed")
		}
	}
}
