package auth

import (
	"context"
	"regexp"
	"time"

	"github.com/mhdns/web_server/auth_service/auth_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emailRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
)

// AuthenticationServer server
type AuthenticationServer struct {
	auth_pb.UnimplementedAuthServiceServer
	store *InMemoryUserStore
	jwt   *JWTManager
}

// NewAuthenticationServer returns a pointer to an AuthenticationServer with new in memory user store
func NewAuthenticationServer(tokenDuration time.Duration, jwtSecret string) *AuthenticationServer {
	return &AuthenticationServer{
		store: NewInMemoryUserStore(),
		jwt:   NewJWTManager(jwtSecret, tokenDuration),
	}
}

// Login describe
func (server *AuthenticationServer) Login(ctx context.Context, req *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {

	email := req.GetLoginCred().GetEmail()
	password := req.GetLoginCred().GetPassword()

	// Retrieve user from user store
	// Check email address is valid and user exists
	user, err := server.store.Get(email)
	if user == nil || err != nil {
		return nil, status.Errorf(codes.Internal, "invalid credentials")
	}

	// Check Password is the same
	if user.Password != password {
		return nil, status.Errorf(codes.Internal, "invalid credentials")
	}
	// Create JWT Token
	token, err := server.jwt.Generate(user.Email, user.Password, user.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create token: %v", err)
	}
	// create response
	res := &auth_pb.LoginResponse{
		Token: &auth_pb.Token{
			Token: token,
		},
	}

	return res, nil
}

// Register describe
func (server *AuthenticationServer) Register(ctx context.Context, req *auth_pb.RegisterRequest) (*auth_pb.RegisterResponse, error) {

	email := req.GetRegisterDetails().GetEmail()
	password := req.GetRegisterDetails().GetPassword()
	name := req.GetRegisterDetails().GetName()
	// birthdate := req.GetRegisterDetails().GetBirthdate()

	// Check if user already exist
	user, err := server.store.Get(email)
	if user != nil {
		return nil, status.Errorf(codes.Internal, "user already exits")
	}
	if err != nil && err.Error() != "user not found" {
		return nil, status.Errorf(codes.Internal, "unable to check if user exits: %v", err)
	}
	// Check if email is valid
	validEmail, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate email: %v", err)
	}
	if !validEmail {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email format: %v", err)
	}
	// Insert user into store
	user = NewUser(email, password, name)
	err = server.store.Save(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save user: %v", err)
	}
	// Create token
	token, err := server.jwt.Generate(email, password, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create token: %v", err)
	}

	// create response
	res := &auth_pb.RegisterResponse{
		Token: &auth_pb.Token{
			Token: token,
		},
	}

	return res, nil
}
