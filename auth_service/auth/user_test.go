package auth_test

import (
	"testing"

	"github.com/mhdns/web_server/auth_service/auth"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	user := auth.NewUser("email@monkey.com", "password", "name name")
	require.NotNil(t, user)
}

func TestClone(t *testing.T) {
	user := auth.NewUser("email@monkey.com", "password", "name name")
	clonedUser := user.Clone()
	require.NotSame(t, user, clonedUser)
	require.Equal(t, user.Email, clonedUser.Email)
	require.Equal(t, user.Password, clonedUser.Password)
	require.Equal(t, user.Name, clonedUser.Name)

}
