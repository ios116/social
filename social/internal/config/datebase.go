package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// DateBaseConf config
type DateBaseConf struct {
	BdPassword string `env:"MYSQL_PASSWORD" envDefault:"qwerty"`
	BdUser     string `env:"MYSQL_USER" envDefault:"soc_user"`
	BdHost     string `env:"MYSQL_HOST" envDefault:"0.0.0.0"`
	BdPort     string `env:"MYSQL_PORT" envDefault:"3306"`
	BdName     string `env:"MYSQL_DATABASE" envDefault:"soc_db"`
}

func NewDateBaseConf() *DateBaseConf {
	c := &DateBaseConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}

// DBConnection - connection for BD
func DBConnection(c *DateBaseConf) (*sqlx.DB, error) {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	//  user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.BdUser, c.BdPassword, c.BdHost, c.BdPort, c.BdName)
	fmt.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}
