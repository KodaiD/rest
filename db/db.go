// データベースの設定

package db

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gopher dbname=gopher password=gopher sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}