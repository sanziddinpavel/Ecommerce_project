package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load: ", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)

	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("service name is required")
		os.Exit(1)

	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("httpPort is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}
	JwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if JwtSecretKey == "" {
		fmt.Println("jwt secret key required")
		os.Exit(1)
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("Port is required")
		os.Exit(1)
	}
	dbprt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("Name is required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("Password is required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("ENABLE_SSL_MODE")

	enablSSLMode, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbprt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMODE: enablSSLMode,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: JwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}

	return configuration
}
