// データベースの設定

package db

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user= dbname=gwp password=gwp sslmode=disable")
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}