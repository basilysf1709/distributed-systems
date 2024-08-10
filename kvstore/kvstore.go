package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// KVStore represents our key-value store
type KVStore struct {
	mu       sync.RWMutex
	store    map[string]string
	filePath string
}

// NewKVStore creates a new KVStore
func NewKVStore(filePath string) *KVStore {
	kv := &KVStore{
		store:    make(map[string]string),
		filePath: filePath,
	}
	kv.load() // Load existing data from file
	return kv
}

// load reads the JSON file and populates the store
func (kv *KVStore) load() {
	file, err := ioutil.ReadFile(kv.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	err = json.Unmarshal(file, &kv.store)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
	}
}

// save writes the store to the JSON file
func (kv *KVStore) save() {
	// Remove the RLock here, as it's causing the deadlock
	file, err := json.MarshalIndent(kv.store, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}

	err = ioutil.WriteFile(kv.filePath, file, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

// Set adds or updates a key-value pair in the store and saves to file
func (kv *KVStore) Set(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = value
	kv.save() // This is now safe because we hold the write lock
	fmt.Printf("Set %s to %s\n", key, value)
}

// Get retrieves a value from the store given a key
func (kv *KVStore) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	value, exists := kv.store[key]
	return value, exists
}

// Delete removes a key-value pair from the store and saves to file
func (kv *KVStore) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	delete(kv.store, key)
	kv.save() // This is now safe because we hold the write lock
	fmt.Printf("Deleted %s\n", key)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  set <key> <value>")
		fmt.Println("  get <key>")
		fmt.Println("  delete <key>")
		return
	}

	kv := NewKVStore("kvstore.json")

	command := os.Args[1]
	key := os.Args[2]

	switch command {
	case "set":
		if len(os.Args) != 4 {
			fmt.Println("Usage: set <key> <value>")
			return
		}
		value := os.Args[3]
		kv.Set(key, value)

	case "get":
		if value, exists := kv.Get(key); exists {
			fmt.Printf("%s: %s\n", key, value)
		} else {
			fmt.Printf("Key '%s' not found\n", key)
		}

	case "delete":
		kv.Delete(key)

	default:
		fmt.Println("Unknown command. Use 'set', 'get', or 'delete'.")
	}
}