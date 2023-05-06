package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/madhiemw/mini_project/models"
    "gorm.io/gorm"
)

type AdminController struct{
    db *gorm.DB
}


func NewAdminController(db *gorm.DB) *AdminController {
    return &AdminController{db: db}
}


func (ac *AdminController) RegisterAdmin(c echo.Context) error {
    var user models.Admin
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
    }

    var existingUser models.Admin
    if err := uc.db.Select("username").Where("username = ?", user.Email).First(&existingUser).Error; err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "username sudah terdaftar"})
    } 

    if err := uc.db.Create(&Admin).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert Admin into database"})
    }

    return c.JSON(http.StatusCreated, map[string]int64{"user_id": int64(user.ID)})
}
