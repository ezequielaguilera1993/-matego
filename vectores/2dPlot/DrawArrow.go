package R2Plotter

import (
	v "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	. "vectores/vectores"
)

func (g *Graficadora) drawArrowPoint(vector Vector, bidirectional bool, clr color.Color) {

	startX := vector.Inicio.GetX()
	startY := vector.Inicio.GetY()
	endX := vector.Fin.GetX()
	endY := vector.Fin.GetY()

	// Calcula la dirección del vector
	dx := startX - endX
	dy := endY - startY

	angle := math.Atan2(dy, dx)

	// Define la longitud de las alas de la flecha
	arrowLength := float64(10)
	arrowAngle := math.Pi / 6 // 30 grados

	// Calcula los puntos de las alas
	x1 := vector.Fin.GetX() - arrowLength*math.Cos(angle-arrowAngle)
	y1 := endY - arrowLength*math.Sin(angle-arrowAngle)

	x2 := endX - arrowLength*math.Cos(angle+arrowAngle)
	y2 := endY - arrowLength*math.Sin(angle+arrowAngle)

	// Dibuja las alas de la flecha
	v.StrokeLine(g.screen, endX, endY, x1, y1, 2, clr, true)
	v.StrokeLine(g.screen, endX, endY, x2, y2, 2, clr, true)

	if bidirectional {
		// Calcula los puntos de las alas
		x1 = startX + arrowLength*math.Cos(angle-arrowAngle)
		y1 = startY + arrowLength*math.Sin(angle-arrowAngle)

		x2 = startX + arrowLength*math.Cos(angle+arrowAngle)
		y2 = startY + arrowLength*math.Sin(angle+arrowAngle)

		// Dibuja las alas de la flecha
		v.StrokeLine(g.screen, startX, startY, x1, y1, 2, clr, true)
		v.StrokeLine(g.screen, startX, startY, x2, y2, 2, clr, true)
	}
}
func (g *Graficadora) drawArrowLowLevel(start, end ParOrdenado, bidirectional bool, clr color.Color) {
	v.StrokeLine(g.screen, start.X, start.Y, end.X, end.Y, 2, clr, true)
	g.drawArrowPoint(start, end, bidirectional, clr)
}

func (g *Graficadora) drawArrowBidirectional(start, end ParOrdenado, clr color.Color) {
	g.drawArrowLowLevel(start, end, true, clr)
}

func (g *Graficadora) drawArrow(start, end ParOrdenado, clr color.Color) {
	g.drawArrowLowLevel(start, end, false, clr)
}
