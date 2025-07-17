package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"clean_arc/config"
	"clean_arc/internal/controller/http"
	"clean_arc/internal/repo/persistent"
	"clean_arc/internal/repo/webapi"
	"clean_arc/internal/usecase/translation"
	"clean_arc/pkg/httpserver"
	"clean_arc/pkg/logger"
	"clean_arc/pkg/postgres"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found or couldn't be loaded. Continuing anyway.")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("main - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use-Case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, cfg, translationUseCase, l)

	// Start servers
	httpServer.Start()

	// Waiting for interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("main - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("main - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("main - httpServer.Shutdown: %w", err))
	}
}
