package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env       string `yaml:"env" env-required:"true"`
	Server    Server `yaml:"server" env-required:"true"`
	DB        DBConfig
	REDIS     string `env:"REDIS_URL" env-required:"true"`
	JWTSecret string `env:"JWT_SECRET"`
}

type Server struct {
	Addr string `yaml:"address"`
	Port int    `yaml:"port"`
}
type DBConfig struct {
	Driver string `env:"DB_DRIVER" env-required:"true"`
	DSN    string `env:"DB_DSN" env-required:"true"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("%s: Config path does not exist", op)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("%s: Error reading config file: %v", op, err)
	}
	return &cfg
}
