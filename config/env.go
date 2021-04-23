package config

import "os"

var (
	ServerPort    = GetEnv("SERVER_PORT", "9000")
	MongoUrl      = GetEnv("MONGODB_URL", "mongodb://root:root@localhost:27017")
	MongoDatabase = GetEnv("MONGODB_DATABASE", "demo")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
