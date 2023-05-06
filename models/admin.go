package models

import (
    "gorm.io/gorm"
)


type admin struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
}