package main

import (
	"fmt"
	"os"

	"github.com/basilysf1709/distributed-systems/kvstore/internal/cli"
	"github.com/basilysf1709/distributed-systems/kvstore/internal/store"
	"github.com/basilysf1709/distributed-systems/kvstore/pkg/kvstore"
)

func main() {
	persistence := &store.JSONPersistence{FilePath: "kvstore.json"}
	kvs := kvstore.NewKVStore(persistence)
	
	if err := cli.Run(os.Args, kvs); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}