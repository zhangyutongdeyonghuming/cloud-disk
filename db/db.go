package db

import (
	"cloud-disk/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDb() {
	db, err := gorm.Open(sqlite.Open("cloud-disk.db"), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}

	// 迁移 schema
	db.AutoMigrate(&model.File{})

	Db = db

}
