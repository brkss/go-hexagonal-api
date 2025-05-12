package http

import (
	"log/slog"
	"strings"

	"github.com/brkss/dextrace-server/internal/adapter/config"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	sloggin "github.com/samber/slog-gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	token port.TokenService,
	userHandler UserHandler,
) (*Router, error) { 
	if config.Env == "poroduction" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS 
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigin;
	originList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// v, ok := binding.Validator.Engine().(*validator.Validate)
	// if ok {
	// 	if err := v.RegisterValidation("user")
	// }

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", userHandler.Register)
		}
	}

	return &Router{
		router,
	}, nil	
}
