package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/madhiemw/mini_project/models"
    "gorm.io/gorm"
)


type UserAccount struct{
    db *gorm.DB
}

func UserAccountController(db *gorm.DB) *UserAccount {
    return &UserAccount{db: db}
}

func (uc *UserAccount) RegisterUser(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
    }

    var existingUser models.User
    if err := uc.db.Select("email").Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email sudah terdaftar"})
    } else if err := uc.db.Select("phone_number").Where("phone_number = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Nomor Telepon sudah terdaftar"})
    }

    if err := uc.db.Create(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert user into database"})
    }

    return c.JSON(http.StatusCreated, map[string]int64{"user_id": int64(user.ID)})
}

func (uc *UserAccount) DeleteUser(c echo.Context) error {
    id := c.Param("id")

    if err := uc.db.Delete(&models.User{}, id).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus user"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dihapus"})
}

func (uc *UserAccount) ChangePassword(c echo.Context) error {
    id := c.Param("id")
    var user models.User
    if err := uc.db.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
    }

    var password struct {
        Password string `json:"password"`
    }
    if err := c.Bind(&password); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
    }

    if err := uc.db.Model(&user).Update("password", password.Password).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update password"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "password updated successfully"})
}
