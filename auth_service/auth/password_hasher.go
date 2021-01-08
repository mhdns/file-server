package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword takes the password and salt and returns the hashed password
func HashPassword(password, salt string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
}

// ComparePassword compares the password
func ComparePassword(givenPassword, salt, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(givenPassword+salt))
	if err != nil {
		return false
	}

	return true
}
