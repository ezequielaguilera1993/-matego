package R2Plotter

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	. "vectores/vectores"
)

type plotVector struct {
	Vector
	Color color.Color
}

func (g *Graficadora) Update() error {

	if g.AnimationEnabled == true {
		g.vectorsBehavior()
	}
	if g.PlayerEnabled {
		g.playerBehavior()
	}
	return nil
}

func (g *Graficadora) Draw(screen *ebiten.Image) {
	g.screen = screen

	if g.AxesEnabled {
		g.drawAxes()
	}
	if g.AxesDetailEnabled {
		g.drawAxesDetails()
	}

	if g.TobsEnabled {
		g.drawTobs()
	}
	if g.PlayerEnabled {
		g.drawPlayer()
	}
	g.setDebugClick()
	g.PlayerVector.Color = randomColor()

}

func (g *Graficadora) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func randomColor() color.Color {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func (g *Graficadora) setDebugClick() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		xRelative := x - int(g.width/2)
		yRelative := int(g.height/2) - y
		ebitenutil.DebugPrintAt(g.screen, fmt.Sprintf("(%.1f, %.1f)", float64(xRelative)/g.xUnitLength, float64(yRelative)/g.yUnitLength), x+20, y)
	}
}
