package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"karedoro/application"
	"karedoro/presentation"
)

func main() {
	// ebitenuiを使用したアプリケーションのテスト
	services := application.NewServices()
	app := presentation.NewEbitenUIApp(services)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Karedoro - EbitenUI Version")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}