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
	userStore, user, _ := createUserAndStore(t)
	err := userStore.Save(user)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))
}

func TestSaveDuplicate(t *testing.T) {
	userStore, user, _ := createUserAndStore(t)
	err := userStore.Save(user)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))

	err = userStore.Save(user)
	require.Error(t, err)

}

func TestGet(t *testing.T) {
	userStore, user, email := createUserAndStore(t)
	err := userStore.Save(user)
	require.NoError(t, err)

	userFromStore, err := userStore.Get(email)
	require.NoError(t, err)
	require.Equal(t, user, userFromStore)
}

func TestGetUserMissing(t *testing.T) {
	// Don't save user to check if error is returned when queried for the user
	userStore, _, email := createUserAndStore(t)
	_, err := userStore.Get(email)
	require.Error(t, err)
}

func TestUpdate(t *testing.T) {
	userStore, user, email := createUserAndStore(t)
	err := userStore.Save(user)
	require.NoError(t, err)

	updatedUser := auth.NewUser(email, "random", "Jack Ma")
	err = userStore.Update(updatedUser)
	require.NoError(t, err)
	require.Equal(t, 1, len(userStore.Store))

}

func TestUpdateUserMissing(t *testing.T) {
	userStore, _, email := createUserAndStore(t)

	updatedUser := auth.NewUser(email, "random", "Jack Ma")
	err := userStore.Update(updatedUser)
	require.Error(t, err)

}

func TestDelete(t *testing.T) {
	userStore, user, email := createUserAndStore(t)
	err := userStore.Save(user)
	require.NoError(t, err)

	err = userStore.Delete(email)
	require.NoError(t, err)
	require.Equal(t, 0, len(userStore.Store))
}

func TestDeleteUserMissing(t *testing.T) {
	userStore, _, email := createUserAndStore(t)
	err := userStore.Delete(email)
	require.Error(t, err)
}

func createUserAndStore(t *testing.T) (*auth.InMemoryUserStore, *auth.User, string) {
	userStore := auth.NewInMemoryUserStore()
	email := "random@gmail.com"
	password := "password"
	name := "Randy Random"
	user := auth.NewUser(email, password, name)

	return userStore, user, email
}
