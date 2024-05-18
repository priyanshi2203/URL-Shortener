package internal

import (
    "crypto/sha256"
    "encoding/hex"
    "strings"
)

// GenerateUniqueHash generates a unique 5-digits hash for a given URL
func (us *URLShortener)GenerateUniqueHash(url string) string {
    mu.Lock()
    defer mu.Unlock()

    if hash, ok := us.UrlHashes[url]; ok {
        return hash
    }

    // Calculate hash of the URL
    hashBytes := sha256.Sum256([]byte(url))
    hash := hex.EncodeToString(hashBytes[:])

    // Take the first 5 characters as the unique hash
    uniqueHash := hash[:5]

    // Store the hash for future reference
    us.UrlHashes[url] = uniqueHash

    return uniqueHash
}

// getDomain extracts the domain from a URL
func getDomain(url string) string {
    parts := strings.Split(url, "/")
    if len(parts) >= 3 && strings.HasPrefix(parts[2], "www.") {
        return parts[2][4:]
    }
    return parts[2]
}