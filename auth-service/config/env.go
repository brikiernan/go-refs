package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost         string
	Port               string
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	JWTExpiryInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		PublicHost:         getEnv("PUBLIC_HOST"),
		Port:               getEnv("PORT"),
		DBUser:             getEnv("DB_USER"),
		DBPassword:         getEnv("DB_PASSWORD"),
		DBName:             getEnv("DB_NAME"),
		JWTSecret:          getEnv("JWT_SECRET"),
		JWTExpiryInSeconds: getEnvAsInt("JWT_EXPIRY_IN_SECONDS"),
	}

}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Fatalf("Failed to load %s env variable", key)
	return ""
}

func getEnvAsInt(key string) int64 {
	if value, ok := os.LookupEnv(key); ok {
		numb, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0
		}

		return numb
	}

	log.Fatalf("Failed to load %s env variable", key)
	return 0
}
