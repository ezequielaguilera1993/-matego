package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	v "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (g *Graficadora) xAbsoluto(x float64) float64 {
	return g.width/2 + x*g.xUnitLength
}

func (g *Graficadora) yAbsoluto(y float64) float64 {
	return g.height/2 - y*g.yUnitLength
}

func (g *Graficadora) drawLineLowLevel(x1, y1, x2, y2 float64, width int, color color.Color, antiAliasEnabled bool) {
	v.StrokeLine(g.screen, float32(g.xAbsoluto(x1)), float32(g.yAbsoluto(y1)), float32(g.xAbsoluto(x2)), float32(g.yAbsoluto(y2)), float32(width), color, antiAliasEnabled)
}

func (g *Graficadora) debugPrintRelativeAtLowLevel(x0, x1 float64, message string) {
	x0Absolute := g.xAbsoluto(x0)
	y0Absolute := g.yAbsoluto(x1)
	ebitenutil.DebugPrintAt(g.screen, message, int(x0Absolute), int(y0Absolute))
}
