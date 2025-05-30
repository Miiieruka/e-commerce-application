package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"local"`
	DB         DBConfig   `yaml:"db"`
	Redis      string     `yaml:"redis"`
	HttpServer HTTPServer `yaml:"http_server"`
	JwtSecret  string     `yaml:"jwt_secret"`
}

type DBConfig struct {
	Driver string `yaml:"driver" env-default:"postgre"`
	DSN    string `yaml:"dsn" env-required:"true"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
	Port    string `yaml:"port"`
}

func NewConfig() *Config {

	const op = "config.configgo.newconfig"

	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("%s: config file does not exist: %s", op, configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("%s: cannot read config file: %s", op, configPath)
	}

	return &cfg

}
