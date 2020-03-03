package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"social/internal/domain/entities"
	"time"
)

type MockedUseCaseUser struct {
	mock.Mock
}

func (m MockedUseCaseUser) AddUser(ctx context.Context, user *entities.User) (int64, error) {
	args := m.Called(ctx,user)
	return args.Get(0).(int64), args.Error(1)
}

func (m MockedUseCaseUser) UpdateUser(ctx context.Context, user *entities.User) error {
	panic("implement me")
}

func (m MockedUseCaseUser) DeleteUser(ctx context.Context, ID int64) error {
	panic("implement me")
}

func (m MockedUseCaseUser) GetUserById(ctx context.Context, ID int64) (*entities.User, error) {
	panic("implement me")
}

func (m MockedUseCaseUser) GetUserByLogin(ctx context.Context, login string) (*entities.User, error) {
	args := m.Called(ctx,login)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m MockedUseCaseUser) SetPassword(ctx context.Context, password string, ID int64, modify time.Time) error {
	panic("implement me")
}

func (m MockedUseCaseUser) GetUsersWithLimitAndOffset(ctx context.Context, limit int64, offset int64) ([]*entities.User, error) {
	panic("implement me")
}

func (m MockedUseCaseUser) FindByName(ctx context.Context, q string, id int64, limit int64, direction string) ([]*entities.User, error) {
	panic("implement me")
}
