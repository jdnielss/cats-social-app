package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// TODO
/*
1. Konfigurasi untuk API -> Port
2. Konfigurasi untuk Database -> User, Host, Port, Password, DbName, DbDriver
3. FileConfig -> Penyimpanan file atau upload
4. Token -> untuk konfigurasi authentication
5. dll...
*/

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type FileConfig struct{}

type TokenConfig struct{}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" ||
		c.DbConfig.Name == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" {
		return errors.New("all environment required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
