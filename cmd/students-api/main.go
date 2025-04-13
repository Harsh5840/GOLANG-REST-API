package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Harsh5840/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1) //os.Signal is a buffer function so isme ek baar me ek hi value store ho skti hai

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server")
		}
	}() //function in go func is a function call

	<-done //jab tak done k andar signal nhi jayeg to yha block hojayega

	slog.Info("shutiign dow the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
