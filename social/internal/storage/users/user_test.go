package users

import (
	"context"
	"go.uber.org/dig"
	"social/internal/config"
	"social/internal/domain/entities"
	"testing"
	"time"
)

func TestUserStore(t *testing.T) {
	var pg *UserStorage
	container := dig.New()
	container.Provide(config.NewDateBaseConf)
	container.Provide(config.DBConnection)
	container.Provide(NewUserStorage)
	err := container.Invoke(func(st *UserStorage) {
		pg = st
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	user := &entities.User{
		Login:       "Admin23",
		Password:    "13456",
		Email:       "site@mail.ru",
		FirstName:   "Tomas",
		LastName:    "Jonson",
		City:        "Kazan",
		Gender:      "female",
		Interests:   "Some interest",
		DateCreated: time.Now().UTC(),
		DateModify:  time.Now().UTC(),
	}
	t.Run("AddUser", func(t *testing.T) {
		ID, err := pg.AddUser(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
		user.ID = ID
	})

	t.Run("GetUserById", func(t *testing.T) {
		userByID, err := pg.GetUserById(ctx, user.ID)
		if err != nil {
			t.Fatal(err)
		}
		if userByID.Login == "" {
			t.Fatal("user login is empty")
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		user.Email = "new@mail.ru"
		user.Login = "admin3"
		err := pg.UpdateUser(ctx, user)
		if err != nil {
			t.Fatal(err)
		}

		userByID, err := pg.GetUserById(ctx, user.ID)
		if err != nil {
			t.Fatal(err)
		}
		if userByID.Email != "new@mail.ru" {
			t.Fatal("user email is not equal")
		}
	})

	t.Run("SetPassword", func(t *testing.T) {
		err := pg.SetPassword(ctx, "new", user.ID, time.Now().UTC())
		if err != nil {
			t.Fatal(err)
		}
		userByID, err := pg.GetUserById(ctx, user.ID)
		if err != nil {
			t.Fatal(err)
		}
		if userByID.Password != "new" {
			t.Fatal("user password is not equal")
		}
	})

	t.Run("Select users", func(t *testing.T) {
		users, err := pg.GetUsersWithLimitAndOffset(ctx, 10, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(users) == 0 {
			t.Fatal("users must be more then 0")
		}
	})



	t.Run("Get by name", func(t *testing.T) {
		users, err := pg.FindByName(ctx, "Tom")
		if err != nil {
			t.Fatal(err)
		}
		if len(users) == 0 {
			t.Fatal("users must be =  2 ")
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		err = pg.DeleteUser(ctx, user.ID)
		if err != nil {
			t.Fatal(err)
		}
	})
}
