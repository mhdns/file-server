package auth_test

import (
	"testing"

	"github.com/mhdns/web_server/auth_service/auth"
	"github.com/stretchr/testify/require"
)

func TestNewInMemoryUserStore(t *testing.T) {
	userStore := auth.NewInMemoryUserStore()
	require.NotNil(t, userStore)
	require.Empty(t, userStore.Store)
	require.IsType(t, map[string]*auth.User{}, userStore.Store)
}

func TestSave(t *testing.T) {
	userStore, user, _ := createAndSaveUser(t)
	err := userStore.Save(user)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))
}

func TestGet(t *testing.T) {
	userStore, user, email := createAndSaveUser(t)
	userFromStore, err := userStore.Get(email)
	require.NoError(t, err)
	require.Equal(t, user, userFromStore)
}

func TestUpdate(t *testing.T) {
	userStore, _, email := createAndSaveUser(t)

	updatedUser := auth.NewUser(email, "random", "Jack Ma")
	err := userStore.Update(updatedUser)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))

}

func TestDelete(t *testing.T) {
	userStore, _, email := createAndSaveUser(t)
	err := userStore.Delete(email)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))
}

func createAndSaveUser(t *testing.T) (*auth.InMemoryUserStore, *auth.User, string) {
	userStore := auth.NewInMemoryUserStore()
	email := "random@gmail.com"
	password := "password"
	name := "Randy Random"
	user := auth.NewUser(email, password, name)

	return userStore, user, email
}
