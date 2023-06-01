package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denistakeda/mpass/internal/client"
	"github.com/denistakeda/mpass/internal/client_service"
	"github.com/denistakeda/mpass/internal/client_storage"
	"github.com/denistakeda/mpass/internal/printer"
	"github.com/denistakeda/mpass/internal/scanner"
)

func main() {
	statePath := fmt.Sprintf("%s/.mpass/state.gob", os.Getenv("HOME"))

	clientStorage := client_storage.New(statePath)
	defer clientStorage.Close()

	clientService := client_service.New(clientStorage)

	printer := printer.New(os.Stdout, os.Stderr)
	scanner := scanner.New(os.Stdin)

	c := client.New(client.NewClientParams{
		Printer:       printer,
		Scanner:       scanner,
		ClientService: clientService,
	})

	if err := c.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
