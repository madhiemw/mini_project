package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/controller"
	"gorm.io/gorm"

)


func UserRoutes(e *echo.Echo, db *gorm.DB) {
	uc := controllers.NewUserController(db)
	e.POST("/users/register", uc.RegisterUser)
	e.PUT("/users/edit-pass/:id", uc.ChangePassword)
	e.DELETE("/users/delete-acc/:id", uc.DeleteUser)
}

