package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string     `yaml:"env" env-required:"true"`
	GRPCServer GRPCServer `yaml:"grpc_server" env-required:"true"`
	DB         DBConfig
	JWTSecret  string `env:"JWT_SECRET"`
}

type DBConfig struct {
	Driver string `env:"DB_DRIVER" env-required:"true"`
	DSN    string `env:"DB_DSN" env-required:"true"`
}

type GRPCServer struct {
	Addr    string        `yaml:"address"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
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
