package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	App      *App      `json:"app,omitempty"`
	Database *Database `json:"database,omitempty"`
}

type App struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type Database struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
}

func LoadConfig() *AppConfig {
	cfg := viper.New()
	cfg.SetConfigFile("config.json")
	cfg.SetConfigType("json")
	cfg.AddConfigPath("./")

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatalf("cant load config : %v", err)
	}

	configApp := &AppConfig{
		App: &App{
			Name:   cfg.GetString("app.name"),
			Author: cfg.GetString("app.author"),
		},
		Database: &Database{
			User:     cfg.GetString("database.user"),
			Password: cfg.GetString("database.password"),
			Host:     cfg.GetString("database.host"),
			Port:     cfg.GetInt("database.port"),
			Name:     cfg.GetString("database.name"),
		},
	}

	return configApp
}
