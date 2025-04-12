package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Harsh5840/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter , r *http.Request)  {
		w.Write([]byte("welcome to students api"))
	})

	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}
	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server")
	}
	
} 