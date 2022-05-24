package main

import (
	"log"

	"github.com/adnanmhd/go-porto-homies/config"
	"github.com/adnanmhd/go-porto-homies/internal/app"
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
