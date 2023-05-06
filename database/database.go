package database 

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
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
	// Set database connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
