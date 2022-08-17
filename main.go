package main

import (
	"github.com/adnanmhd/go-porto-homies/app"
	"log"

	"github.com/adnanmhd/go-porto-homies/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error : %s", err)
	}

	// Run
	app.Run(cfg)
}
