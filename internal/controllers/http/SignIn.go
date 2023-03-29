package http

import (
	"fmt"
	"net/http"
	"server/domain/auth"
	"server/storage"
)

var userMock = map[string]string{
	"username": "username1",
	"password": "password1",
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials
	// create instance of local DB
	mydb := storage.NewMapDb()

	// set current creds from response
	username, password, ok := r.BasicAuth()
	creds.Username = username
	creds.Password = password

	passwordHash, err := auth.GetPasswordHash(userMock["password"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mydb.Set(userMock["username		"], passwordHash)

	hashedPassword, ok := mydb.Get(creds.Username).([]byte)

	// return Unauthorized code if user is not found or password not valid
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := auth.ComparePasswords(hashedPassword, creds.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, expirationTime, errGenerateToken := auth.GenerateToken(&creds)

	if errGenerateToken != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	fmt.Printf("/login")
}
