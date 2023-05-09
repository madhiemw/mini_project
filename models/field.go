package models

import (
    "gorm.io/gorm"
)

type Fields struct {
	gorm.Model
	Field_name	string `json:"field_name"`
	Type		string `json:"type"`
	AdminID   	int   `json:"admin_id"`
	FieldID		int   `json:"field_id"`
}