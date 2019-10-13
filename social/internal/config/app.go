package config

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"log"
)

// Additional conf
type AppConf struct {
	Build string `env:"APP_BUILD" envDefault:"dev"`
}

func NewAppConf() *AppConf {
	c := &AppConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}

// CreateLogger creating the logger
func CreateLogger(c *AppConf) (logger *zap.Logger, err error) {
	if c.Build == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	return logger, err
}
