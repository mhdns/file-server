package auth

import (
	"fmt"
	"sync"
)

// Store is an interface that holds in-memory data
type Store interface {
	Save(*User) error
	Get(string) (*User, error)
	Update(*User) error
	Delete(string) error
}

// InMemoryUserStore is a Store that holds User data in-memory
type InMemoryUserStore struct {
	mu    sync.Mutex // learn about embedding
	Store map[string]*User
}

// NewInMemoryUserStore returns a InMemoryUserStore with an empty store
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		Store: make(map[string]*User),
	}
}

// Save commits User into in-memory store
func (store *InMemoryUserStore) Save(user *User) error {
	_, ok := store.Store[user.Email]
	if ok {
		return fmt.Errorf("user already exits")
	}

	store.Store[user.Email] = user
	return nil
}

// Get retrieves a User based on email address and returns an error if user
// doesn't exist
func (store *InMemoryUserStore) Get(email string) (*User, error) {
	if user, ok := store.Store[email]; ok {
		return user, nil
	}

	return nil, fmt.Errorf("user not found")
}

// Update takes a pointer to the User and updates the details in the store
func (store *InMemoryUserStore) Update(user *User) error {
	if user, ok := store.Store[user.Email]; ok {
		store.Store[user.Email] = user
		return nil
	}
	return fmt.Errorf("user not found")
}

// Delete removes a User from the store
func (store *InMemoryUserStore) Delete(email string) error {
	if _, ok := store.Store[email]; ok {
		delete(store.Store, email)
		return nil
	}
	return fmt.Errorf("user not found")
}
