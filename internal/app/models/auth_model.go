package models

type AuthUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"passowrd"`
}
