package R2Plotter

import (
	v "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

func (g *Graficadora) drawArrowPoint(vector PlotVector, bidirectional bool, scale float64, clr color.Color) {

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

func (g *Graficadora) drawArrowLowLevel(vector PlotVector, bidirectional bool, scale float64) {
	v.StrokeLine(g.screen, float32(vector.GetStartX()), float32(vector.GetStartY()), float32(vector.GetEndX()), float32(vector.GetEndY()), 2, vector.Color, true)
	g.drawArrowPoint(vector, bidirectional, scale, vector.Color)
}

func (g *Graficadora) drawArrowBidirectional(vector PlotVector) {
	g.drawArrowLowLevel(vector, true, 0.025)

}

func (g *Graficadora) drawArrow(vector PlotVector) {
	g.drawArrowLowLevel(vector, false, 0.30)
}
