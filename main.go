package main

import (
	"log"
	"net/http"
	"github.com/madhiemw/mini_project/controller"
	"github.com/madhiemw/mini_project/database"
	"github.com/madhiemw/mini_project/middleware"
	"github.com/madhiemw/mini_project/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)


func main() {
	db := initDB()
    defer db.Close()

	e := echo.New()

    e.Use(auth.BasicAuth())

	db.AutoMigrate(&models.User{})

	
	e.Logger.Fatal(e.Start(":8080"))

}
