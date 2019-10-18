package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

// HttpConf for config
type HttpConf struct {
	Port        int    `env:"HTTP_PORT" envDefault:"8080"`
	Host        string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	SessionKey  string `env:"SESSION_KEY" envDefault:"Auth"`
	SessionTime int    `env:"SESSION_TIME" envDefault:"24"`
	ContextKey  string `env:"CONTEXT_KEY" envDefault:"user"`
}

func NewHttpConf() *HttpConf {
	c := &HttpConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}
