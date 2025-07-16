package main

import (
	"random-service/internal/app"
	"random-service/pkg/logger"
)

func main() {
	log := logger.New()
	application := app.NewApp(log)

	if err := application.Run(); err != nil {
		log.Fatal("failed to start app: ", err)
	}
}
