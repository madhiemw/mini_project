package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/models"
	"gorm.io/gorm"
)

type UserBooking struct {
	db *gorm.DB
}

func UserBookingController(db *gorm.DB) *UserBooking {
	return &UserBooking{db: db}
}

func (ub *UserBooking) GetAllField(c echo.Context) error {
	var field []struct {
        Type string
        Field_name string
    }
	if err := ub.db.Find(&models.Fields{}).Select("type, field_name").Scan(&field).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, field)
}
