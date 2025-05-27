package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andreyxaxa/Quotation-Book-service/config"
	"github.com/andreyxaxa/Quotation-Book-service/internal/controller/http"
	"github.com/andreyxaxa/Quotation-Book-service/internal/repo/persistent"
	"github.com/andreyxaxa/Quotation-Book-service/internal/usecase/quotes"
	"github.com/andreyxaxa/Quotation-Book-service/pkg/httpserver"
)

func Run(cfg *config.Config) {
	// Repo
	rep := persistent.New()

	// Use-Case
	quotesUseCase := quotes.New(rep)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	http.NewRouter(httpServer.Router, quotesUseCase)

	// Start server
	httpServer.Start()
	log.Printf("http server started at :%s", cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	log.Println("graceful shutdown complete")
}
