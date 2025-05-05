package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App 	*App
		DB 		*DB
		Token 	*Token
		Redis	*Redis
		HTTP	*HTTP	
	}
	// App contains all the env variables for the application 
	App struct {
		Name 	string
		Env 	string
	}
	// Tokne contain all the envs variables got the token service 
	Token struct {
		Duration string
	}
	// Redis contains all the env variables for the cache service 
	Redis struct {
		Addr	 string
		Password string
	}
	// Database contains all the env variable for the database 
	DB struct {
		Connection 	string
		Host 		string
		Port 		string
		User 		string
		Password 	string
		Name		string
	}
	// HTTP contain all the env variable for the http server 
	HTTP struct {
		Env 			string
		URL 			string
		Port 			string
		AllowedOrigin	string
	}
)


// New create new Container instance 
func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: 	os.Getenv("APP_NAME"),
		Env: 	os.Getenv("APP_ENV"),
	}

	token := &Token{
		Duration: os.Getenv("TOKEN_DURATION"),
	}

	redis := &Redis{
		Addr: 		os.Getenv("REDIS_ADDR"),
		Password: 	os.Getenv("REDIS_PASSWORD"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Name: 		os.Getenv("DB_NAME"),
		User: 		os.Getenv("DB_USER"),
		Password: 	os.Getenv("DB_PASSWORD"),
		Port: 		os.Getenv("DB_PORT"),
		Host: 		os.Getenv("DB_HOST"),
	}

	http := &HTTP{
		Env: 			os.Getenv("APP_ENV"),
		URL: 			os.Getenv("HTTP_URL"),
		Port: 			os.Getenv("HTTP_PORT"),
		AllowedOrigin: 	os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		app,
		db,
		token,
		redis,
		http,
	}, nil

}