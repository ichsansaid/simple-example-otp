package app

import (
	"github.com/ichsansaid/multi-artha-otp/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {

	//Asumsikan saya pake memory DB aja tidak masalah.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entity.OtpEntity{}); err != nil {
		panic(err)
	}
	return db
}
