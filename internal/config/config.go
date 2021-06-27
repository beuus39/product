package config

import (
	"os"
	"strconv"
)

const (
	MONGO_PREFIX         string = "mongodb://"
	MONGO_COLLECTION  string = "MONGODB_COLLECTION"
	MONGO_TIMEOUT     string = "MONGODB_TIMEOUT"
	MONGO_HOST        string = "MONGODB_HOST"
	MONGO_PORT        string = "MONGODB_PORT"
	SERVER_PORT       string = "PRODUCT_SERVER_PORT"
	GRPC_PORT         string = "PRODUCT_GRPC_PORT"
)
type Config struct {
	URL        string
	DB         string
	Timeout    int
	ServerPort string
	GrpcPort   uint
}

func LoadEnv() *Config {
	mongoHost          := os.Getenv(MONGO_HOST)
	mongoPort          := os.Getenv(MONGO_PORT)
	mongoCollection    := os.Getenv(MONGO_COLLECTION)
	mongoTimeout, _    := strconv.Atoi(os.Getenv(MONGO_TIMEOUT))
	serverPort         := os.Getenv(SERVER_PORT)
	grpcPort, _        := strconv.Atoi(os.Getenv(GRPC_PORT))

	if mongoPort == "" {
		mongoPort = "27017"
	}

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoUri           := MONGO_PREFIX + mongoHost + ":" + mongoPort

	return &Config{
		URL: mongoUri,
		DB: mongoCollection,
		Timeout: mongoTimeout,
		ServerPort: serverPort,
		GrpcPort: uint(grpcPort),
	}
}