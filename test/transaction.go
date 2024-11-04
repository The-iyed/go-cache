package test

import (
	"testing"

	"github.com/go-redis-v1/internal/transaction"
)

func TestStartTransaction(t *testing.T) {
	tx := transaction.NewTransaction()
	err := tx.StartTransaction()
	if err != nil {
		t.Fatalf("Unexpected error on StartTransaction: %s", err)
	}
	if !tx.IsActive {
		t.Fatal("Expected transaction to be active after StartTransaction")
	}
	err = tx.StartTransaction()
	if err == nil || err.Error() != "transaction start failed: active transaction already running" {
		t.Fatalf("Expected error for starting an active transaction, got: %v", err)
	}
}

func TestAddCommand(t *testing.T) {
	tx := transaction.NewTransaction()
	tx.StartTransaction()
	tx.AddCommand("SET", []string{"SET", "key1", "value1"})

	if len(tx.Commands) != 1 {
		t.Fatalf("Expected 1 command in transaction, got %d", len(tx.Commands))
	}

	if tx.Commands[0].Name != "SET" || tx.Commands[0].Args[1] != "key1" || tx.Commands[0].Args[2] != "value1" {
		t.Fatalf("Expected command to be 'SET key1 value1', got '%s %v'", tx.Commands[0].Name, tx.Commands[0].Args)
	}
}

func TestAbortTransaction(t *testing.T) {
	tx := transaction.NewTransaction()
	tx.StartTransaction()
	tx.AddCommand("SET", []string{"SET", "key1", "value1"})
	err := tx.AbortTransaction()
	if err != nil {
		t.Fatalf("Unexpected error on AbortTransaction: %s", err)
	}
	if tx.IsActive {
		t.Fatal("Expected transaction to be inactive after AbortTransaction")
	}
	if len(tx.Commands) != 0 {
		t.Fatalf("Expected no commands after AbortTransaction, got %d", len(tx.Commands))
	}
	err = tx.AbortTransaction()
	if err == nil || err.Error() != "transaction abort failed: no active transaction" {
		t.Fatalf("Expected error for aborting without an active transaction, got: %v", err)
	}
}
