package usecase

import (
	"context"
	"social/internal/domain/entities"
	"testing"
	"time"
)

func TestUseCase(t *testing.T) {
	// Assert that Bar() is invoked.
	repoUser := new(MockedUseCaseUser)
	service := NewService(repoUser)
	ctx := context.Background()
	user := &entities.User{
		ID:          1,
		Login:       "Login",
		Password:    "123456",
		Email:       "site@mail.ru",
		DateCreated: time.Now().UTC(),
		DateModify:  time.Now().UTC(),
		Age:         33,
	}

	t.Run("AddUser", func(t *testing.T) {
         repoUser.On("AddUser",ctx,user).Return(int64(24),nil)
		_, err := service.AddUserUseCase(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("CheckAuthUseCase", func(t *testing.T) {
		user.Password, _ = HashPassword("123456")
		repoUser.On("GetUserByLogin",ctx,"Admin").Return(user,nil)
		user, err := service.CheckAuthUseCase(ctx, "Admin", "123456")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)
	})
}
