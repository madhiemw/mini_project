package main

import (
	"os"
	"github.com/madhiemw/mini_project/database"
	"github.com/madhiemw/mini_project/middleware"
	"github.com/madhiemw/mini_project/models"
	"github.com/labstack/echo/v4"
)


func main() {
	db := database.ConnectDB()

	e := echo.New()
    e.Use(auth.BasicAuth())

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})

	AdminRoutes(e, db)
	UserRoutes(e, db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
