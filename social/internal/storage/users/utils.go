package users

import (
	"database/sql"
	"social/internal/domain/entities"
)

func toUser(userDB *UserDB) *entities.User {
	var lastName string
	var firstName string
	if userDB.FirstName.Valid {
		firstName = userDB.FirstName.String
	}
	if userDB.LastName.Valid {
		lastName = userDB.LastName.String
	}
	return &entities.User{
		ID:          userDB.ID,
		Login:       userDB.Login,
		Password:    userDB.Password,
		Email:       userDB.Email,
		IsActive:    userDB.IsActive,
		IsStaff:     userDB.IsStaff,
		DateCreated: userDB.DateCreated,
		DateModify:  userDB.DateModify,
		LastName:    lastName,
		FirstName:   firstName,
	}
}

func fromUser(user *entities.User) *UserDB {

	var firstName sql.NullString
	var lastName sql.NullString

	firstName.String = user.FirstName
	lastName.String = user.LastName

	return &UserDB{
		ID:          user.ID,
		Login:       user.Login,
		Password:    user.Password,
		Email:       user.Email,
		IsActive:    user.IsActive,
		IsStaff:     user.IsStaff,
		DateCreated: user.DateCreated,
		DateModify:  user.DateModify,
		LastName:    lastName,
		FirstName:   firstName,
	}
}
