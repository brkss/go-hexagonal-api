package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	paseto "github.com/brkss/dextrace-server/internal/adapter/auth"
	"github.com/brkss/dextrace-server/internal/adapter/config"
	http "github.com/brkss/dextrace-server/internal/adapter/handler"
	"github.com/brkss/dextrace-server/internal/adapter/logger"
	"github.com/brkss/dextrace-server/internal/adapter/storage/postgres"
	"github.com/brkss/dextrace-server/internal/adapter/storage/postgres/repository"
	"github.com/brkss/dextrace-server/internal/adapter/storage/redis"
	"github.com/brkss/dextrace-server/internal/core/service"
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

	// Init cache 
	cache, err := redis.New(ctx, config.Redis)
	if err != nil {
		slog.Error("Error Initializing cache connection", "error", err)
		os.Exit(0)
	}
	defer cache.Close()


	// Init Token Service 
	token, err := paseto.New(config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}
	

	// dependency injection 
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cache)
	userHandler := http.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, token)
	authHandler := http.NewAuthHandler(authService)

	// Init Router 
	router, err := http.NewRouter(
		config.HTTP,
		token,
		*userHandler,
		*authHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start Server ! 
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Run(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}