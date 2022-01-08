package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Nonsense struct {
	Id      int    `json:"id"`
	Content string `json:"filename" gorm:"type:string"`
	Time    string `json:"time" gorm:"type:string"`
}

type LocalFile struct {
	Name         string
	Link         string
	Size         string
	IsFolder     bool
	ModifiedTime string
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./.go-nonsense.db")
	if err == nil {
		DB = db
		db.AutoMigrate(&Nonsense{})
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}

func All() ([]*Nonsense, error) {
	var nonsenses []*Nonsense
	var err error
	err = DB.Find(&nonsenses).Error
	return nonsenses, err
}

func (nonsense *Nonsense) Insert() error {
	var err error
	err = DB.Create(nonsense).Error
	return err
}

func (nonsense *Nonsense) Delete() error {
	var err error
	err = DB.Delete(nonsense).Error
	return err
}
