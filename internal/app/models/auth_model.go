package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password []byte `json:"passowrd"`
}

type LoginData struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}
