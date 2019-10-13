package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

// HttpConf for config
type HttpConf struct {
	Port int    `env:"HTTP_PORT" envDefault:"8080"`
	Host string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
}

func NewHttpConf() *HttpConf {
	c := &HttpConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}
