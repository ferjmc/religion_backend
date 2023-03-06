package main

import (
	"context"
	"log"
	"os"
	"religion/config"
	"religion/internal/server"
	"religion/pkg/database"
	"religion/pkg/firebase"
	"religion/pkg/logger"
)

func main() {
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("loading config: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("Starting religion server mode: %s", cfg.HttpServer.Mode)

	dbConn, err := database.New(cfg)
	if err != nil {
		appLogger.Fatal("cannot connect with the database:", err)
	}

	authClient, err := firebase.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("connecting with authenticator service: %v", err)
	}

	s := server.NewServer(appLogger, cfg, dbConn, authClient)
	appLogger.Fatal(s.Run())
}
