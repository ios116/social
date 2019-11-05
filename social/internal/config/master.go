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

	MasterPassword string `env:"MASTER_PASSWORD" envDefault:"qwerty"`
	MasterUser     string `env:"MASTER_USER" envDefault:"soc_user"`
	MasterHost     string `env:"MASTER_HOST" envDefault:"master"`
	MasterPort     string `env:"MASTER_PORT" envDefault:"3306"`
	MasterName     string `env:"MASTER_DATABASE" envDefault:"soc_db"`
}

func NewDateBaseConf() *DateBaseConf {
	c := &DateBaseConf{}
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return c
}

// DBConnection - connection for BD master
func DBConnection(c *DateBaseConf) (*sqlx.DB, error) {
	// "user:password@tcp(127.0.0.1:3306)/hello"
	//  user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.MasterUser, c.MasterPassword, c.MasterHost, c.MasterPort, c.MasterName)
	fmt.Println("master=",dsn)
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}


