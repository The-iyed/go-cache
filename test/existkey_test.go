package test

import (
	"log"
	"testing"

	"github.com/go-redis-v1/internal/store"
)

func TestExistKey(t *testing.T) {
	kvStore := store.NewKeyValueStore()

	exist := kvStore.Exist("key11")
	if exist {
		t.Errorf("TestExistKey: Expected key11 existence check should return false")
	} else {
		kvStore.Set("key11", "value1", 0)
		exist = kvStore.Exist("key11")
		if !exist {
			t.Errorf("TestExistKey: Expected key11 existence check should return true")
		} else {
			log.Println("TestExistKey: Passed")
		}
	}
}
