package main

import (
	"context"
	"fmt"
	"log"

	"github.com/luquxSentinel/kcrypt/crypt"
	"github.com/luquxSentinel/kcrypt/service"
	"github.com/luquxSentinel/kcrypt/storage"
	"github.com/luquxSentinel/kcrypt/types"
)

func main() {
	storage := storage.NewNopStorage()
	authService := service.NewAuthService(storage, &crypt.BcryptHasher{})

	passwordService := service.NewPasswordService(&crypt.NopEncrypter{}, &crypt.BcryptHasher{}, storage)

	authService.CreateUser(context.Background(), &types.CreateUserData{
		Email:    "luqus@gmail.com",
		Password: "12341234",
	})

	user, err := authService.Login(context.Background(), "luqus@gmail.com", "12341234")
	if err != nil {
		log.Fatal(err)
	}

	err = passwordService.Save(context.Background(), user.UID, "Netflix.com", user.Password, "netflix1234")
	if err != nil {
		log.Fatal(err)
	}

	passwords, err := passwordService.GetPasswords(context.Background(), user.UID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", passwords)

	if len(passwords) > 0 {
		plain, err := passwordService.ShowPassword(context.Background(), "luqus@gmail.com", "12341234", passwords[0].Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Password\nEncrypted: %s\nPlain: %s\n", passwords[0].Password, plain)
	}
}
