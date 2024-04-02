package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/luquxSentinel/kcrypt/crypt"
	"github.com/luquxSentinel/kcrypt/storage"
	"github.com/luquxSentinel/kcrypt/types"
)

type AuthService interface {
	CreateUser(ctx context.Context, data *types.CreateUserData) error
	Login(ctx context.Context, email string, password string) (*types.User, error)
}

type authService struct {
	storage storage.Storage
	hasher  crypt.Hasher
}

func NewAuthService(storage storage.Storage, hasher crypt.Hasher) *authService {
	return &authService{
		storage: storage,
		hasher:  hasher,
	}
}

func (s *authService) CreateUser(ctx context.Context, data *types.CreateUserData) error {
	// TODO: check if email is not already in use

	// TODO: create new user
	user := new(types.User)

	user.UID = uuid.NewString()
	user.Email = data.Email

	// TODO: hash user password
	b, err := s.hasher.Hash([]byte(data.Password))
	if err != nil {
		return err
	}

	user.Password = string(b)

	// TODO: persist user into database
	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(ctx context.Context, email string, password string) (*types.User, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = s.hasher.Verify([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong email or password")
	}

	return user, nil
}
