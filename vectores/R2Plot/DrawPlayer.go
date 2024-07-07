package R2Plotter

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Graficadora) drawPlayer() {
	g.drawArrow(g.PlayerVector)

	xc, yc := g.PlayerVector.GetComponents()
	g.drawArrow(PlotVector{xc, Blue})
	g.drawArrow(PlotVector{yc, Green})

	x1 := g.PlayerVector.GetStartX()
	y1 := g.PlayerVector.GetStartY()
	x2 := g.PlayerVector.GetEndX()
	y2 := g.PlayerVector.GetEndY()

	x1f :=
		g.x(x1)
	y1f :=
		g.y(y1)
	x2f :=
		g.x(x2)
	y2f :=
		g.y(y2)

	ebitenutil.DebugPrintAt(g.screen, fmt.Sprintf("x1: %f , y1: %f ; x2: %f , y2: %f", x1f, y1f, x2f, y2f), int(x1), int(y1))

}
