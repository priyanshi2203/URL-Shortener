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
}

// Pair represents a key-value pair
type Pair struct {
	Key   string
	Value int
}