package initializer

import (
	confid "auth-service/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error in dot env loading: %s", err.Error())
	}
}

func LoadConfig() *confid.Config {
	config := confid.MustLoad()
	return config
}
