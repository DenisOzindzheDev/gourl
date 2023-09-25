package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-default:"./storage" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Adress      string        `yaml:"adress" env-default:"localhost:8080" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s" env-required:"true"`
}

// readConfig
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH") //get config path
	if configPath == "" {
		configPath = "./config/local.yaml" // default config path
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath) //if file not found
	}
	var cfg Config //initialize config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil { //read config file if error throw error
		log.Fatalf("cannot read config file: %s", err)
	}

	return &cfg
}
