package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPassword []byte, stringPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(stringPassword))
}
