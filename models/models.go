package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	uid       string `json:"uid"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"email"`
}
