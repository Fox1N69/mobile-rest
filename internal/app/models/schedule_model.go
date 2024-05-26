package models

import (
	"time"

	"gorm.io/gorm"
)

type Prepod struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName  string `gorm:"not null" json:"first_name"`
	SecondName string `gorm:"not null" json:"second_name"`
	Surname    string `gorm:"not null" json:"surname"`
}

type Group struct {
	gorm.Model
	NAME string `gorm:"primaryKey;unique" json:"name"`
}

type Subject struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NAME string `gorm:"not null" json:"name"`
}

type Change struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;autoIncrement" json:"id"`
	IsChange   bool
	Date       time.Time `gorm:"not null" json:"date"`
	Urok       string
	GroupID    string  `gorm:"not null;unique"`
	Number     int     `json:"number"`
	Classroom  string  `json:"classroom"`
	PrepodID   uint    `json:"prepod_id"`
	SubjectID  uint    `json:"subject_id"`
	Comment    string  `json:"comment"`
	DeleteUrok *int    `json:"delete_urok"`
	Prepod     Prepod  `gorm:"foreignKey:PrepodID" json:"prepod"`
	Subject    Subject `gorm:"foreignKey:SubjectID" json:"subject"`
}

type Urok struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	IsChange  bool
	SubjectID uint      `gorm:"not null" json:"subject_id"`
	PrepodID  uint      `gorm:"not null" json:"prepod_id"`
	GroupID   string    `gorm:"not null;unique"`
	Date      time.Time `gorm:"not null" json:"date"`
	Number    int       `gorm:"not null" json:"number"`
	Classroom string    `json:"classroom"`
	Prepod    Prepod    `gorm:"foreignKey:PrepodID" json:"prepod"`
	Subject   Subject   `gorm:"foreignKey:SubjectID" json:"subject"`
}
