package store

import (
	"sync"
)

type Store struct {
	mu           sync.RWMutex
	data         map[string]string
	persistence  Persistence
}

type Persistence interface {
	Save(data map[string]string) error
	Load() (map[string]string, error)
}

func NewStore(persistence Persistence) *Store {
	s := &Store{
		data:        make(map[string]string),
		persistence: persistence,
	}
	if data, err := s.persistence.Load(); err == nil {
		s.data = data
	}
	return s
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return s.persistence.Save(s.data)
}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return s.persistence.Save(s.data)
}