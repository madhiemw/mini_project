package controllers

import (
    // "net/http"

    // "github.com/labstack/echo/v4"
    // "github.com/madhiemw/mini_project/models"
    "gorm.io/gorm"
)


type UserBooking struct{
    db *gorm.DB
}

func UserBookingController(db *gorm.DB) *UserBooking {
    return &UserBooking{db: db}
}
