package main

import (
	"log"

	"github.com/denistakeda/mpass/internal/config"
)

func main() {
	conf, err := config.ParseServerCfg()
	if err != nil {
		log.Fatal("failed to read the configuration")
	}

	log.Printf("Configuration: %v\n", conf)
}
