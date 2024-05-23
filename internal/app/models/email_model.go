package models

type FormDataAboutTraning struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	Direction  string `json:"direction"`
	Group      string `json:"group"`
	Quantity   string `json:"quantity"`
	Message    string `json:"message"`
}

type FormArmy struct {
	Fio       string `json:"fio"`
	Direction string `json:"direction"`
	Group     string `json:"group"`
	ArmyName  string `json:"army_name"`
	Message   string `json:"message"`
}

type ScholarshipForm struct {
	FirstName     string `json:"first_name"`
	Patronymic    string `json:"patronymic"`
	LastName      string `json:"last_name"`
	Specialty     string `json:"specialty"`
	Group         string `json:"group"`
	PaymentPeriod string `json:"payment_period"`
	Quantity      string `json:"quantity"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Message       string `json:"message"`
}
