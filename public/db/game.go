package db

import "gorm.io/gorm"

var gameDB *gorm.DB

func InitGameDB() {
}

func GetGameDB() *gorm.DB {
	return gameDB
}
