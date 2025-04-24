package confid

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBCOnfig struct {
	Driver string `yaml:"driver" env-default:"postgre"`
	DSN    string `yaml:"dsn" env-required:"true"`
}

type Config struct {
	Env        string   `yaml:"env" env-default:"local"`
	DB         DBCOnfig `yaml:"db"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	address      string        `yaml:"address" env-default:"localhost:8080"`
	timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	idle_timeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func MustLoad() *Config {
	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist")
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("couldn't read config file %s", configPath)
	}

	return &config

}
