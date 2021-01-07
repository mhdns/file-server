package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mhdns/web_server/auth_service/auth"
	"github.com/mhdns/web_server/auth_service/auth_pb"
)

func TestLogin(t *testing.T) {
	server := auth.NewAuthenticationServer(time.Hour, "Secret")

	email := "sample@gmail.com"
	password := "randomP"
	name := "Joker Tan"

	registerUser(email, password, name, t, server)

	req := &auth_pb.LoginRequest{
		LoginDate: time.Now().Unix(),
		LoginCred: &auth_pb.LoginCred{
			Email:    "sample@gmail.com",
			Password: "randomP",
		},
	}
	res, err := server.Login(context.Background(), req)
	require.NoError(t, err)
	require.IsType(t, "", res.GetToken().GetToken())
}

func TestRegister(t *testing.T) {
	email := "sample@gmail.com"
	password := "randomP"
	name := "Joker Tan"

	registerUser(email, password, name, t, nil)
}

func registerUser(email, password, name string, t *testing.T, server *auth.AuthenticationServer) {
	if server == nil {
		server = auth.NewAuthenticationServer(time.Hour, "Secret")
	}

	req := &auth_pb.RegisterRequest{
		RegisterDate: time.Now().Unix(),
		RegisterDetails: &auth_pb.RegisterDetails{
			Email:     email,
			Password:  password,
			Name:      name,
			Birthdate: time.Date(1993, time.Month(6), 11, 0, 0, 0, 0, time.UTC).Unix(),
		},
	}
	res, err := server.Register(context.Background(), req)
	require.NoError(t, err)
	require.IsType(t, "", res.GetToken().GetToken())
}
