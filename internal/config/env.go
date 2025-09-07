package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var SvcCfg svcConfig

type Server struct {
	Addr string `env:"ADDR" envDefault:":7000"`
	Name string `env:"NAME" envDefault:"full-time-go-dev"`
}

type Database struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"full-time-go-dev"`
	Schema   string `env:"SCHEMA" envDefault:"public"`
	SSLMode  string `env:"SSL_MODE" envDefault:"disable"`
}

type svcConfig struct {
	Environment string   `env:"ENVIRONMENT" envDefault:"development"`
	Server      Server   `envPrefix:"SERVER_"`
	Database    Database `envPrefix:"DATABASE_"`
	Jwt         Jwt      `envPrefix:"JWT_"`
}

type Jwt struct {
	Secret string `env:"SECRET" envDefault:"secret"`
}

func init() {
	log := zap.S().Named("config")
	if err := godotenv.Load(); err != nil {
		log.Errorf("Can not read env from file system, please check the right this program owned.")
	}

	SvcCfg = svcConfig{}

	if err := env.Parse(&SvcCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}
}
