package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/luquxSentinel/kcrypt/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)

	CreatePassword(ctx context.Context, uid string, v *types.Password) error
	GetPasswords(ctx context.Context, uid string) ([]*types.Password, error)
}

type NopStorage struct {
	passwordStore map[string][]*types.Password
	userStore     []*types.User
}

func NewNopStorage() *NopStorage {
	return &NopStorage{
		passwordStore: make(map[string][]*types.Password),
		userStore:     make([]*types.User, 0),
	}
}

func (s *NopStorage) CreateUser(ctx context.Context, user *types.User) error {
	s.userStore = append(s.userStore, user)

	return nil
}

func (s *NopStorage) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	for _, user := range s.userStore {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *NopStorage) CreatePassword(ctx context.Context, uid string, v *types.Password) error {
	log.Println(v)

	if _, ok := s.passwordStore[uid]; !ok {
		l := make([]*types.Password, 0)

		l = append(l, v)
		s.passwordStore[uid] = l
		return nil
	}

	s.passwordStore[uid] = append(s.passwordStore[uid], v)

	log.Printf("%+v", s.passwordStore)

	return nil
}

func (s *NopStorage) GetPasswords(ctx context.Context, uid string) ([]*types.Password, error) {
	v, ok := s.passwordStore[uid]
	if ok {
		fmt.Println("----------------")
		log.Printf("%+v", v)
		fmt.Println("----------------")

		return v, nil
	}

	return nil, errors.New("user has not passwords saved")
}
