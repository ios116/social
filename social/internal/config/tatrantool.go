package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/tarantool/go-tarantool"
	"log"
	"time"
)

// TarantoolConf env for tarantool
type TarantoolConf struct {
	Port     int    `env:"TARANTOOL_PORT" envDefault:"3301"`
	Host     string `env:"TARANTOOL_HOST" envDefault:"0.0.0.0"`
	UserName string `env:"TARANTOOL_USER_NAME" envDefault:"admin"`
	Password string `env:"TARANTOOL_USER_PASSWORD" envDefault:"123456"`
}

func NewTarantoolConf() *TarantoolConf {
	c := &TarantoolConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}

func TarantoolConnection(conf *TarantoolConf) (*tarantool.Connection, error) {
	opts := tarantool.Opts{
		User:          conf.UserName,
		Pass:          conf.Password,
		Reconnect:     2 * time.Second,
		MaxReconnects: 3,
	}
	dsn := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	fmt.Println("tarantool=> ", dsn)
	conn, err := tarantool.Connect(dsn, opts)
	if err != nil {
		return nil, err
	}
	resp, err := conn.Ping()
	if err !=nil {
		return nil, err
	}
	log.Println("====>",resp.Code, resp.Data)
	return conn, nil
	//resp, err := conn.Insert(999, []interface{}{99999, "BB"})
	//if err != nil {
	//	fmt.Println("Error", err)
	//	fmt.Println("Code", resp.Code)
	//}
}
