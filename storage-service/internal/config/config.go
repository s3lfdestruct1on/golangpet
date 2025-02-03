package config

import (
	"os"
	"log/slog"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbHost  string `yaml:"db_host"`
	DbUser  string `yaml:"db_user" env-default:"postgres"`
	DbPass  string `yaml:"db_pass"`
	DbName  string `yaml:"db_name"`
	DbPort  string `yaml:"db_port"`
	Sslmode string `yaml:"sslmode"`
}

func GetConfig() Config {
	var cfg Config
	path := os.Getenv("CONFIG_PATH")
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		slog.Info("CFG path", slog.String("path", path))
		panic(err)
	}
	
	return cfg
}