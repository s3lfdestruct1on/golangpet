package main

import (
	"storage-service/internal/config"
	"storage-service/internal/storage"
)

func main() {

	cfg := config.GetConfig()

	storage.CreateUser(cfg.DbUser,cfg.DbPass,cfg.DbHost,cfg.DbHost,cfg.DbName)
	
}