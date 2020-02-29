package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

// TarantoolConf env for tarantool
type TarantoolConf struct {
	Port     int    `env:"TARANTOOL_PORT" envDefault:"3301"`
	Host     string `env:"TARANTOOL_HOST" envDefault:"0.0.0.0"`
	UserName string `env:"TARANTOOL_USER_NAME" envDefault:"admin"`
	Password string `env:"TARANTOOL_USER_PASSWORD" envDefault:"123456"`
}

func NewTarantoolConf () *TarantoolConf {
	c := &TarantoolConf {}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}
