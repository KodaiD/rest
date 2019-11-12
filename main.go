package main

import (
	"github.com/KodaiD/rest/db"
	"github.com/KodaiD/rest/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}