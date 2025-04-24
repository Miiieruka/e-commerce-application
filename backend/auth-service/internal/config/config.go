package confid

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string   `yaml:"env" env-default:"local"`
	DB          DBCOnfig `yaml:"db"`
	HTTPServer  `yaml:"http_server"`
	GoogleOAuth GoogleOAuth `yaml:"google_oauth"`
	JwtSecret   string      `yaml:"jwt_secret"`
}

type DBCOnfig struct {
	Driver string `yaml:"driver" env-default:"postgre"`
	DSN    string `yaml:"dsn" env-required:"true"`
}

type GoogleOAuth struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	Idle_timeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
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
