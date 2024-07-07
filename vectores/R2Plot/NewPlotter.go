package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	. "vectores/vectores"
)

func NewPlotter(vectores []Vector, animated bool, player bool) {
	width, height, scale := 640.0, 480.0, 1.5
	g := &Graficadora{
		width:  width,
		height: height,
	}
	g.setCoordinates()

	plotVectors := make([]PlotVector, len(vectores))
	for i, v := range vectores {
		plotVectors[i] = PlotVector{Vector: Vector{
			Inicio: Punto{v.GetStartX() + g.Origin.GetX(), -v.GetStartY() + g.Origin.GetY()},
			Fin:    Punto{v.GetEndX() + g.Origin.GetX(), -v.GetEndY() + g.Origin.GetY()},
		}, Color: randomColor()}
	}

	g.Vectores = plotVectors

	if player {
		g.AddPlayer()
	}

	g.Coordinates.xUnits = 20
	g.Coordinates.yUnits = 20

	g.xUnitLength = g.width / g.xUnits
	g.yUnitLength = g.height / g.yUnits

	ebiten.SetWindowSize(int(width*scale), int(height*scale))
	ebiten.SetWindowTitle("Graficadora de Vectores")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// Graficadora es la estructura principal que contiene los vectores a graficar
type Graficadora struct {
	Vectores     []PlotVector
	PlayerVector PlotVector
	screen       *ebiten.Image
	Coordinates
	width, height float64
	Animated      bool
}

type Coordinates struct {
	xUnits      float64
	xUnitLength float64
	yUnits      float64
	yUnitLength float64
	xSize       int
	ySize       int
	xZero       float64
	yZero       float64
	xTopRight   float64
	xTopLeft    float64
	yTopUp      float64
	yTopDown    float64
	left        Punto
	right       Punto
	Up          Punto
	Down        Punto
	Origin      Punto
}
