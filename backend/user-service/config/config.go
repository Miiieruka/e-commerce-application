package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env       string     `yaml:"env" env-default:"local"`
	DB        DBCOnfig   `yaml:"db"`
	server    HTTPServer `yaml:"http_server"`
	JwtSecret string     `yaml:"jwt_secret"`
}

type DBCOnfig struct {
	Driver string `yaml:"driver" env-default:"postgre"`
	DSN    string `yaml:"dsn" env-required:"true"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	Idle_timeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
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
