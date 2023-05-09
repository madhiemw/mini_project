package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/madhiemw/mini_project/controller"
	"gorm.io/gorm"

)


func AdminAccount(e *echo.Echo, db *gorm.DB) {
	ac := controllers.NewAdminController(db)
	e.POST("/admin/register", ac.RegisterAdmin)
	e.DELETE("/admin/delete-acc/:id", ac.DeleteAdmin)

}

func AdminManagement(e *echo.Echo, db *gorm.DB){
	at := controllers.NewAdminManagement(db)
	e.GET("/admin/all-user-info", at.GetAllUser)
	e.DELETE("/admin/delete-usr/:id", at.DeleteUserByID)
	e.POST("/admin/add-field/:id", at.AddField)
	e.PUT("/admin/edit-field/:id",at.UpdateField)
}

