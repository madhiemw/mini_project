package models

import (
    "gorm.io/gorm"
)

type booking struct{
	gorm.Model
	userID			int `json:"user_id"`
	FieldID			int `json:"field_id"`
	BookingDate		string `json:"booking_date"`
	BookingHours	int `json:"booking_hours"`
	StartTime		string `json:"username"`
	EndTime			int `json:""`
	Confirmed		bool `json:"confirmed"`
}