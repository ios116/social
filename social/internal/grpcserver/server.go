package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"social/internal/config"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
	"social/internal/domain/usecase"
)

type RPCServer struct {
	UserService usecase.UserService
	Logger      *zap.Logger
	conf        *config.GrpcConf
}

// NewRPCServer constructor for new rpc server
func NewRPCServer(userService usecase.UserService, logger *zap.Logger, conf *config.GrpcConf) *RPCServer {
	return &RPCServer{UserService: userService, Logger: logger, conf: conf}
}

// Start - init RPC server
func (s *RPCServer) Start() {
	address := fmt.Sprintf("%s:%d", s.conf.GrpcHost, s.conf.GrpcPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.Logger.Fatal("Cannot start RPC server", zap.String("err", err.Error()))
	}
	// server := grpc.NewServer(grpc.UnaryInterceptor(newInterceptor(g.logger, g.conf.GrpcToken)))
	server := grpc.NewServer()
	RegisterUsersServer(server, s)
	s.Logger.Info("Starting RPC server", zap.String("address", address))
	err = server.Serve(lis)
	if err != nil {
		s.Logger.Fatal("Cannot start listen port", zap.String("err", err.Error()))
	}
}

// Add new user
func (s *RPCServer) AddUser(ctx context.Context, in *AddUserRequest) (*UserAddResponse, error) {
	user := &entities.User{
		Login:     in.Login,
		Password:  in.Password,
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		City:      in.City,
		Gender:    in.Gender,
		Interests: in.Interests,
	}
	ID, err := s.UserService.AddUserUseCase(ctx, user)

	switch err {
	case nil:
		return &UserAddResponse{
			Status: true,
			Detail: "Successes",
			UserId: ID,
		}, nil

	default:
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &UserAddResponse{
				Status: false,
				Detail: domainErr.Error(),
				UserId: 0,
			}, nil
		}
		s.Logger.Error("rpc", zap.String("err", err.Error()))
		return nil, err
	}
}

// Ger user by id
func (s *RPCServer) GetUserById(ctx context.Context, in *UserByIdRequest) (*UserResponse, error) {
	user, err := s.UserService.GetUserByIdUseCase(ctx, in.UserId)

	switch err {
	case nil:
		dateCreated, err := ptypes.TimestampProto(user.DateCreated)
		if err != nil {
			return nil, err
		}
		dateModify, err := ptypes.TimestampProto(user.DateModify)

		if err != nil {
			return nil, err
		}

		userRpc := &User{
			UserId:      user.ID,
			Login:       user.Login,
			Password:    user.Password,
			Email:       user.Email,
			City:        user.City,
			Gender:      user.Gender,
			Interests:   user.Interests,
			DateCreated: dateCreated,
			DateModify:  dateModify,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
		}
		return &UserResponse{
			Status: true,
			Detail: "Success",
			Date:   userRpc,
		}, nil
	default:
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &UserResponse{
				Status: false,
				Detail: domainErr.Error(),
				Date:   nil,
			}, nil
		}
		s.Logger.Error("rpc", zap.String("err", err.Error()))
		return nil, err
	}
}

// Update user
func (s *RPCServer) UpdateUser(ctx context.Context, in *UpdateUserRequest) (*StatusResponse, error) {
	dateCreated, err := ptypes.Timestamp(in.DateCreated)
	if err != nil {
		return nil, err
	}
	user := &entities.User{
		ID:          in.UserId,
		Login:       in.Login,
		Email:       in.Email,
		FirstName:   in.FirstName,
		DateCreated: dateCreated,
		LastName:    in.LastName,
		City:        in.City,
		Gender:      in.Gender,
		Interests:   in.Interests,
	}

	err = s.UserService.UpdateUserUseCase(ctx, user)
	switch err {
	case nil:
		return &StatusResponse{
			Status: true,
			Detail: "Successes",
		}, nil
	default:
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &StatusResponse{
				Status: false,
				Detail: domainErr.Error(),
			}, nil
		}

		return nil, err
	}
}

// Delete user by id
func (s *RPCServer) DeleteUser(ctx context.Context, in *UserByIdRequest) (*StatusResponse, error) {
	err := s.UserService.DeleteUserUseCase(ctx, in.UserId)
	switch err {
	case nil:
		return &StatusResponse{
			Status: true,
			Detail: "Successes",
		}, nil
	default:
		s.Logger.Error("rpc", zap.String("err", err.Error()))
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &StatusResponse{
				Status: false,
				Detail: domainErr.Error(),
			}, nil
		}
		return nil, err
	}
}

// Set password for new user
func (s *RPCServer) SetPassword(ctx context.Context, in *SetPasswordRequest) (*StatusResponse, error) {
	err := s.UserService.SetPasswordUseCase(ctx, in.Password, in.UserId)
	switch err {
	case nil:
		return &StatusResponse{
			Status: true,
			Detail: "Successes",
		}, nil
	default:
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &StatusResponse{
				Status: false,
				Detail: domainErr.Error(),
			}, nil
		}
		s.Logger.Error("rpc", zap.String("err", err.Error()))
		return nil, err
	}
}

// Check credential of an user
func (s *RPCServer) CheckAuth(ctx context.Context, in *CheckAuthRequest) (*UserResponse, error) {
	user, err := s.UserService.CheckAuthUseCase(ctx, in.Login, in.Password)
	switch err {
	case nil:
		dateCreated, err := ptypes.TimestampProto(user.DateCreated)
		if err != nil {
			return nil, err
		}
		dateModify, err := ptypes.TimestampProto(user.DateModify)

		if err != nil {
			return nil, err
		}
		userRpc := &User{
			UserId:      user.ID,
			Login:       user.Login,
			Password:    user.Password,
			Email:       user.Email,
			City:        user.City,
			Gender:      user.Gender,
			Interests:   user.Interests,
			DateCreated: dateCreated,
			DateModify:  dateModify,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
		}
		return &UserResponse{
			Status: true,
			Detail: "Success",
			Date:   userRpc,
		}, nil
	default:
		var domainErr exceptions.DomainError
		if errors.As(err, &domainErr) {
			return &UserResponse{
				Status: false,
				Detail: domainErr.Error(),
			}, nil
		}
		s.Logger.Error("rpc", zap.String("err", err.Error()))
		return nil, err
	}
}
