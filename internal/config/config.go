package config

import "os"

type Config struct {
	MongoURI string
	Port     string
}

func Load() Config {
	return Config{
		MongoURI: getEnv("MONGO_URI", "mongodb+srv://dbuserTest:w61rynzfdRpXjx4h@test.qsj0mrb.mongodb.net/?retryWrites=true&w=majority&appName=Test"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}