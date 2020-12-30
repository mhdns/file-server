package auth

import (
	"context"

	"github.com/mhdns/web_server/auth_service/auth_pb"
)

// AuthenticationServer server
type AuthenticationServer struct {
	auth_pb.UnimplementedAuthServiceServer
}

// Login describe
func (server *AuthenticationServer) Login(ctx context.Context, req *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	// Check email address is valid and user exists

	// Retrieve user from user store

	// Check Password is the same

	// Create JWT Token

	// return response

	return nil, nil
}

// Register describe
func (server *AuthenticationServer) Register(ctx context.Context, req *auth_pb.RegisterRequest) (*auth_pb.RegisterResponse, error) {

	// Check if user already exist

	// Check if email is valid

	// Insert user into store

	// Create token

	// return token

	return nil, nil
}
