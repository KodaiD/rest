// データベースの設定

package db

import (
	"github.com/KodaiD/rest/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "port=5432 user=gopher dbname=gopher password=gopher sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Post{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}