package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppsConfig   AppsConfig     `envconfig:"APP"`
	UserWalletDB DatabaseConfig `envconfig:"USER_WALLET_DB"`
}

type AppsConfig struct {
	Port  string `envconfig:"PORT" default:"8080"`
	Debug bool   `envconfig:"DEBUG" default:"false"`
	Env   string `envconfig:"ENV" default:"dev"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"HOST"`
	Port     int    `envconfig:"PORT"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Name     string `envconfig:"NAME"`
}

func InitConfig() Config {
	_ = godotenv.Load(".env")

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to process environment variables: %v", err)
	}
	if cfg.AppsConfig.Env == "local" {
		cfg.UserWalletDB.Host = "localhost"
	}
	return cfg
}
