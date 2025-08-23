package mysql

import (
	"studyonline/dao/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/studyonline?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&entity.Admin{})
	db.AutoMigrate(&entity.Student{})
	db.AutoMigrate(&entity.Teacher{})
	db.AutoMigrate(&entity.Resource{})
	db.AutoMigrate(&entity.Dataset{})

	DB = db
}
