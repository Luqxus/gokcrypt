package service

import (
	"context"
	"errors"

	"github.com/luquxSentinel/kcrypt/crypt"
	"github.com/luquxSentinel/kcrypt/storage"
	"github.com/luquxSentinel/kcrypt/types"
)

type PasswordService interface {
	Save(ctx context.Context, ui, resource, password string) error
	GetPasswords() ([]*types.Password, error)
	ShowPassword(ctx context.Context, email, loginPassword, password string) (string, error)
}

type passwordService struct {
	encrytor crypt.Encrytor
	hasher   crypt.Hasher
	storage  storage.Storage
}

func NewPasswordService(encrytor crypt.Encrytor, hasher crypt.Hasher, storage storage.Storage) *passwordService {
	return &passwordService{
		encrytor: encrytor,
		hasher:   hasher,
		storage:  storage,
	}
}

func (s *passwordService) Save(ctx context.Context, uid, resource, hashedPassword, password string) error {
	// TODO: create new password type
	newpassword := new(types.Password)

	newpassword.Resource = resource

	//TODO: encrypt password
	pwd, err := s.encrytor.Encrypt([]byte(password), []byte(hashedPassword))
	if err != nil {
		return nil
	}

	newpassword.Password = string(pwd)

	// TODO: persist password type
	err = s.storage.CreatePassword(ctx, uid, newpassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *passwordService) GetPasswords(ctx context.Context, uid string) ([]*types.Password, error) {
	return s.storage.GetPasswords(ctx, uid)
}

func (s *passwordService) ShowPassword(ctx context.Context, email, loginPassword, password string) (string, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return password, err
	}
	err = s.hasher.Verify([]byte(user.Password), []byte(loginPassword))
	if err != nil {
		return password, errors.New("wrong password")
	}

	b, err := s.encrytor.Decrypt([]byte(password), []byte(user.Password))
	if err != nil {
		return password, err
	}

	return string(b), nil
}
