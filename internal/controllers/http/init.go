package http

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Adapter struct {
	port string
}

func New(port string) *Adapter {
	return &Adapter{port}
}

func (a *Adapter) Run() error {
	PORT := fmt.Sprintf(":%s", a.port)

	http.HandleFunc("/login", Login)
	http.HandleFunc("/verify", Verify)
	http.HandleFunc("/logout", Logout)

	err := http.ListenAndServe(PORT, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	return nil
}
