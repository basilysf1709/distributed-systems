package kvstore

import (
	"github.com/basilysf1709/distributed-systems/kvstore/internal/store"
)

type KVStore struct {
	store *store.Store
}

func NewKVStore(persistence store.Persistence) *KVStore {
	return &KVStore{
		store: store.NewStore(persistence),
	}
}

func (kvs *KVStore) Get(key string) (string, bool) {
	return kvs.store.Get(key)
}

func (kvs *KVStore) Set(key, value string) error {
	return kvs.store.Set(key, value)
}

func (kvs *KVStore) Delete(key string) error {
	return kvs.store.Delete(key)
}