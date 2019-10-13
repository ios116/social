package grpcserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"integration_tests/internal/config"
	"testing"
)

type tokenAuth struct {
	Token string
}

func (t *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.Token,
	}, nil
}

func (t *tokenAuth) RequireTransportSecurity() bool {
	return false
}

func TestGrpc(t *testing.T) {
	option := grpc.WithPerRPCCredentials(&tokenAuth{"Bearer secret"})
	conf := config.NewGrpcConf()
	address := fmt.Sprintf("%s:%d", conf.GrpcHost, conf.GrpcPort)
	conn, err := grpc.Dial(address, option, grpc.WithInsecure())
	if err != nil {
		t.Fatal("Can't connect to GRPC: ", address)
	}

	server := NewUsersClient(conn)
	ctx := context.Background()

	user:=&User{
		Login:                "login2",
		Password:             "123456",
		Email:                "site@mail.ru",
		IsActive:             false,
		IsStaff:              false,
		FirstName:            "Ivanov",
		LastName:             "Popov",
	}


	t.Run("add", func(t *testing.T) {
		newUser := &AddUserRequest{
			Login:     user.Login,
			Password:  user.Password,
			Email:     user.Email,
			IsActive:  false,
			IsStaff:   false,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		userRPC, err := server.AddUser(ctx, newUser)
		if err != nil {
			t.Fatal("test add:",err)
		}
		t.Log(userRPC.UserId)
		user.UserId = userRPC.UserId
	})


	t.Run("Get user by id", func(t *testing.T) {
		byId := &UserByIdRequest{
			UserId: user.UserId,
		}
		ctx = context.WithValue(ctx, "id", "Some Test")
		userResp, err := server.GetUserById(ctx, byId)
		if err != nil {
			t.Fatal(err)
		}
		user = userResp.Date
		if user.FirstName != "Ivanov" {
			t.Fatal(err)
		}
	})

	t.Log(user)

	t.Run("update", func(t *testing.T) {
		userUpdate := &UpdateUserRequest{
			UserId:      user.UserId,
			Login:       user.Login,
			Email:       user.Email,
			IsActive:    true,
			IsStaff:     user.IsStaff,
			DateCreated: user.DateCreated,
			FirstName:   "Vladimir",
			LastName:    "Popov",
		}

		st, err := server.UpdateUser(ctx, userUpdate)
		if err != nil {
			t.Fatal("test update:", err)
		}
		t.Log(st)
	})

	t.Run("Set password", func(t *testing.T) {
		newPass := SetPasswordRequest{
			Password: "23456789",
			UserId:   user.UserId,
		}
		st, err := server.SetPassword(ctx, &newPass)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(st)
	})

	t.Run("Auth", func(t *testing.T) {
		authReq := CheckAuthRequest{
			Login:    user.Login,
			Password: "23456789",
		}
		st, err := server.CheckAuth(ctx, &authReq)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(st)
	})

	t.Run("Delete user", func(t *testing.T) {
		delUser := UserByIdRequest{
			UserId: user.UserId,
		}
		_, err := server.DeleteUser(ctx, &delUser)
		if err != nil {
			t.Fatal(err)
		}
	})

}
