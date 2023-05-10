package routes

import (
	"github.com/labstack/echo/v4"
	controllers "github.com/madhiemw/mini_project/controller"
	"gorm.io/gorm"
)

func AdminAccount(e *echo.Echo, db *gorm.DB) {
	ac := controllers.NewAdminController(db)
	e.POST("/admin/register", 			ac.RegisterAdmin)
	e.DELETE("/admin/delete-acc/:id", 	ac.DeleteAdmin)

}

func AdminManagement(e *echo.Echo, db *gorm.DB) {
	at := controllers.NewAdminManagement(db)
	e.GET("/admin/all-user-info",		at.GetAllUser)
	e.GET("/admin/all-booking-info", 	at.GetAllBookings)
	e.GET("/admin/all-field", 			at.GetAllField)
	e.POST("/admin/add-field/:id", 		at.AddField)
	e.POST("/admin/add-schedule/:id", 	at.AddSchedules)
	e.PUT("/admin/edit-field/:id", 		at.UpdateField)
	e.DELETE("/admin/delete-usr/:id", 	at.DeleteUserByID)
	// e.DELETE("/admin/delete-field/:id". at.DeleteField)
}
