package R2Plotter

import (
	v "github.com/hajimehoshi/ebiten/v2/vector"
	. "vectores/vectores"
)

const (
	lineLength     = 10
	halfLineLength = lineLength / 2
	width          = 1
)

// drawAxesDetails dibuja las leyendas de los ejes. Las líneas que cortan los ejes y los números que indican las coordenadas
func (g *Graficadora) drawAxesDetails() {

	// Dibuja las líneas de los ejes
	g.drawAxes()

	for x := 1 + (-g.xUnits / 2); x < g.xUnits/2; x++ { //evita los extremos
		//a la primera iter no le suma ninguno y despues se va dezplaznad
		g.drawDetailLineX(x * g.xUnitLength)
	}

	for y := 1 + (-g.yUnits / 2); y < g.yUnits/2; y++ { //evita los extremos
		//a la primera iter no le suma ninguno y despues se va dezplaznad
		g.drawDetailLineY(y * g.yUnitLength)
	}
}

func (g *Graficadora) drawDetailLineX(x float64) {
	v.StrokeLine(g.screen, float32(g.xZero+x), float32(g.yZero)-halfLineLength, float32(g.xZero+x), float32(g.yZero)+halfLineLength, width, White, true)
}

func (g *Graficadora) drawDetailLineY(y float64) {
	v.StrokeLine(g.screen, float32(g.xZero)-halfLineLength, float32(g.yZero+y), float32(g.xZero)+halfLineLength, float32(g.yZero+y), width, White, true)
}

func (g *Graficadora) drawAxes() {
	g.drawArrowBidirectional(PlotVector{Vector: Vector{
		Inicio: g.left,
		Fin:    g.right,
	},
		Color: Blue,
	})
	g.drawArrowBidirectional(PlotVector{Vector: Vector{
		Inicio: g.Up,
		Fin:    g.Down,
	},
		Color: Green,
	})
}
