package test

import (
	"testing"

	"github.com/go-redis-v1/internal/liststore"
)

func TestLPUSH(t *testing.T) {
	store := liststore.NewListStore()
	store.LPUSH("mylist", "value1")
	store.LPUSH("mylist", "value2")

	expected := []string{"value2", "value1"}
	actual := store.LRANGE("mylist", 0, -1)
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
		}
	}
}

func TestRPUSH(t *testing.T) {
	store := liststore.NewListStore()
	store.RPUSH("mylist", "value1")
	store.RPUSH("mylist", "value2")

	expected := []string{"value1", "value2"}
	actual := store.LRANGE("mylist", 0, -1)
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
		}
	}
}

func TestLPOP(t *testing.T) {
	store := liststore.NewListStore()
	store.LPUSH("mylist", "value1")
	store.LPUSH("mylist", "value2")

	poppedValue, ok := store.LPOP("mylist")
	if !ok {
		t.Errorf("Expected to pop a value, but got none")
	}
	if poppedValue != "value2" {
		t.Errorf("Expected value2, got %s", poppedValue)
	}

	expected := []string{"value1"}
	actual := store.LRANGE("mylist", 0, -1)
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
		}
	}
}

func TestRPOP(t *testing.T) {
	store := liststore.NewListStore()
	store.RPUSH("mylist", "value1")
	store.RPUSH("mylist", "value2")

	poppedValue, ok := store.RPOP("mylist")
	if !ok {
		t.Errorf("Expected to pop a value, but got none")
	}
	if poppedValue != "value2" {
		t.Errorf("Expected value2, got %s", poppedValue)
	}

	expected := []string{"value1"}
	actual := store.LRANGE("mylist", 0, -1)
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
		}
	}
}

func TestLRANGE(t *testing.T) {
	store := liststore.NewListStore()
	store.LPUSH("mylist", "value1")
	store.LPUSH("mylist", "value2")
	store.LPUSH("mylist", "value3")

	actual := store.LRANGE("mylist", 0, 2)

	expected := []string{"value3", "value2", "value1"}
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
		}
	}
}

func TestLRANGEInvalidIndex(t *testing.T) {
	store := liststore.NewListStore()
	store.LPUSH("mylist", "value1")
	store.LPUSH("mylist", "value2")

	actual := store.LRANGE("mylist", -1, -1)

	if len(actual) != 0 {
		t.Errorf("Expected empty result for invalid range, got %v", actual)
	}
}

func TestLPOPEmpty(t *testing.T) {
	store := liststore.NewListStore()

	poppedValue, ok := store.LPOP("emptylist")
	if ok {
		t.Errorf("Expected to not pop a value, but got %s", poppedValue)
	}
}

func TestRPOPEmpty(t *testing.T) {
	store := liststore.NewListStore()

	poppedValue, ok := store.RPOP("emptylist")
	if ok {
		t.Errorf("Expected to not pop a value, but got %s", poppedValue)
	}
}
