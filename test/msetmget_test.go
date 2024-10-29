package test

import (
	"testing"

	"github.com/go-redis-v1/internal/store"
)

func TestMSETAndMGET(t *testing.T) {
	kvStore := store.NewKeyValueStore()
	kvStore.MSET("key1", "value1", "key2", "value2", "key3", "value3")

	values := kvStore.MGET("key1", "key2", "key3")
	expected := []string{"value1", "value2", "value3"}

	for i, expectedValue := range expected {
		if values[i] != expectedValue {
			t.Errorf("TestMSETAndMGET: Expected %s, got %s", expectedValue, values[i])
		}
	}

	values = kvStore.MGET("nonExistentKey")
	if values[0] != "(nil)" {
		t.Errorf("TestMSETAndMGET: Expected (nil), got %s", values[0])
	}
}
