package users

import (
	"context"
	"database/sql"
	"fmt"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
	"time"
)

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

type UserRole struct {
	ID     int64
	UserID int64 `Db:"user_id"`
	RoleID int64 `Db:"role_id"`
}

func (p *UserStorage) AddUser(ctx context.Context, user *entities.User) (int64, error) {
	query := "INSERT INTO users(login, password, email, city, gender, interests ,date_created,date_modify, first_name, last_name, age) VALUES(?,?, ?, ?, ?, ?, ?, ?, ?,?,?);"
	result, err := p.Db.ExecContext(ctx, query, user.Login, user.Password, user.Email, user.City, user.Gender, user.Interests, user.DateCreated, user.DateModify, user.FirstName, user.LastName, user.Age)
	switch err {
	case nil:
		id, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("sql: last insert id: %s", err)
		}
		return id, nil
	default:
		return 0, fmt.Errorf("sql: %s", err)
	}
}

func (p *UserStorage) UpdateUser(ctx context.Context, user *entities.User) error {
	query := "UPDATE users SET login = :login, email = :email, city = :city, gender = :gender, interests = :interests, date_created = :date_created, date_modify = :date_modify,first_name = :first_name, last_name = :last_name, age=:age  WHERE id=:id"

	result, err := p.Db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"login":        user.Login,
			"email":        user.Email,
			"gender":       user.Gender,
			"city":         user.City,
			"interests":    user.Interests,
			"date_created": user.DateCreated,
			"date_modify":  time.Now(),
			"id":           user.ID,
			"last_name":    user.LastName,
			"first_name":   user.FirstName,
			"age":          user.Age,
		})
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("sql: %s", exceptions.ObjectDoesNotExist)
	}
	return nil
}

func (p *UserStorage) DeleteUser(ctx context.Context, ID int64) error {
	result, err := p.Db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", ID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("sql: %s", exceptions.ObjectDoesNotExist)
	}
	return nil
}

func (p *UserStorage) GetUserByLogin(ctx context.Context, login string) (*entities.User, error) {

	query := "SELECT * FROM users WHERE login=?"
	dest := &UserDB{}
	err := p.Db.GetContext(ctx, dest, query, login)
	if err != nil {
		return nil, err
	}
	return toUser(dest), nil
}

func (p *UserStorage) GetUserById(ctx context.Context, ID int64) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id=?"
	dest := &UserDB{}
	err := p.Db.GetContext(ctx, dest, query, ID)
	switch err {
	case nil:
		return toUser(dest), nil
	case sql.ErrNoRows:
		return nil, fmt.Errorf("sql: %s", exceptions.ObjectDoesNotExist)
	default:
		return nil, err
	}
}

func (p *UserStorage) SetPassword(ctx context.Context, password string, ID int64, modify time.Time) error {
	query := "UPDATE users SET password = ?, date_modify=? WHERE id=?"
	result, err := p.Db.ExecContext(ctx, query, password, modify, ID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("sql: %s", exceptions.ObjectDoesNotExist)
	}
	return nil
}

func (p *UserStorage) GetUsersWithLimitAndOffset(ctx context.Context, limit int64, offset int64) ([]*entities.User, error) {
	query := "SELECT * FROM users ORDER BY id DESC LIMIT ? OFFSET ?"
	rows, err := p.Db.QueryxContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	user := &UserDB{}
	var users []*entities.User
	for rows.Next() {
		err := rows.StructScan(user)
		if err != nil {
			return nil, err
		}
		users = append(users, toUser(user))
	}
	return users, nil
}

func (p *UserStorage) FindByName(ctx context.Context, firstName string, lastName string) ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE first_name LIKE ? and last_name LIKE ? ORDER BY id DESC "
	rows, err := p.Db.QueryxContext(ctx, query, firstName+"%", lastName+"%")
	if err != nil {
		return nil, err
	}

	user := &UserDB{}
	var users []*entities.User
	for rows.Next() {
		err := rows.StructScan(user)
		if err != nil {
			return nil, err
		}
		users = append(users, toUser(user))
	}
	return users, nil
}
