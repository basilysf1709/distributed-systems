package main

import (
	"fmt"
	"strings"
	"sync"
)

// KeyValue holds a key-value pair
type KeyValue struct {
	Key   string
	Value int
}

// MapFunc takes a string and returns a slice of KeyValue pairs
func MapFunc(input string) []KeyValue {
	fmt.Printf("Mapping input: %s\n", input)
	words := strings.Fields(input)
	result := make([]KeyValue, len(words))
	for i, word := range words {
		result[i] = KeyValue{word, 1}
	}
	fmt.Printf("Map result: %v\n", result)
	return result
}

// ReduceFunc takes a key and a slice of values, returns a single KeyValue pair
func ReduceFunc(key string, values []int) KeyValue {
	fmt.Printf("Reducing key: %s with values: %v\n", key, values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	result := KeyValue{key, sum}
	fmt.Printf("Reduce result: %v\n", result)
	return result
}

// MapReduce orchestrates the MapReduce process
func MapReduce(inputs []string, mapF func(string) []KeyValue, reduceF func(string, []int) KeyValue) []KeyValue {
	fmt.Println("Starting MapReduce process...")

	// Map phase
	fmt.Println("Starting Map phase...")
	intermediate := []KeyValue{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, input := range inputs {
		wg.Add(1)
		go func(inp string) {
			defer wg.Done()
			mapped := mapF(inp)
			mu.Lock()
			intermediate = append(intermediate, mapped...)
			mu.Unlock()
		}(input)
	}
	wg.Wait()
	fmt.Printf("Map phase complete. Intermediate results: %v\n", intermediate)

	// Shuffle phase
	fmt.Println("Starting Shuffle phase...")
	groups := make(map[string][]int)
	for _, kv := range intermediate {
		groups[kv.Key] = append(groups[kv.Key], kv.Value)
	}
	fmt.Printf("Shuffle phase complete. Grouped results: %v\n", groups)

	// Reduce phase
	fmt.Println("Starting Reduce phase...")
	var results []KeyValue
	for key, values := range groups {
		results = append(results, reduceF(key, values))
	}
	fmt.Printf("Reduce phase complete. Final results: %v\n", results)

	return results
}

func main() {
	inputs := []string{
		"the quick brown fox",
		"the fox jumped over",
		"the lazy dog",
	}

	fmt.Println("Input strings:", inputs)

	results := MapReduce(inputs, MapFunc, ReduceFunc)

	fmt.Println("\nFinal word counts:")
	for _, kv := range results {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}