package store

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {

	kvStore := NewKeyValueStore()

	kvStore.Set("key1", "value1", 5*time.Second)

	value, exist := kvStore.Get("key1")
	if !exist || value != "value1" {
		t.Errorf("TestSetAndGet: Expected value1, got %s (exist: %v)", value, exist)
	}

}

func TestExpireKey(t *testing.T) {
	kvStore := NewKeyValueStore()

	kvStore.Set("key1", "value1", 1*time.Second)

	value, exist := kvStore.Get("key1")
	if !exist || value != "value1" {
		t.Errorf("TestExpireKey: Expected value1, got %s (exist: %v)", value, exist)
	}
	time.Sleep(2 * time.Second)
	_, exist = kvStore.Get("key2")
	if exist {
		t.Errorf("TestExpireKey: Expected key2 to be expired")
	}
}

func TestDeleteKey(t *testing.T) {
  kvStore := NewKeyValueStore()

  kvStore.Set("key3", "value3", 0) 

  kvStore.Delete("key3")

  _, exist := kvStore.Get("key3")
  if exist {
      t.Errorf("TestDeleteKey: Expected key3 to be deleted")
  }
}

func TestExistKey(t *testing.T){
	kvStore := NewKeyValueStore()

	exist := kvStore.Exist("key11")
	if exist {
		t.Errorf("TestExistKey: Expected key11 exsitance check should return false")
	}

	kvStore.Set("key11","value1",0)
	exist = kvStore.Exist("key11")
	if !exist {
		t.Errorf("TestExistKey: Expected key11 exsitance check should return true")
	}

}

func TestTTL(t *testing.T) {
    store := NewKeyValueStore()

    store.Set("key1", "value1", 5*time.Second)
    ttl := store.TTL("key1")
    if ttl <= 0 {
        t.Errorf("Expected positive TTL for key1, got %d", ttl)
    }

    store.Set("key2", "value2", 0)
    ttl2 := store.TTL("key2")
    if ttl2 != -1 {
        t.Errorf("Expected TTL of -1 for key2, got %d", ttl2)
    }

    ttl3 := store.TTL("nonExistentKey")
    if ttl3 != -2 {
        t.Errorf("Expected TTL of -2 for non-existent key, got %d", ttl3)
    }

    store.Set("key3", "value3", 1*time.Second)
    time.Sleep(2 * time.Second) 
    ttl4 := store.TTL("key3")
    if ttl4 != 0 {
        t.Errorf("Expected TTL of 0 for expired key3, got %d", ttl4)
    }
}