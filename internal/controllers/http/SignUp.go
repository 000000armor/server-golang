package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(creds.Username) == 0 || len(creds.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users[creds.Username] = creds.Password
	fmt.Println(users)
}
