package main

import (
	"fmt"
	"os"

	"github.com/yourusername/kvstore/internal/cli"
	"github.com/yourusername/kvstore/internal/store"
	"github.com/yourusername/kvstore/pkg/kvstore"
)

func main() {
	persistence := &store.JSONPersistence{FilePath: "kvstore.json"}
	kvs := kvstore.NewKVStore(persistence)
	
	if err := cli.Run(os.Args, kvs); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}