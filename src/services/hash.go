package services

import (
	"crypto/sha256"
	"fmt"
)

type Hash struct {
	hash string
}

// create hash function to hash password
func (h *Hash) Create(password string) {
	h.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
