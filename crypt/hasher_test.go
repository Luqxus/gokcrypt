package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasherHash(t *testing.T) {
	password := "12345678"

	var hasher Hasher = &BcryptHasher{}

	_, err := hasher.Hash([]byte(password))

	assert.Nil(t, err)

}

func TestHasherVerify(t *testing.T) {
	password := "12345678"

	var hasher Hasher = &BcryptHasher{}

	h, err := hasher.Hash([]byte(password))

	assert.Nil(t, err)

	passwordInvalid := "123456546"

	err = hasher.Verify(h, []byte(passwordInvalid))

	assert.NotNil(t, err)

	err = hasher.Verify(h, []byte(password))

	assert.Nil(t, err)
}
