package main

import (
	"log"

	"karedoro/presentation"
)

func main() {
	app := presentation.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}