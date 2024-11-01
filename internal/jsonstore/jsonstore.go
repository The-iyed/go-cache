package jsonstore

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type JSONStore struct {
	mu         sync.RWMutex
	jsonData   map[string][]byte
	expiration map[string]time.Time
}

func NewJSONStore() *JSONStore {
	return &JSONStore{
		jsonData:   make(map[string][]byte),
		expiration: make(map[string]time.Time),
	}
}

func (js *JSONStore) SetJSON(key string, value interface{}, ttl time.Duration) error {
	js.mu.Lock()
	defer js.mu.Unlock()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	js.jsonData[key] = data
	if ttl > 0 {
		js.expiration[key] = time.Now().Add(ttl)
	} else {
		delete(js.expiration, key)
	}
	return nil
}

func (js *JSONStore) GetJSON(key string, dest interface{}) error {
	js.mu.RLock()
	defer js.mu.RUnlock()
	data, exists := js.jsonData[key]
	if !exists {
		return errors.New("key not found or expired")
	}
	if exp, ok := js.expiration[key]; ok && exp.Before(time.Now()) {
		delete(js.jsonData, key)
		delete(js.expiration, key)
		return errors.New("key not found or expired")
	}
	return json.Unmarshal(data, dest)
}

func (js *JSONStore) DeleteJSON(key string) bool {
	js.mu.Lock()
	defer js.mu.Unlock()
	_, exists := js.jsonData[key]
	if exists {
		delete(js.jsonData, key)
		delete(js.expiration, key)
	}
	return exists
}

func (js *JSONStore) UpdateJSON(key string, field string, value interface{}) error {
	js.mu.Lock()
	defer js.mu.Unlock()
	data, exists := js.jsonData[key]
	if !exists {
		return errors.New("key not found")
	}
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}
	jsonData[field] = value
	updatedData, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	js.jsonData[key] = updatedData
	return nil
}

func (js *JSONStore) TTL(key string) (time.Duration, error) {
	js.mu.RLock()
	defer js.mu.RUnlock()
	exp, exists := js.expiration[key]
	if !exists || exp.Before(time.Now()) {
		return -1, errors.New("key not found or expired")
	}
	return time.Until(exp), nil
}
