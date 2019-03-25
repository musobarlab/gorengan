package config

import (
	"errors"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

var (

	// PORT config
	Port int

	// DBHost config
	DBHost string
	// DBName config
	DBName string
	// DBUser config
	DBUser string
	// DBPassword config
	DBPassword string
	// DBPort config
	DBPort int
)

// Load function will load all config from environment variable
func Load() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("PORT env is not loaded")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	Port = port

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return errors.New("DB_HOST env is not loaded")
	}

	// set DBHost
	DBHost = dbHost

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return errors.New("DB_NAME env is not loaded")
	}

	// set DBName
	DBName = dbName

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return errors.New("DB_USER env is not loaded")
	}

	// set DBUser
	DBUser = dbUser

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return errors.New("DB_PASSWORD env is not loaded")
	}

	// set DBPassword
	DBPassword = dbPassword

	dbPortStr, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return errors.New("DB_PORT env is not loaded")
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return err
	}

	// set DBPort
	DBPort = dbPort
	// ------------------------------------

	return nil
}
