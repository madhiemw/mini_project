package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/controller"
	"gorm.io/gorm"

)


func AdminRoutes(e *echo.Echo, db *gorm.DB) {
	ac := controllers.NewAdminController(db)
	e.POST("/admin/register", ac.RegisterAdmin)
	e.DELETE("/admin/delete-acc/:id", ac.DeleteAdmin)
	e.DELETE("/admin/delete-usr/:id", ac.DeleteUserByID)
	e.POST("/admin/add-field/:id, ", ac.AddField)
}