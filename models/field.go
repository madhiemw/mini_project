package models

import (
    "gorm.io/gorm"
)

type Lapangan struct {
	gorm.Model
	Field_name	string `json:"field_name"`
	Type	string `json:"type"`
	AdminID   uint   `json:"admin_id"`
}