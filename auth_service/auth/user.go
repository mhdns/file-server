package auth

// User is a structure to hold user details
type User struct {
	email    string
	password string
	name     string
}

// Clone returns a copy of the user
func (user *User) Clone() *User {
	return &User{
		email:    user.email,
		password: user.password,
		name:     user.name,
	}
}
