package auth

import (
	"context"
	"fmt"

	"github.com/mhdns/web_server/auth_service/auth_pb"
)

// AuthenticationServer server
type AuthenticationServer struct {
	auth_pb.UnimplementedAuthServiceServer
	store *InMemoryUserStore
}

// NewAuthenticationServer returns a pointer to an AuthenticationServer with new in memory user store
func NewAuthenticationServer() *AuthenticationServer {
	return &AuthenticationServer{
		store: NewInMemoryUserStore(),
	}
}

// Login describe
func (server *AuthenticationServer) Login(ctx context.Context, req *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	// Check email address is valid and user exists

	// Retrieve user from user store

	// Check Password is the same

	// Create JWT Token

	// return response

	return nil, fmt.Errorf("error")
}

// Register describe
func (server *AuthenticationServer) Register(ctx context.Context, req *auth_pb.RegisterRequest) (*auth_pb.RegisterResponse, error) {

	// Check if user already exist

	// Check if email is valid

	// Insert user into store

	// Create token

	// return token

	return nil, fmt.Errorf("error")
}
