package main

import (
	"fmt"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image

const (
	screenWidth  = 640
	screenHeight = 480
)

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("ebi.png")
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct {
	x    float64
	y    float64
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	keyStrs := []string{}
	for _, p := range g.keys {
		if p.String() != "" {
			keyStrs = append(keyStrs, p.String())
		}
	}
	switch strings.Join(keyStrs, ", ") {
	case ebiten.KeyArrowRight.String():
		g.x += float64(2)
	case ebiten.KeyArrowLeft.String():
		g.x -= float64(2)
	case ebiten.KeyArrowUp.String():
		g.y -= float64(2)
	case ebiten.KeyArrowDown.String():
		g.y += float64(2)
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("ebi position\n(%0.2f, %0.2f)", g.x, g.y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{x: 0.0, y: 0.0}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("moving ebi!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
