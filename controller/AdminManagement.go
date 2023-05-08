package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/models"
	"gorm.io/gorm"
)


type AdminManagement struct {
	db *gorm.DB
}



func NewAdminManagement (db *gorm.DB) *AdminManagement {
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
	var users []models.User
    if err := at.db.Find(&users).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get users"})
    }
    return c.JSON(http.StatusOK, users)
}

