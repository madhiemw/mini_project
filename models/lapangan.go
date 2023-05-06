package models

import (
    "gorm.io/gorm"
)



type lapangan struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}