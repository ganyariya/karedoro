package main

import (
	"log"

	"karedoro/application"
	"karedoro/presentation"
)

func main() {
	// Build dependency graph
	services := application.NewServices()
	
	// Create and run the application
	app := presentation.NewAppWithServices(services)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}