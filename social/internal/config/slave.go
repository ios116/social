package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// SlaveConf config
type SlaveConf struct {
	SlavePassword string `env:"SLAVE_PASSWORD" envDefault:"qwerty"`
	SlaveUser     string `env:"SLAVE_USER" envDefault:"soc_user"`
	SlaveHost     string `env:"SLAVE_HOST" envDefault:"slave"`
	SlavePort     string `env:"SLAVE_PORT" envDefault:"3306"`
	SlaveName     string `env:"SLAVE_DATABASE" envDefault:"soc_db"`
}

// NewSlaveConf new config for master
func NewSlaveConf() *SlaveConf {
	c := &SlaveConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}

// SlaveConnection - connection for BD salve
func SlaveConnection(c *SlaveConf) (*sqlx.DB, error) {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	//  user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.SlaveUser, c.SlavePassword, c.SlaveHost, c.SlavePort, c.SlaveName)
	fmt.Println("SlaveConnection ===>",dsn)
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}


