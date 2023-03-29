package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// how to generate in prod env?
var jwtKey = []byte("my_secret_key")

// creds schema
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// claims schema
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(creds *Credentials) (token string, expirationTime time.Time, errGenerateToken error) {
	// generate expiration time
	expirationTime = time.Now().Add(time.Minute * 5)

	// why use pointer? 🤔
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// generate token with specific algo with claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	tokenString, err := rawToken.SignedString(jwtKey)
	token = tokenString

	if err != nil {
		errGenerateToken = err
		return
	}

	return
}
