package users

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// UserStorage - event repository
type UserStorage struct {
	Db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{Db: db}
}
