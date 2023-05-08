package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/models"
	"gorm.io/gorm"
)

type AdminController struct {
	db *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{db: db}
}

func (ac *AdminController) RegisterAdmin(c echo.Context) error {
	var admin models.Admin
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
	}

	var existingAdmin models.Admin
	if err := ac.db.Select("username").Where("username = ?", admin.Username).First(&existingAdmin).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "username sudah terdaftar"})
	}

	if err := ac.db.Create(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal untuk Menambahkan Admin ke dalam database"})
	}

	return c.JSON(http.StatusCreated, map[string]int64{"admin_id": int64(admin.ID)})
}

func (ac *AdminController) DeleteAdmin(c echo.Context) error {
	id := c.Param("id")

	if err := ac.db.Delete(&models.Admin{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dihapus"})
}

func (ac *AdminController) DeleteUserByID(c echo.Context) error {
	type requestBody struct {
		UserID string `json:"user_id"`
	}

	rb := new(requestBody)

	if err := c.Bind(rb); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	userID := rb.UserID

	var user models.User
	if err := ac.db.Where("user_id = ?", userID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user account"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User account successfully deleted"})
}


func (ac *AdminController) AddField(c echo.Context) error {
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

    if err := ac.db.Create(Fields).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create field"})
    }

    return c.JSON(http.StatusOK, Fields)
}