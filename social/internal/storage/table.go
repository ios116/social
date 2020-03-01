package storage

import (
	"database/sql"
	"time"
)
// UserDB is table for user in postgres
type UserDB struct {
	ID          int64
	Login       string
	Password    string
	Email       string
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	City        sql.NullString
	Gender      sql.NullString
	Interests   sql.NullString
	DateCreated time.Time `db:"date_created"`
	DateModify  time.Time `db:"date_modify"`
	Age         int32
}

// UserRole  is table for users role in postgres
type UserRole struct {
	ID     int64
	UserID int64 `db:"user_id"`
	RoleID int64 `db:"role_id"`
}

// Subscribe  is table for subscribe user to user in postgres
type SubscribeDb struct {
   ID int64 `db:"id"`
   UserId int64 `db:"user_id"`
   SubscribeId int64 `db:"subscribe_id"`
}