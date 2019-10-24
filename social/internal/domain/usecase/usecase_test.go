package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"social/internal/domain/entities"
	"testing"
	"time"
)

func TestUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()
	repoUser := NewMockUserRepository(ctrl)

	service := NewService(repoUser)
	ctx := context.Background()

	user := &entities.User{
		ID:          1,
		Login:       "Login",
		Password:    "123456",
		Email:       "site@mail.ru",
		DateCreated: time.Now().UTC(),
		DateModify:  time.Now().UTC(),
		Age: 33,
	}

	t.Run("AddUser", func(t *testing.T) {
		repoUser.EXPECT().AddUser(ctx, user).Return(int64(24), nil).AnyTimes()
		_, err := service.AddUserUseCase(ctx, user)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CheckAuthUseCase", func(t *testing.T) {
		user.Password, _ = HashPassword("123456")
		repoUser.EXPECT().GetUserByLogin(ctx, "Admin").Return(user, nil).AnyTimes()
		user, err := service.CheckAuthUseCase(ctx, "Admin", "123456")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)

	})

}
