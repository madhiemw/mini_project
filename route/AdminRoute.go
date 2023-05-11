package routes

import (
	"github.com/labstack/echo/v4"
	 "github.com/madhiemw/mini_project/controller"
	"gorm.io/gorm"
)

func AdminAccount(e *echo.Echo, db *gorm.DB) {
    ac := controllers.NewAdminController(db)

    // tambahkan middleware JWT untuk setiap route pada grup admin
    g := e.Group("/admin", middleware.JWT([]byte(getSecretKey())))
    g.POST("/register", ac.RegisterAdmin)
    g.DELETE("/delete-acc/:id", ac.DeleteAdmin)
}

func AdminManagement(e *echo.Echo, db *gorm.DB) {
    at := controllers.NewAdminManagement(db)

    // tambahkan middleware JWT untuk setiap route pada grup admin
    g := e.Group("/admin", middleware.JWT([]byte(getSecretKey())))
    g.GET("/all-user-info", at.GetAllUser)
    g.GET("/all-booking-info", at.GetAllBookings)
    g.GET("/all-field", at.GetAllField)
    g.POST("/add-field/:id", at.AddField)
    g.POST("/add-schedule/:id", at.AddSchedules)
    g.POST("/confirmation/:id", at.ConfirmBooking)
    g.PUT("/edit-field/:id", at.UpdateField)
    g.DELETE("/delete-usr/:id", at.DeleteUserByID)
    // g.DELETE("/delete-field/:id", at.DeleteField)
}

func getSecretKey() []byte {
    // ganti dengan kunci rahasia Anda
    return []byte("your_secret_key")
}