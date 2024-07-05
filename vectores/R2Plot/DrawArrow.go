package R2Plotter

import (
	v "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	. "vectores/vectores"
)

func (g *Graficadora) drawArrowPoint(vector Vector, bidirectional bool, clr color.Color) {

	startX := vector.GetStartX()
	startY := vector.GetStartY()
	endX := vector.GetEndX()
	endY := vector.GetEndY()

	// Calcula la dirección del vector
	dx := endX - startX
	dy := endY - startY

	angle := math.Atan2(dy, dx)

	// Define la longitud de las alas de la flecha
	arrowLength := float64(10)
	arrowAngle := math.Pi / 6 // 30 grados

	// Calcula los puntos de las alas
	x1 := endX - arrowLength*math.Cos(angle-arrowAngle)
	y1 := endY - arrowLength*math.Sin(angle-arrowAngle)

	x2 := endX - arrowLength*math.Cos(angle+arrowAngle)
	y2 := endY - arrowLength*math.Sin(angle+arrowAngle)

	// Dibuja las alas de la flecha
	v.StrokeLine(g.screen, float32(endX), float32(endY), float32(x1), float32(y1), 2, clr, true)
	v.StrokeLine(g.screen, float32(endX), float32(endY), float32(x2), float32(y2), 2, clr, true)

	if bidirectional {
		// Calcula los puntos de las alas
		x1 = startX + arrowLength*math.Cos(angle-arrowAngle)
		y1 = startY + arrowLength*math.Sin(angle-arrowAngle)

		x2 = startX + arrowLength*math.Cos(angle+arrowAngle)
		y2 = startY + arrowLength*math.Sin(angle+arrowAngle)

		// Dibuja las alas de la flecha
		v.StrokeLine(g.screen, float32(startX), float32(startY), float32(x1), float32(y1), 2, clr, true)
		v.StrokeLine(g.screen, float32(startX), float32(startY), float32(x2), float32(y2), 2, clr, true)
	}
}
func (g *Graficadora) drawArrowLowLevel(vector Vector, bidirectional bool, clr color.Color) {

	v.StrokeLine(g.screen, float32(vector.GetStartX()), float32(vector.GetStartY()), float32(vector.GetEndX()), float32(vector.GetEndY()), 2, clr, true)
	g.drawArrowPoint(vector, bidirectional, clr)
}

func (g *Graficadora) drawArrowBidirectional(vector Vector, clr color.Color) {
	g.drawArrowLowLevel(vector, true, clr)

}

func (g *Graficadora) drawArrow(vector Vector, clr color.Color) {
	g.drawArrowLowLevel(vector, false, clr)
}
