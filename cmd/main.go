package main

import (
	"errors"
	"fmt"
	"github.com/timur-pro/ueba-get-items/config"
	"github.com/timur-pro/ueba-get-items/internal/handlers"
	"github.com/timur-pro/ueba-get-items/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Конфигурация сервера
	cfg := config.GetConfig()

	// Создание хэндлера
	repo := repository.New(cfg.FilePath, cfg.IdColumnIndex)
	getItemHandler := handlers.New(repo)

	http.HandleFunc("/get-items", getItemHandler.Get)

	log.Printf("Running HTTP server on %s\n", cfg.ServerAddress)
	go func() {
		err := http.ListenAndServe(cfg.ServerAddress, nil)
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Ожидаем сигналы от системы для корректного завершения работы
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGKILL)

	<-sig
	fmt.Println("Shutting down server...")
	return nil
}
