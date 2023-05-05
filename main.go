package main

import (
	"log"
	"net/http"
    "main/middleware/auth"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func main() {
    dsn := "root:44mkcrZP7F3sK2t81XSv@tcp(containers-us-west-124.railway.app:6014)/railway?parseTime=true"
	// dsn := "root:44mkcrZP7F3sK2t81XSv@tcp(localhost:3306)/futsal?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&User{})

	e := echo.New()

    e.Use(auth.BasicAuth())

	e.POST("/user/register", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
		}
		var existingUser User
		if err := db.Select("email").Where("email = ?", user.Email).First(&existingUser).Error; err == nil {

			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email sudah terdaftar"})
		} else if err := db.Select("phone_number").Where("phone_number = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {

			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Nomor Telepon sudah terdaftar"})
		}
		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert user into database"})
		}

		return c.JSON(http.StatusCreated, map[string]int64{"user_id": int64(user.ID)})
	})

	e.DELETE("/user/delete/:id", func(c echo.Context) error {
		id := c.Param("id")

		if err := db.Delete(&User{}, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus user"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dihapus"})
	})
	e.PUT("/user/change-password/:id", func(c echo.Context) error {
		id := c.Param("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		var password struct {
			Password string `json:"password"`
		}
		if err := c.Bind(&password); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind request body"})
		}

		if err := db.Model(&user).Update("password", password.Password).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update phone number"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "password updated successfully"})
	})
	e.Logger.Fatal(e.Start(":8080"))

}
