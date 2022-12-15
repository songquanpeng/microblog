package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"microblog/common"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	time.Now().Unix()
	db, err := gorm.Open(sqlite.Open(common.SQLitePath), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	err = db.AutoMigrate(&Post{})
	return err
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}
