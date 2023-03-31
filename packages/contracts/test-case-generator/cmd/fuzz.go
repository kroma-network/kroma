package main

import (
	"flag"
	"log"

	t "github.com/wemixkanvas/kanvas/packages/contracts/test-case-generator/trie"
)

// Mode enum
const (
	// Enables the `trie` fuzzer
	trie string = "trie"
)

func main() {
	mode := flag.String("m", "", "Fuzzer mode")
	variant := flag.String("v", "", "Mode variant")
	flag.Parse()

	if len(*mode) < 1 {
		log.Fatal("Must pass a mode for the fuzzer!")
	}

	switch *mode {
	case trie:
		t.FuzzTrie(*variant)
	default:
		log.Fatal("Invalid mode!")
	}
}
