package internal

import (
    "sync"
)

var (
    mu        sync.Mutex
)

type URLShortener struct {
    Urls map[string]string
    UrlHashes map[string]string
    DomainFreq map[string]int
}

// Pair represents a key-value pair
type Pair struct {
	Key   string
	Value int
}