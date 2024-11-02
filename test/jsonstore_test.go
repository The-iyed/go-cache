package test

import (
	"testing"
	"time"

	"github.com/go-redis-v1/internal/jsonstore"
)

func TestSetJSON(t *testing.T) {
	store := jsonstore.NewJSONStore()
	data := map[string]interface{}{"name": "John", "age": 30}

	err := store.SetJSON("user1", data, 0)
	if err != nil {
		t.Fatalf("SetJSON failed: %v", err)
	}

	var result map[string]interface{}
	err = store.GetJSON("user1", &result)
	if err != nil {
		t.Fatalf("GetJSON failed after SetJSON: %v", err)
	}

	if result["name"] != "John" || result["age"] != float64(30) {
		t.Errorf("Unexpected data retrieved: %v", result)
	}
}

func TestSetJSONWithTTL(t *testing.T) {
	store := jsonstore.NewJSONStore()
	data := map[string]interface{}{"name": "John"}

	err := store.SetJSON("user2", data, 1*time.Second)
	if err != nil {
		t.Fatalf("SetJSON failed: %v", err)
	}

	time.Sleep(2 * time.Second)

	var result map[string]interface{}
	err = store.GetJSON("user2", &result)
	if err == nil {
		t.Error("Expected key to be expired, but it was found")
	}
}

func TestDeleteJSON(t *testing.T) {
	store := jsonstore.NewJSONStore()
	data := map[string]interface{}{"name": "Jane"}

	store.SetJSON("user3", data, 0)
	deleted := store.DeleteJSON("user3")
	if !deleted {
		t.Error("Expected key to be deleted")
	}

	var result map[string]interface{}
	err := store.GetJSON("user3", &result)
	if err == nil {
		t.Error("Expected key to be not found after deletion")
	}
}

func TestUpdateJSON(t *testing.T) {
	store := jsonstore.NewJSONStore()
	data := map[string]interface{}{"name": "Doe", "age": 25}

	store.SetJSON("user4", data, 0)
	err := store.UpdateJSON("user4", "age", 26)
	if err != nil {
		t.Fatalf("UpdateJSON failed: %v", err)
	}

	var result map[string]interface{}
	err = store.GetJSON("user4", &result)
	if err != nil {
		t.Fatalf("GetJSON failed after UpdateJSON: %v", err)
	}

	if result["age"] != float64(26) {
		t.Errorf("Expected age to be updated to 26, got %v", result["age"])
	}
}

func TestTTLJSON(t *testing.T) {
	store := jsonstore.NewJSONStore()
	data := map[string]interface{}{"name": "Test"}

	store.SetJSON("user5", data, 2*time.Second)

	ttl, err := store.TTL("user5")
	if err != nil || ttl <= 0 {
		t.Fatalf("Expected TTL to be set, but got error: %v", err)
	}

	time.Sleep(3 * time.Second)

	ttl, err = store.TTL("user5")
	if err == nil {
		t.Error("Expected TTL to fail after expiration, but it did not")
	}
}
