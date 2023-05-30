package main

import (
	"log"
	"os"

	"github.com/denistakeda/mpass/internal/client"
	"github.com/denistakeda/mpass/internal/printer"
	"github.com/denistakeda/mpass/internal/scanner"
)

func main() {
	printer := printer.New(os.Stdout, os.Stderr)
	scanner := scanner.New(os.Stdin)

	c := client.New(client.NewClientParams{
		Printer: printer,
		Scanner: scanner,
	})

	if err := c.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
