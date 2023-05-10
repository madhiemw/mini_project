package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/models"
	"gorm.io/gorm"
)

type AdminManagement struct {
	db *gorm.DB
}

func NewAdminManagement(db *gorm.DB) *AdminManagement {
	return &AdminManagement{db: db}
}

func (at *AdminManagement) AddField(c echo.Context) error {
	id := c.Param("id")

	Fields := new(models.Fields)
	if err := c.Bind(Fields); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Admin ID is required"})
	}

	adminID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid admin ID"})
	}

	Fields.AdminID = int(adminID)

	if err := at.db.Create(Fields).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create field"})
	}

	return c.JSON(http.StatusOK, Fields)
}

func (at *AdminManagement) DeleteUserByID(c echo.Context) error {
	type requestBody struct {
		UserID string `json:"user_id"`
	}

	rb := new(requestBody)

	if err := c.Bind(rb); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	userID := rb.UserID

	var user models.User
	if err := at.db.Where("user_id = ?", userID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user account"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User account successfully deleted"})
}
func (at *AdminManagement) GetAllUser(c echo.Context) error {
	var users []struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber int    `json:"phone_number"`
	}
	if err := at.db.Table("users").Select("username, email, phone_number").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get bookings"})
	}
	return c.JSON(http.StatusOK, users)
}

func (at *AdminManagement) UpdateField(c echo.Context) error {
	id := c.Param("id")
	var field models.Fields
	if err := at.db.First(&field, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	var Fields struct {
		Field_name string `json:"field_name"`
		Type       string `json:"type"`
	}
	if err := c.Bind(&Fields); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
	}
	if err := at.db.Model(&field).Updates(map[string]interface{}{
		"field_name": Fields.Field_name,
		"type":       Fields.Type,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update field"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "update field"})
}

func (at *AdminManagement) GetAllBookings(c echo.Context) error {
	var bookings []struct {
		UserID       int       `json:"user_id"`
		FieldID      int       `json:"field_id"`
		BookingDate  time.Time `json:"booking_date"`
		BookingHours int       `gorm:"column:BookingHours"`
		StartTime    int       `gorm:"column:StartTime"`
		EndTime      int       `gorm:"column:EndTime"`
	}

	if err := at.db.Table("bookings").Select("user_id, field_id, booking_date, BookingHours, StartTime, EndTime").Find(&bookings).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get bookings"})
	}

	return c.JSON(http.StatusOK, bookings)
}

func (at *AdminManagement) GetAllField(c echo.Context) error {
	var field []struct {
		Type       string
		Field_name string
	}
	if err := at.db.Find(&models.Fields{}).Select("type, field_name").Scan(&field).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get users"})
	}
	return c.JSON(http.StatusOK, field)
}

func (at *AdminManagement) AddSchedules(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid admin ID"})
    }
    schedules := new(models.Schedules)

	schedules.AdminID = id

    if err := at.db.Create(schedules).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create schedules"})
    }

    if err := c.Bind(schedules); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
    }
   
    return c.JSON(http.StatusOK, schedules)
}
