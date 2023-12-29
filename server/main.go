package main

import (
	"auth/server/auth"
	"auth/server/configuration"
	"auth/server/connection"
	"auth/server/router"
	"fmt"
)

func main() {
	// Load the configuration from the file
	cfg, err := configuration.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Extract the database connection parameters
	dbConfig := connection.DBConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Dbname:   cfg.Database.Dbname,
	}

	// Set the database connection parameters
	connection.Connection(dbConfig)

	// Set the authorization parameters
	auth.SetAuthorizationConfig(cfg.Authorization.AuthorizedUsername, cfg.Authorization.AuthorizedPassword)

	// Start the server
	router.Router(dbConfig)
}
