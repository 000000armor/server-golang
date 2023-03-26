package application

import (
	"log"
	"server/internal/controllers/http"
)

type App struct {
}

func (a *App) Run() error {
	log.Println("run")
	server := http.New("3000")
	server.Run()
	return nil
}

func (a *App) Stop() error {
	return nil
}
