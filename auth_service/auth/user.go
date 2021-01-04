package auth

// User is a structure to hold user details
type User struct {
	Email    string
	Password string
	Name     string
}

// NewUser returns a pointer to a User struct with email, password and name
func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
}

// Clone returns a copy of the user
func (user *User) Clone() *User {
	return &User{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}
}
