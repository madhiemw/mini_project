package controllers

import (
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "github.com/labstack/echo/v4"
    "github.com/madhiemw/mini_project/models"
    "gorm.io/gorm"
    "encoding/json"
)


type UserAccount struct{
    db *gorm.DB
}

func UserAccountRoute(db *gorm.DB) *UserAccount {
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

    passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
    }
    user.Password = string(passwordHash)

    if err := uc.db.Create(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert user into database"})
    }

    return c.JSON(http.StatusCreated, map[string]int64{"user_id": int64(user.ID)})
}


func (uc *UserAccount) DeleteUser(c echo.Context) error {
    id := c.Param("id")

    if err := uc.db.Where("user_id = ?", id).Delete(&models.User{}).Error; err != nil {
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

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
    }

    if err := uc.db.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update password"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
}

func (uc *UserAccount) LoginUser(c echo.Context) error {
    var login models.User

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully Login", token)
}
