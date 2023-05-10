package models

import (
    "gorm.io/gorm"
)


type SchedulesTimeStart struct{
	gorm.Model
	SchedulesDay        string `json:"schedules_day"`
	SchedulesTimeStart  int `json:"schedules_time_start"`
	SchedulesTimeEnd  int `json:"schedules_time_start"`
	
}