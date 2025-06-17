package main

import (
	"log"

	"karedoro/presentation"
)

func main() {
	app, audioService := presentation.NewApp()
	if err := app.Run(audioService); err != nil {
		log.Fatal(err)
	}
}