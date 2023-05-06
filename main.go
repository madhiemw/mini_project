package main

import (
	"os"
	"github.com/madhiemw/mini_project/controller"
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

	uc := controllers.NewUserController(db)
	e.POST("/users/register", uc.RegisterUser)
	e.PUT("/users/edit-pass/:id", uc.ChangePassword)
	e.DELETE("/users/delete-acc/:id", uc.DeleteUser)

	ac := controllers.NewAdminController(db)
	e.POST("/admin/register", ac.RegisterAdmin)
	e.DELETE("/admin/delete-acc/:id", ac.DeleteAdmin)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))}
