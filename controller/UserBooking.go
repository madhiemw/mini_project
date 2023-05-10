package controllers

import (
	"net/http"
	"strconv"
	"errors"

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
func (ub *UserBooking) CreateBooking(c echo.Context) error {
    var booking models.Booking
    if err := c.Bind(&booking); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
    }

    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
    }

    booking.UserID = userID

    if booking.FieldID == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "FieldID cannot be empty"})
     }
    // if booking.BookingDate.IsZero() {
    //     return c.JSON(http.StatusBadRequest, map[string]string{"message": "BookingDate cannot be empty"})
    // }
    // if booking.StartTime < 0 || booking.StartTime > 23 {
    //     return c.JSON(http.StatusBadRequest, map[string]string{"message": "StartTime must be between 0 and 23"})
    // }
    // if booking.EndTime < 0 || booking.EndTime > 23 {
    //     return c.JSON(http.StatusBadRequest, map[string]string{"message": "EndTime must be between 0 and 23"})
    // }
    // if booking.StartTime >= booking.EndTime {
    //     return c.JSON(http.StatusBadRequest, map[string]string{"message": "StartTime must be before EndTime"})
    // }

    booking.BookingHours = booking.EndTime - booking.StartTime

    var overlappingBookings []models.Booking
    if err := ub.db.Where("field_id = ? AND booking_date = ? AND StartTime < ? AND EndTime > ?", booking.FieldID, booking.BookingDate, booking.EndTime, booking.StartTime).Find(&overlappingBookings).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to query existing bookings"})
    }
	if err := ub.db.Where("field_id = ? AND booking_date = ? AND StartTime <= ? AND EndTime >= ?", booking.FieldID, booking.BookingDate, booking.EndTime, booking.StartTime).Find(&overlappingBookings).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to query existing bookings"})
}
    if len(overlappingBookings) > 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "jadwal tidak tersedia"})
    }

    if err := ub.db.Create(&booking).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert booking into database"})
    }

    return c.JSON(http.StatusCreated, map[string]int64{"booking_id": int64(booking.ID)})
}
