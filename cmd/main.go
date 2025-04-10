package main

import (
	"fmt"
	"github.com/Egorpalan/api-pvz/config"
	"github.com/Egorpalan/api-pvz/pkg/db"
	"github.com/Egorpalan/api-pvz/pkg/logger"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	// Init Logger
	logger.Init()

	// Init DB
	database := db.NewPostgresDB(cfg.DB)
	defer database.Close()

	// TODO: Init router, routes, server handlers

	log.Printf("Starting server on port %s...", cfg.AppPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), nil))
}
