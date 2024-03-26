package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Fullname string `json:"fullname"`
	Group    string `json:"group"`
	Username string `json:"username"`
	Password []byte `json:"passowrd"`
}

type LoginData struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}
