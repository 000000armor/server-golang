package auth

import "golang.org/x/crypto/bcrypt"

func GetPasswordHash(passwordString string) (passwordHash []byte, err error) {
	password, errorHash := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		err = errorHash
		return
	}

	passwordHash = password
	return
}
