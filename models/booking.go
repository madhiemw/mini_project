package models

import (
    "gorm.io/gorm"
	"time"
)

type Booking struct{
	gorm.Model
	UserID			int 	`json:"user_id"`
	FieldID			int 	`json:"field_id"`
	BookingDate		time.Time`json:"booking_date"`
	BookingHours	int	 	`gorm:"column:BookingHours"`
	StartTime		int 	`gorm:"column:StartTime"`
	EndTime			int 	`gorm:"column:EndTime"`
	BookingID 		int 	`json:"booking_id"`
}