package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"server/internal/application"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)

	defer cancel()

	app := application.App{}
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()

	app.Stop()
}
