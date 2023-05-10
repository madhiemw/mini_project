package models

import (
    "gorm.io/gorm"
)

type ConfirmedBooking struct{
	gorm.Model
	BookingID		int `json:"booking_id"`
	Confirmed		bool `json:"confirmed"`
}