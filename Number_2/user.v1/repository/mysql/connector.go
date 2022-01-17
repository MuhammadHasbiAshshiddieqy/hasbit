package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"user.v1/model"
)
  
var db *gorm.DB
var err error

// Init - mysql init
func Init() {
	dsn := "root:@tcp(127.0.0.1:3306)/hasbit_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
		}), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.User{})
}

// DbManager - return db connection
func DbManager() *gorm.DB {
	return db
}