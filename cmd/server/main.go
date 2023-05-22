package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/server"
)

func main() {
	conf, err := config.ParseServerCfg()
	if err != nil {
		log.Fatal("failed to read the configuration")
	}
	log.Printf("Configuration: %v\n", conf)

	interruptChan := handleInterrupt()

	srv := server.New(server.NewServerParams{
		Host: conf.Host,
	})
	serverErrors := srv.Start()
	defer srv.Stop()

	select {
	case serverError := <-serverErrors:
		log.Println(serverError)
	case <-interruptChan:
		log.Println("Server was interrupted")
	}
}

func handleInterrupt() <-chan os.Signal {
	out := make(chan os.Signal, 2)
	signal.Notify(out, os.Interrupt)
	signal.Notify(out, syscall.SIGTERM)
	return out
}
