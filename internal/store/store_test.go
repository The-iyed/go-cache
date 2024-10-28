package store


import (
  "testing"
  "time"
)


func TestSetAndGet(t *testing.T){
    
  kvStore := NewKeyValueStore()

  kvStore.Set("key1", "value1", 5*time.Second)
  
  value , exist := kvStore.Get("key1")
  if !exist || value != "value1" {
    t.Errorf("Expected value1, got %s (exist: %v)", value, exist)
  }

}
