package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
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

	configuration = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: JwtSecretKey,
	}
}

func GetConfig() Config {
	loadConfig()
	return configuration
}
