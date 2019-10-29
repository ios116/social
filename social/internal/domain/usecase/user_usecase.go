package usecase

import (
	"context"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
	"time"
)

type UserService interface {
	AddUserUseCase(ctx context.Context, user *entities.User) (int64, error)
	UpdateUserUseCase(ctx context.Context, user *entities.User) error
	DeleteUserUseCase(ctx context.Context, ID int64) error
	GetUserByIdUseCase(ctx context.Context, ID int64) (*entities.User, error)
	GetUserByLoginUseCase(ctx context.Context, login string) (*entities.User, error)
	SetPasswordUseCase(ctx context.Context, password string, ID int64) error
	CheckAuthUseCase(ctx context.Context, login string, password string) (*entities.User, error)
	GetUsersWithLimitAndOffset(ctx context.Context, limit int64, offset int64) ([]*entities.User, error)
	FindByNameUC(ctx context.Context, query string) ([]*entities.User, error)
}

type Service struct {
	userRepository entities.UserRepository
}

func NewService(userRepository entities.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}

func (s *Service) CheckAuthUseCase(ctx context.Context, login string, password string) (*entities.User, error) {
	if login == "" {
		return nil, exceptions.LoginRequired
	}
	if password == "" {
		return nil, exceptions.PasswordRequired
	}
	user, err := s.userRepository.GetUserByLogin(ctx, login)

	if err != nil {
		return nil, exceptions.ObjectDoesNotExist
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, exceptions.Auth
	}
	return user, nil
}

func (s *Service) AddUserUseCase(ctx context.Context, user *entities.User) (int64, error) {

	if err := user.Validation(); err != nil {
		return 0, err
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.DateCreated = time.Now().UTC()
	user.DateModify = time.Now().UTC()
	user.Password = hash
	id, err := s.userRepository.AddUser(ctx, user)
	return id, err
}

func (s *Service) UpdateUserUseCase(ctx context.Context, user *entities.User) error {
	if err := user.Validation(); err != nil {
		return err
	}
	user.DateModify = time.Now().UTC()
	return s.userRepository.UpdateUser(ctx, user)
}

// DeleteUserUseCase
func (s *Service) DeleteUserUseCase(ctx context.Context, ID int64) error {
	return s.userRepository.DeleteUser(ctx, ID)
}

func (s *Service) GetUserByIdUseCase(ctx context.Context, ID int64) (*entities.User, error) {
	return s.userRepository.GetUserById(ctx, ID)
}

func (s *Service) GetUserByLoginUseCase(ctx context.Context, login string) (*entities.User, error) {
	return s.userRepository.GetUserByLogin(ctx, login)
}

func (s *Service) SetPasswordUseCase(ctx context.Context, password string, ID int64) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	modify := time.Now().UTC()
	return s.userRepository.SetPassword(ctx, hash, ID, modify)
}
func (s *Service) GetUsersWithLimitAndOffset(ctx context.Context, limit int64, offset int64) ([]*entities.User, error) {
	return s.userRepository.GetUsersWithLimitAndOffset(ctx, limit, offset)
}

func (s *Service) FindByNameUC(ctx context.Context, query string) ([]*entities.User, error) {
	if query == "" {
		return nil, exceptions.QueryRequired
	}
	return s.userRepository.FindByName(ctx, query)
}
