package R2Plotter

import (
	"fmt"
)

func (g *Graficadora) drawPlayer() {
	g.drawArrowSimple(g.PlayerVector)

	xc, yc := g.PlayerVector.GetComponents()
	g.drawArrowSimple(plotVector{xc, Blue})
	g.drawArrowSimple(plotVector{yc, Green})

	x0 := g.PlayerVector.GetStartX()
	y0 := g.PlayerVector.GetStartY()
	x1 := g.PlayerVector.GetEndX()
	y1 := g.PlayerVector.GetEndY()

	g.debugPrintRelativeAtLowLevel(x0, y0, fmt.Sprintf("(%.1f,%.1f) ; (%.1f,%.1f)", x0, y0, x1, y1))

}
