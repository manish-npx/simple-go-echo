package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Port int    `yaml:"port"`
	Addr string `yaml:"addr"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func LoadConfig() *Config {

	var cfg Config

	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error! config file not readable %v", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error parsing YAML file %v", err)
	}

	return &cfg
}
