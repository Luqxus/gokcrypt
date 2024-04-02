package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	var encryptor Encrytor = &NopEncrypter{}

	data := []byte("Tomorrow will be a good day.")

	password, err := getPassword()
	assert.Nil(t, err)

	hiddedData, err := encryptor.Encrypt(data, password)
	assert.Nil(t, err)

	descryptedData, err := encryptor.Decrypt(hiddedData, password)
	assert.Nil(t, err)

	assert.Equal(t, data, descryptedData)
}

func TestEncryptionInvalidKey(t *testing.T) {
	var encryptor Encrytor = &NopEncrypter{}

	data := []byte("Tomorrow will be a good day.")

	password := "12345678"

	var hasher Hasher = &BcryptHasher{}

	hashed, err := hasher.Hash([]byte(password))

	assert.Nil(t, err)

	hiddedData, err := encryptor.Encrypt(data, hashed)
	assert.Nil(t, err)

	password = "12345679"
	hashed, err = hasher.Hash([]byte(password))

	assert.Nil(t, err)

	_, err = encryptor.Decrypt(hiddedData, hashed)
	assert.NotNil(t, err)

}

func getPassword() ([]byte, error) {
	password := "12345678"

	var hasher Hasher = &BcryptHasher{}

	return hasher.Hash([]byte(password))
}
