package liststore

import (
	"sync"
)

type ListStore struct {
	mu    sync.RWMutex
	lists map[string][]string
}

func NewListStore() *ListStore {
	return &ListStore{
		lists: make(map[string][]string),
	}
}
