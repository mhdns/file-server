package auth

import "sync"

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
	store map[string]*User
}

// NewInMemoryUserStore returns a InMemoryUserStore with an empty store
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		store: make(map[string]*User),
	}
}

// Save commits User into in-memory store
func (store *InMemoryUserStore) Save(user User) error {
	return nil
}

// Get retrieves a User based on email address and returns an error if user
// doesn't exist
func (store *InMemoryUserStore) Get(email string) (*User, error) {
	return nil, nil
}

// Update takes a pointer to the User and updates the details in the store
func (store *InMemoryUserStore) Update(*User) error {
	return nil
}

// Delete removes a User from the store
func (store *InMemoryUserStore) Delete(email string) error {
	return nil
}
