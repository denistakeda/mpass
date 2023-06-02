package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denistakeda/mpass/internal/client"
	"github.com/denistakeda/mpass/internal/client_service"
	"github.com/denistakeda/mpass/internal/client_storage"
	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/printer"
	"github.com/denistakeda/mpass/internal/scanner"
)

func main() {
	homeFolder := fmt.Sprintf("%s/.mpass/", os.Getenv("HOME"))
	statePath := fmt.Sprintf("%s/state.gob", homeFolder)
	configPath := fmt.Sprint("%s/config.json", homeFolder)

	conf, err := config.ParseClientCfg(configPath)
	if err != nil {
		log.Fatal(err)
	}

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
