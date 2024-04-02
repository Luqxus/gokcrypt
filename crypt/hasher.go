package crypt

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(v []byte) ([]byte, error)
	Verify(hashed []byte, plain []byte) error
}

type BcryptHasher struct{}

func (h *BcryptHasher) Hash(v []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(v, 14)
}

func (h *BcryptHasher) Verify(hashed []byte, plain []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plain)
}
