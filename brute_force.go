package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var characters = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

var found bool

func bruteForce(maxLength int, prefix string, targetHash string) {
	if found {
		return
	}

	hash := sha256.Sum256([]byte(prefix))
	hashStr := hex.EncodeToString(hash[:])

	if len(prefix) > 0 && hashStr == targetHash {
		fmt.Println("Match found:", prefix)
		found = true
		return
	}

	if len(prefix) == maxLength {
		return
	}

	for _, char := range characters {
		bruteForce(maxLength, prefix+char, targetHash)
	}
}

func main() {
	var maxLength int
	var targetHash string

	fmt.Println("Enter the SHA256 hash to brute force: ")
	fmt.Scan(&targetHash)

	fmt.Println("Enter the max length of the decrypted string: ")
	fmt.Scan(&maxLength)

	found = false

	for length := 1; length <= maxLength && !found; length++ {
		bruteForce(length, "", targetHash)
	}

	if !found {
		fmt.Println("No match found")
	}
}
