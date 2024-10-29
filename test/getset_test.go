package test

import (
	"log"
	"testing"
	"time"

	"github.com/go-redis-v1/internal/store"
)

func TestSetAndGet(t *testing.T) {
	kvStore := store.NewKeyValueStore()

	kvStore.Set("key1", "value1", 5*time.Second)

	value, exist := kvStore.Get("key1")
	if !exist || value != "value1" {
		t.Errorf("TestSetAndGet: Expected value1, got %s (exist: %v)", value, exist)
	} else {
		log.Println("TestSetAndGet: Passed")
	}
}


func TestGetSet(t *testing.T) {
	kvStore := store.NewKeyValueStore()
	kvStore.Set("key1", "value1", 0)

	oldValue := kvStore.GetSet("key1", "newValue")
	if oldValue != "value1" {
		t.Errorf("TestGetSet: Expected old value to be 'value1', got '%s'", oldValue)
	}

	newValue, exist := kvStore.Get("key1")
	if !exist || newValue != "newValue" {
		t.Errorf("TestGetSet: Expected new value to be 'newValue', got '%s' (exist: %v)", newValue, exist)
	}

	oldValue = kvStore.GetSet("key2", "value2")
	if oldValue != "" {
		t.Errorf("TestGetSet: Expected old value to be '', got '%s'", oldValue)
	}
}
