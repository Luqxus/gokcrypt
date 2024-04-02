package types

type Password struct {
	PasswordID string `json:"password_id"`
	Resource   string `json:"resource"`
	Password   string `json:"password"`
}
