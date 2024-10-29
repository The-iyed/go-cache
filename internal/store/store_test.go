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