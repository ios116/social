package storage

import (
	"database/sql"
	"social/internal/domain/entities"
)

func toUser(userDB *UserDB) *entities.User {
	var lastName string
	var firstName string
	var city string
	var interests string
	var gender string
	if userDB.FirstName.Valid {
		firstName = userDB.FirstName.String
	}
	if userDB.LastName.Valid {
		lastName = userDB.LastName.String
	}

	if userDB.City.Valid {
		city = userDB.City.String
	}
	if userDB.Interests.Valid {
		interests = userDB.Interests.String
	}
	if userDB.Gender.Valid {
		gender = userDB.Gender.String
	}

	return &entities.User{
		ID:          userDB.ID,
		Login:       userDB.Login,
		Password:    userDB.Password,
		Email:       userDB.Email,
		City:        city,
		Gender:      gender,
		Interests:   interests,
		DateCreated: userDB.DateCreated,
		DateModify:  userDB.DateModify,
		LastName:    lastName,
		FirstName:   firstName,
	}
}

func fromUser(user *entities.User) *UserDB {

	var firstName sql.NullString
	var lastName sql.NullString
	var city sql.NullString
	var gender sql.NullString
	var interests sql.NullString

	firstName.String = user.FirstName
	lastName.String = user.LastName
	city.String = user.City
	gender.String = user.Gender
	interests.String = user.Interests

	return &UserDB{
		ID:          user.ID,
		Login:       user.Login,
		Password:    user.Password,
		Email:       user.Email,
		Gender:      gender,
		Interests:   interests,
		City:        city,
		DateCreated: user.DateCreated,
		DateModify:  user.DateModify,
		LastName:    lastName,
		FirstName:   firstName,
	}
}
