package cli

import (
	"fmt"
	"github.com/basilysf1709/distributed-systems/kvstore/pkg/kvstore"
)

func Run(args []string, kvs *kvstore.KVStore) error {
	fmt.Println("Starting KV Store CLI...")

	if len(args) < 3 {
		return fmt.Errorf("Usage:\n  set <key> <value>\n  get <key>\n  delete <key>\n  list")
	}

	command := args[1]
	fmt.Printf("Executing command: %s\n", command)

	switch command {
	case "set":
		if len(args) != 4 {
			return fmt.Errorf("Usage: set <key> <value>")
		}
		key := args[2]
		value := args[3]
		fmt.Printf("Attempting to set key '%s' to value '%s'...\n", key, value)
		if err := kvs.Set(key, value); err != nil {
			return fmt.Errorf("Failed to set value: %v", err)
		}
		fmt.Printf("Successfully set %s to %s\n", key, value)

	case "get":
		key := args[2]
		fmt.Printf("Attempting to get value for key '%s'...\n", key)
		if value, ok := kvs.Get(key); ok {
			fmt.Printf("Successfully retrieved value. %s: %s\n", key, value)
		} else {
			fmt.Printf("Key '%s' not found in the store\n", key)
		}

	case "delete":
		key := args[2]
		fmt.Printf("Attempting to delete key '%s'...\n", key)
		if err := kvs.Delete(key); err != nil {
			return fmt.Errorf("Failed to delete key: %v", err)
		}
		fmt.Printf("Successfully deleted key %s\n", key)

	case "list":
		fmt.Println("Listing all key-value pairs:")
		pairs, err := kvs.List()
		if err != nil {
			return fmt.Errorf("Failed to list key-value pairs: %v", err)
		}
		if len(pairs) == 0 {
			fmt.Println("The store is empty.")
		} else {
			for k, v := range pairs {
				fmt.Printf("%s: %s\n", k, v)
			}
		}

	default:
		return fmt.Errorf("Unknown command '%s'. Use 'set', 'get', 'delete', or 'list'", command)
	}

	fmt.Println("Command executed successfully.")
	return nil
}