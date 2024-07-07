package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	. "vectores/vectores"
)

func NewPlotter(vectores []Vector, configs ...Config) {
	width, height, scale := 640.0, 480.0, 1.5
	g := &Graficadora{
		screenValues: screenValues{
			width:  width,
			height: height,
		},
	}

	config := configs[0]

	g.PlayerEnabled = config.PlayerEnabled
	g.AnimationEnabled = config.AnimationEnabled
	g.AxesEnabled = config.AxesEnabled
	g.AxesDetailEnabled = config.AxesDetailEnabled
	g.TobsEnabled = config.TobsEnabled
	g.coordinates.xUnits = float64(config.ResolutionX)
	g.coordinates.yUnits = float64(config.ResolutionY)

	g.setCoordinates()
	g.setPlotterAbsoluteValues()

	plotVectors := make([]plotVector, len(vectores))
	for i, v := range vectores {
		plotVectors[i] = g.newPlotVector(v)
	}
	g.Tobs = plotVectors

	if g.PlayerEnabled {
		g.PlayerVector = g.newPlotVector(Vector{
			Inicio: Punto{1, 1},
			Fin:    Punto{2, 2},
		})
	}

	ebiten.SetWindowSize(int(width*scale), int(height*scale))
	ebiten.SetWindowTitle("Graficadora de Tobs")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Graficadora) newPlotVector(v Vector) plotVector {
	return plotVector{Vector: v, Color: randomColor()}
}

// Graficadora es la estructura principal que contiene los vectores a graficar
type Graficadora struct {
	Tobs         []plotVector
	PlayerVector plotVector
	screenValues
	plotterAbsoluteValues
	coordinates
	Config
}

type screenValues struct {
	width, height float64
	screen        *ebiten.Image
}

type Config struct {
	PlayerEnabled, AxesEnabled, AxesDetailEnabled, AnimationEnabled, TobsEnabled bool
	ResolutionX, ResolutionY                                                     int
}

type plotterAbsoluteValues struct {
	xUnitLength, yUnitLength, absoluteOriginX, absoluteOriginY float64
	absoluteOrigin                                             Punto
}

type coordinates struct {
	xUnits    float64
	yUnits    float64
	xTopRight float64
	xTopLeft  float64
	yTopUp    float64
	yTopDown  float64
	left      Punto
	right     Punto
	Up        Punto
	Down      Punto
	Origin    Punto
}

func (g *Graficadora) setPlotterAbsoluteValues() {
	g.xUnitLength = g.width / g.xUnits
	g.yUnitLength = g.height / g.yUnits
	g.absoluteOrigin = Punto{g.width / 2, g.height / 2}
	g.absoluteOriginX = g.absoluteOrigin.GetX()
	g.absoluteOriginY = g.absoluteOrigin.GetY()
}

func (g *Graficadora) setCoordinates() {

	xUnitsHalf := g.xUnits / 2
	yUnitsHalf := g.yUnits / 2

	g.xTopRight = xUnitsHalf
	g.xTopLeft = -xUnitsHalf
	g.yTopUp = yUnitsHalf
	g.yTopDown = -yUnitsHalf
	g.left = Punto{g.xTopLeft, 0}
	g.right = Punto{g.xTopRight, 0}
	g.Up = Punto{0, g.yTopUp}
	g.Down = Punto{0, g.yTopDown}
	g.Origin = Punto{0, 0}
}
