package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/brkss/dextrace-server/internal/adapter/config"
	"github.com/brkss/dextrace-server/internal/adapter/logger"
	"github.com/brkss/dextrace-server/internal/adapter/storage/postgres"
)




func main() {

	// load envirenement variables 
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environement variables", "error", err)
		os.Exit(1);
	}

	logger.Set(config.App)
	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database 
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1);
	}

	defer db.Close()

	// Migrate database 
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfuly migrated the database")

	

	fmt.Printf("%v : ", config.App);

	//fmt.Println("Hello, World!")
}