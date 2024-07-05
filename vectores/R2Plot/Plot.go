package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	. "vectores/vectores"
)

// Graficadora es la estructura principal que contiene los vectores a graficar
type Graficadora struct {
	Vectores  []Vector
	screen    *ebiten.Image
	xZero     float64 // la mitad del width de la pantalla
	yZero     float64 // la mitad del height de la pantalla
	xTopRight float64 // la parte más a la derecha de la pantalla
	xTopLeft  float64 // la parte más a la izquierda de la pantalla
	yTopUp    float64 // la parte más alta de la pantalla
	yTopDown  float64 // la parte más baja de la pantalla
	left      Punto
	right     Punto
	Up        Punto
	Down      Punto
	Zero      Punto
}

func (g *Graficadora) Update() error {
	return nil
}

func (g *Graficadora) Draw(screen *ebiten.Image) {
	g.screen = screen
	g.setCoordinates()
	g.drawAxes()
	g.drawVectors()

}

func (g *Graficadora) setCoordinates() {
	width := g.screen.Bounds().Size().X
	height := g.screen.Bounds().Size().Y
	g.xZero, g.yZero = float64(width)/2, float64(height)/2
	g.xTopRight = float64(width)
	g.xTopLeft = 0
	g.yTopUp = 0
	g.yTopDown = float64(height)
	g.left = Punto{g.xTopLeft, g.yZero}
	g.right = Punto{g.xTopRight, g.yZero}
	g.Up = Punto{g.xZero, g.yTopUp}
	g.Down = Punto{g.xZero, g.yTopDown}
	g.Zero = Punto{g.xZero, g.yZero}
}

func (g *Graficadora) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

const (
	unidadesARepresentarX = 14
	unidadesARepresentarY = 14
)

func (g *Graficadora) drawAxes() {
	g.drawArrowBidirectional(Vector{Inicio: g.left, Fin: g.right}, Green)
	g.drawArrowBidirectional(Vector{Inicio: g.Up, Fin: g.Down}, Blue)
}

func NewPlotter(vectores []Vector) {
	graficadora := &Graficadora{Vectores: vectores}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Graficadora de Vectores")

	if err := ebiten.RunGame(graficadora); err != nil {
		log.Fatal(err)
	}
}
