package R2Plotter

import (
	"image/color"
	"math"
)

func (g *Graficadora) drawArrowPoint(vector plotVector, bidirectional bool, scale float64, clr color.Color) {

	startX := vector.GetStartX()
	startY := vector.GetStartY()
	endX := vector.GetEndX()
	endY := vector.GetEndY()

	// Calcula la dirección del vector
	dx := endX - startX
	dy := endY - startY

	angle := math.Atan2(dy, dx)

	// Define la longitud de las alas de la flecha
	arrowLength := vector.Magnitude() * scale
	arrowAngle := math.Pi / 6 // 30 grados

	// Calcula los puntos de las alas
	x1 := endX - arrowLength*math.Cos(angle-arrowAngle)
	y1 := endY - arrowLength*math.Sin(angle-arrowAngle)

	x2 := endX - arrowLength*math.Cos(angle+arrowAngle)
	y2 := endY - arrowLength*math.Sin(angle+arrowAngle)

	// Dibuja las alas de la flecha
	g.drawLineLowLevel(endX, endY, x1, y1, 2, clr, true)
	g.drawLineLowLevel(endX, endY, x2, y2, 2, clr, true)

	if bidirectional {
		// Calcula los puntos de las alas
		x1 = startX + arrowLength*math.Cos(angle-arrowAngle)
		y1 = startY + arrowLength*math.Sin(angle-arrowAngle)

		x2 = startX + arrowLength*math.Cos(angle+arrowAngle)
		y2 = startY + arrowLength*math.Sin(angle+arrowAngle)

		// Dibuja las alas de la flecha
		g.drawLineLowLevel(startX, startY, x1, y1, 2, clr, true)
		g.drawLineLowLevel(startX, startY, x2, y2, 2, clr, true)
	}
}

func (g *Graficadora) drawArrow(vector plotVector, bidirectional bool, scale float64, width float64) {
	g.drawLineLowLevel(vector.GetStartX(), vector.GetStartY(), vector.GetEndX(), vector.GetEndY(), 2, vector.Color, true)
	g.drawArrowPoint(vector, bidirectional, scale, vector.Color)
}

func (g *Graficadora) drawArrowBidirectional(vector plotVector) {
	g.drawArrow(vector, true, 0.025, 1)

}

func (g *Graficadora) drawArrowSimple(vector plotVector) {
	g.drawArrow(vector, false, 0.30, 2)
}
