package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"external.v1/model"
)
  
var db *gorm.DB
var err error

// Init - mysql init
func Init() {
	dsn := "root:@tcp(127.0.0.1:3306)/hasbit_external?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
		}), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.Log{})
}

func DbManager() *gorm.DB {
	return db
}