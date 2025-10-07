package mysql

import (
	"fmt"
	"log"
	"studyonline/dao/entity"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg, err := ini.Load("./init/project.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	host := cfg.Section("mysql").Key("host").String()
	port := cfg.Section("mysql").Key("port").String()
	database := cfg.Section("mysql").Key("database").String()

	// dsn := "user:123456@tcp(127.0.0.1:3306)/studyonline?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

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
	db.AutoMigrate(&entity.Score{})
	db.AutoMigrate(&entity.Unit{})
	db.AutoMigrate(&entity.Homework{})
	db.AutoMigrate(&entity.Discuss{})
	db.AutoMigrate(&entity.Permission{})
	db.AutoMigrate(&entity.Request{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.DownloadLog{})
	DB = db
}
