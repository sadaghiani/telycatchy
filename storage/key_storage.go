package storage

import (
	"errors"
	"sync"
)

// DataStore is an interface to store Datas
type DataStore interface {
	// Save saves a Data to the store
	Set(key string, Data *Result)
	// Find find Data from store
	Get(key string) (*Result, error)
}

// InMemoryDataStore stores Datas in memory
type InMemoryDataStore struct {
	mutex sync.RWMutex
	Data  map[string]*Result
}

// NewInMemoryDataStore returns a new in-memory Data store
func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{
		Data: make(map[string]*Result),
	}
}

// Save saves a Data to the store
func (store *InMemoryDataStore) Set(key string, Data *Result) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.Data[key] = Data
}

// Get return Data from store
func (store *InMemoryDataStore) Get(key string) (*Result, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if value, ok := store.Data[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("Data not found")
	}
}
