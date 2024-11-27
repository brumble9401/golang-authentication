package main

import (
	"github.com/brumble9401/golang-authentication/cmd/api"
	"github.com/brumble9401/golang-authentication/configs"
	"github.com/brumble9401/golang-authentication/db"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	pool, err := db.NewPostgresStorage(configs.Envs.DatabaseUrl)
	if err != nil {
		log.Error("Error initializing PostgreSQL storage:", err)
	}

	defer pool.Close()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Error(err)
	}
}