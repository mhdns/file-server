package auth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mhdns/web_server/auth_service/auth"
)

func TestNewJWTManager(t *testing.T) {
	jwtManager := auth.NewJWTManager("Secret", time.Hour)
	require.NotEmpty(t, jwtManager)
}

func TestGenerateToken(t *testing.T) {
	jwtManager := auth.NewJWTManager("Secret", time.Hour)

	token, err := jwtManager.Generate("random_email@gmail.com", "password", "name anme")
	require.NoError(t, err)
	require.NotZero(t, token)
}

func TestVerifyValidToken(t *testing.T) {
	jwtManager := auth.NewJWTManager("Secret", time.Hour)

	email := "random_email@gmail.com"
	password := "password"
	name := "name anme"
	token, err := jwtManager.Generate(email, password, name)
	require.NoError(t, err)
	require.NotZero(t, token)

	claims, err := jwtManager.Verify(token)
	require.NoError(t, err)
	require.Equal(t, claims.Email, email)
	require.Equal(t, claims.Password, password)
	require.Equal(t, claims.Name, name)
}

func TestVerifyInvalidToken(t *testing.T) {
	jwtManager := auth.NewJWTManager("Secret", time.Hour)

	token := "randomString"
	_, err := jwtManager.Verify(token)
	require.Error(t, err)
}
