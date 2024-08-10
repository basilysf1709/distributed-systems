package store

import (
	"fmt"
	"sync"
)

type Store struct {
	mu          sync.RWMutex
	data        map[string]string
	persistence Persistence
}

type Persistence interface {
	Save(data map[string]string) error
	Load() (map[string]string, error)
}

func NewStore(persistence Persistence) *Store {
	fmt.Println("Initializing new Store...")
	s := &Store{
		data:        make(map[string]string),
		persistence: persistence,
	}
	fmt.Println("Loading data from persistence...")
	if data, err := s.persistence.Load(); err == nil {
		s.data = data
		fmt.Printf("Loaded %d key-value pairs from persistence\n", len(data))
	} else {
		fmt.Printf("Failed to load data from persistence: %v\n", err)
	}
	return s
}

func (s *Store) Get(key string) (string, bool) {
	fmt.Printf("Getting value for key '%s'...\n", key)
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	if ok {
		fmt.Printf("Found value for key '%s'\n", key)
	} else {
		fmt.Printf("Key '%s' not found in store\n", key)
	}
	return value, ok
}

func (s *Store) Set(key, value string) error {
	fmt.Printf("Setting key '%s' to value '%s'...\n", key, value)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	fmt.Println("Saving updated data to persistence...")
	err := s.persistence.Save(s.data)
	if err != nil {
		fmt.Printf("Failed to save data to persistence: %v\n", err)
	} else {
		fmt.Println("Data saved successfully")
	}
	return err
}

func (s *Store) Delete(key string) error {
	fmt.Printf("Deleting key '%s'...\n", key)
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	fmt.Println("Saving updated data to persistence...")
	err := s.persistence.Save(s.data)
	if err != nil {
		fmt.Printf("Failed to save data to persistence: %v\n", err)
	} else {
		fmt.Println("Data saved successfully")
	}
	return err
}

func (s *Store) List() (map[string]string, error) {
	fmt.Println("Listing all key-value pairs...")
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Create a copy of the data to return
	dataCopy := make(map[string]string)
	for k, v := range s.data {
		dataCopy[k] = v
	}
	fmt.Printf("Found %d key-value pairs\n", len(dataCopy))
	return dataCopy, nil
}