package main

import (
	"cloud-disk/config"
	"cloud-disk/db"
	"cloud-disk/log"
	"cloud-disk/router"
)

func main() {
	log.InitLogger()
	config.InitConfig()
	db.InitDb()
	router.InitRouter()
}
