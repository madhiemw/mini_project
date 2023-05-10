package models

import (
    "gorm.io/gorm"
)


type Schedules struct{
	gorm.Model
	AdminID   	int   `json:"admin_id"`
	SchedulesDay        string `json:"schedules_day"`
	SchedulesTimeStart  int `json:"schedules_time_start"`
	SchedulesTimeEnd  int `json:"schedules_time_start"`
}