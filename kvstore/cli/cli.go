package cli

import (
	"fmt"
	"github.com/basilysf1709/distributed-systems/kvstore/pkg/kvstore"
)

func Run(args []string, kvs *kvstore.KVStore) error {
	if len(args) < 3 {
		return fmt.Errorf("Usage:\n  set <key> <value>\n  get <key>\n  delete <key>")
	}

	command := args[1]
	key := args[2]

	switch command {
	case "set":
		if len(args) != 4 {
			return fmt.Errorf("Usage: set <key> <value>")
		}
		value := args[3]
		if err := kvs.Set(key, value); err != nil {
			return err
		}
		fmt.Printf("Set %s to %s\n", key, value)

	case "get":
		if value, ok := kvs.Get(key); ok {
			fmt.Printf("%s: %s\n", key, value)
		} else {
			fmt.Printf("Key '%s' not found\n", key)
		}

	case "delete":
		if err := kvs.Delete(key); err != nil {
			return err
		}
		fmt.Printf("Deleted %s\n", key)

	default:
		return fmt.Errorf("Unknown command. Use 'set', 'get', or 'delete'")
	}

	return nil
}