package models

import "gorm.io/gorm"

type AboutOrganization struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
