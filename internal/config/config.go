package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env          string       `yaml:"environment" env-required:"true"`
	HTTPServer   HTTPServer   `yaml:"http_server"`
	MetricServer MetricServer `yaml:"metric_server"`
	Database     Database     `yaml:"database"`
}

type HTTPServer struct {
	Port        uint16        `yaml:"port"`
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type MetricServer struct {
	Port    uint16 `yaml:"port"`
	Address string `yaml:"address"`
}

type Database struct {
	Engine       string `yaml:"engine"`
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Name         string `yaml:"name"`
	MigrationDir string `yaml:"migration_path"`
}

func MustLoad() *AppConfig {
	godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists")
	}

	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config")
	}

	return &cfg
}
