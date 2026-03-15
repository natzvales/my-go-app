package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret     string
	JWTExpiration int // in hours
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string

	MEMORY      int
	ITERATIONS  int
	PARALLELISM int
	KEYLENGTH   int
	SALTLENGTH  int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment")
	}
	stringToInt := func(key string) int {
		value, err := strconv.Atoi(os.Getenv(key))
		if err != nil {
			log.Fatalf("Error converting %s to int: %v", key, err)
		}
		return value
	}

	return &Config{
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTExpiration: stringToInt("JWT_EXPIRATION_HOURS"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),

		MEMORY:      stringToInt("MEMORY") * 1024,
		ITERATIONS:  stringToInt("ITERATIONS"),
		PARALLELISM: stringToInt("PARALLELISM"),
		KEYLENGTH:   stringToInt("KEYLENGTH"),
		SALTLENGTH:  stringToInt("SALTLENGTH"),
	}
}
