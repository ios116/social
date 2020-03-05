package storage

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"
)

// Storage - event repository
type Storage struct {
	Db      *sqlx.DB
	DbSlave *sqlx.DB
	TNT     *tarantool.Connection
}

func NewStorage(db *sqlx.DB, dbSlave *sqlx.DB, tar *tarantool.Connection) *Storage {
	return &Storage{Db: db, DbSlave: dbSlave, TNT: tar}
}
