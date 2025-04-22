package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)

	serv := server.NewServer(logger)

	logger.Println("Starting server on :8080")

	err := serv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Server failed: ", err)
	}
}
