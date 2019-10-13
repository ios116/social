package entities

import (
	"context"
	"social/internal/domain/exceptions"
	"time"
)

type User struct {
	ID          int64
	Login       string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	IsActive    bool
	IsStaff     bool
	DateCreated time.Time
	DateModify  time.Time
}

func (u *User) Validation() error {
	if u.Login == "" {
		return exceptions.LoginRequired
	}

	if u.Email == "" {
		return exceptions.EmailRequired
	}
	return nil
}

type UserRepository interface {
	AddUser(ctx context.Context, user *User) (int64, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, ID int64) error
	GetUserById(ctx context.Context, ID int64) (*User, error)
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	SetPassword(ctx context.Context, password string, ID int64, modify time.Time) error
}
