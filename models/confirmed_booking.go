package models

import (
    "gorm.io/gorm"
)

type ConfirmedBooking struct{
	gorm.Model
	AdminID			int  `json:"admin_id"`
	BookingID		int `json:"booking_id"`
	Confirmed		bool `json:"confirmed"`
}