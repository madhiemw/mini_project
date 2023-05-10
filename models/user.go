package models

import (
    "gorm.io/gorm"
)


type User struct {
	gorm.Model
    ID          uint   `json:"user_id"`
    Username    string `json:"username"`
    Password    string `json:"password"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
    Address     string `json:"address"`
}