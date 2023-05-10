package controllers

import (
	"net/http"
	"strconv"
	"errors"
    // "time"
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


// func (ub *UserBooking) ShowBookingConfirmationByUserID(c echo.Context) error {
// 	userID := c.Param("user_id") // assuming the user ID is passed as a URL parameter

// 	var bookings []struct {
//         FieldName  string `json:"field_name"`
// 		Confirmed  bool   `json:"confirmed"`
// 	}

// 	result := ub.db.Table("bookings").
// 		Select("bookings.*, fields.field_name, confirmed_booking.confirmed").
// 		Joins("JOIN fields ON bookings.field_id = fields.field_id").
// 		Joins("LEFT JOIN confirmed_booking ON bookings.bookings_id = confirmed_booking.bookings_id").
// 		Where("bookings.user_id = ?", userID).
// 		Scan(&bookings)

// 	if result.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get bookings"})
// 	}

// 	return c.JSON(http.StatusOK, bookings)
// }

func (ub *UserBooking) ShowBookingConfirmationByUserID(c echo.Context) error {
	userID := c.Param("id") // assuming the user ID is passed as a URL parameter

	var bookings []struct {
		BookingsID int       `json:"bookings_id"`
		Confirmed  string    `json:"confirmed"`
		FieldName  string    `json:"field_name"`
		StartTime  int       `gorm:"column:StartTime"`
		EndTime    int       `gorm:"column:EndTime"`
	}

    result := ub.db.Table("bookings").
    Select("bookings.bookings_id, fields.field_name, confirmed_booking.confirmed, bookings.StartTime, bookings.EndTime").
    Joins("JOIN fields ON bookings.field_id = fields.field_id").
    Joins("LEFT JOIN confirmed_booking ON bookings.bookings_id = confirmed_booking.bookings_id").
    Where("bookings.user_id = ?", userID).
    Scan(&bookings)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get bookings"})
	}

	return c.JSON(http.StatusOK, bookings)
}
