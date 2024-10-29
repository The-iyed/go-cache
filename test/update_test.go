package test

import (
	"testing"

	"github.com/go-redis-v1/internal/store"
)

func TestUpdate(t *testing.T) {
	kvStore := store.NewKeyValueStore()
	kvStore.Set("key1", "value1", 0)

	oldValue := kvStore.Update("key1", "value2")
	if oldValue != "value1" {
		t.Errorf("TestUpdate: Expected old value 'value1', got '%s'", oldValue)
	}

	value, exist := kvStore.Get("key1")
	if !exist || value != "value2" {
		t.Errorf("TestUpdate: Expected value 'value2', got '%s' (exist: %v)", value, exist)
	}

	oldValue = kvStore.Update("key2", "value3")
	if oldValue != "" {
		t.Errorf("TestUpdate: Expected old value to be '', got '%s'", oldValue)
	}
}