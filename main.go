package main

import (
	"cloud-disk/db"
	"cloud-disk/log"
	"cloud-disk/router"
)

func main() {
	log.InitLogger()
	db.InitDb()
	router.InitRouter()
}
