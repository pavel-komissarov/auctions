package main

import (
	"auctions/internal/logger"
	"auctions/internal/server/handlers/createAuction"
	"auctions/internal/server/handlers/makeBet"
	"auctions/internal/server/mw"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log := logger.New()

	log.Info("starting service")

	router := chi.NewMux()

	router.Use(mw.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/auctions", createAuction.New(log))
	router.Post("/auction/{id}", makeBet.New(log))

	log.Info("starting server", slog.String("address", "localhost:8888"))

	srv := &http.Server{
		Addr:         "localhost:8888",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
}
