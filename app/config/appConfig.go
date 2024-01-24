package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppConfig *AppConfig
}

// function provider
func LoadConfig() IConfig {
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
			Port:   cfg.GetInt("app.port"),
		},
		Database: &Database{
			User:     cfg.GetString("database.user"),
			Password: cfg.GetString("database.password"),
			Host:     cfg.GetString("database.host"),
			Port:     cfg.GetInt("database.port"),
			Name:     cfg.GetString("database.name"),
		},
		Jaeger: &Jaeger{
			Host: cfg.GetString("jaeger.host"),
			Port: cfg.GetInt("jaeger.port"),
		},
		Jwt: &Jwt{
			SecretKey: cfg.GetString("jwt.secret_key"),
		},
	}

	return &Config{AppConfig: configApp}
}

func (c *Config) GetConfig() *AppConfig {
	return c.AppConfig
}

type AppConfig struct {
	App      *App      `json:"app,omitempty"`
	Database *Database `json:"database,omitempty"`
	Jaeger   *Jaeger   `json:"jaeger,omitempty"`
	Jwt      *Jwt      `json:"jwt,omitempty"`
}

type App struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Port   int    `json:"port,omitempty"`
}

type Database struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Jaeger struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

type Jwt struct {
	SecretKey string `json:"secret_key,omitempty"`
}
