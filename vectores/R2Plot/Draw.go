package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	. "vectores/vectores"
)

type PlotVector struct {
	Vector
	Color color.Color
}

// X get X(relative to the origin) grom any number
func (g *Graficadora) x(x float64) float64 {
	return (g.Origin.GetX() + x) / g.xUnitLength
}

// Y get Y(relative to the origin) grom any number
func (g *Graficadora) y(y float64) float64 {
	return (g.Origin.GetY() - y) / g.yUnitLength
}

func (g *Graficadora) Update() error {

	if g.Animated == true {
		g.vectorsBehavior()
	}

	g.playerBehavior()
	return nil
}

func (g *Graficadora) Draw(screen *ebiten.Image) {
	g.screen = screen

	g.drawAxes()
	g.drawAxesDetails()
	g.drawVectors()
	g.drawPlayer()
	g.PlayerVector.Color = randomColor()

}

func (g *Graficadora) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Graficadora) AddPlayer() {
	g.PlayerVector = PlotVector{Vector: Vector{
		Inicio: Punto{10 + g.Origin.GetX(), -10 + g.Origin.GetY()},
		Fin:    Punto{30 + g.Origin.GetX(), -40 + g.Origin.GetY()},
	}, Color: Green}
}

func randomColor() color.Color {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func (g *Graficadora) setCoordinates() {
	width := g.width
	height := g.height
	g.xZero, g.yZero = float64(width)/2, float64(height)/2
	g.xTopRight = float64(width)
	g.xTopLeft = 0
	g.yTopUp = 0
	g.yTopDown = float64(height)
	g.left = Punto{g.xTopLeft, g.yZero}
	g.right = Punto{g.xTopRight, g.yZero}
	g.Up = Punto{g.xZero, g.yTopUp}
	g.Down = Punto{g.xZero, g.yTopDown}
	g.Origin = Punto{g.xZero, g.yZero}
}
