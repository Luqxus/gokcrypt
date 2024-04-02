package types

type User struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type CreateUserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
