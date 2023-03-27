package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// in-memory "DB"
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	var creds Credentials
	// decode creds json
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check is user is there or no
	passwordLocal, ok := users[username]

	// return Unauthorized code if user is not found or password not valid
	if !ok || passwordLocal != password {
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
}
