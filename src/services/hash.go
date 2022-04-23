package services

import (
	"crypto/sha256"
	"fmt"
)

type Hash struct {
	hash string
}

// create a new hash function to hash a string
func NewHash(str string) Hash {
	return Hash{hash: fmt.Sprintf("%x", sha256.Sum256([]byte(str)))}

}
