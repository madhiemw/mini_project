package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/models"
	"gorm.io/gorm"
	
)

type AdminAccount struct {
	db *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminAccount {
	return &AdminAccount{db: db}
}

func (ac *AdminAccount) RegisterAdmin(c echo.Context) error {
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

func (ac *AdminAccount) DeleteAdmin(c echo.Context) error {
	id := c.Param("id")

	if err := ac.db.Delete(&models.Admin{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dihapus"})
}
