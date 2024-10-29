package test

import (
	"log"
	"testing"

	"github.com/go-redis-v1/internal/store"
)

func TestDeleteKey(t *testing.T) {
	kvStore := store.NewKeyValueStore()

	kvStore.Set("key3", "value3", 0)

	kvStore.Delete("key3")

	_, exist := kvStore.Get("key3")
	if exist {
		t.Errorf("TestDeleteKey: Expected key3 to be deleted")
	} else {
		log.Println("TestDeleteKey: Passed")
	}
}
