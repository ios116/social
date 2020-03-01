package users

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"
)

// UserStorage - event repository
type UserStorage struct {
	Db      *sqlx.DB
	DbSlave *sqlx.DB
	Tar     *tarantool.Connection
}

func NewUserStorage(db *sqlx.DB, dbSlave *sqlx.DB, tar *tarantool.Connection) *UserStorage {
	return &UserStorage{Db: db, DbSlave: dbSlave, Tar: tar}
}
