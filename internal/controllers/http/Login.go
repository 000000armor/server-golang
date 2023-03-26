package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// in-memory "DB"
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// decode creds json
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check is user is there or no
	password, ok := users[creds.Username]

	// return Unauthorized code if user is not found or password not valid
	if !ok || password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, expirationTime, errorCode := GenerateToken(&creds)

	if errorCode != 0 {
		w.WriteHeader(errorCode)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	fmt.Printf("/login")
	io.WriteString(w, "/login")
}
