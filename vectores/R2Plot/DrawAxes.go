package R2Plotter

import (
	. "vectores/vectores"
)

const (
	lineLength     = 0.3
	halfLineLength = lineLength / 2
	lineWidth      = 1
)

func (g *Graficadora) drawAxes() {
	g.drawArrowBidirectional(plotVector{Vector: Vector{
		Inicio: g.left,
		Fin:    g.right,
	},
		Color: Blue,
	})
	g.drawArrowBidirectional(plotVector{Vector: Vector{
		Inicio: g.Up,
		Fin:    g.Down,
	},
		Color: Green,
	})
}

// drawAxesDetails dibuja las leyendas de los ejes. Las líneas que cortan los ejes yRelative los números que indican las coordenadas
func (g *Graficadora) drawAxesDetails() {

	for x := 1 + (-g.xUnits / 2); x < g.xUnits/2; x++ { //evita los extremos
		//a la primera iter no le suma ninguno yRelative despues se va dezplazna
		g.drawDetailLineX(x)
	}

	for y := 1 + (-g.yUnits / 2); y < g.yUnits/2; y++ { //evita los extremos
		//a la primera iter no le suma ninguno yRelative despues se va dezplaznad
		g.drawDetailLineY(y)
	}
}

func (g *Graficadora) drawDetailLineX(x float64) {
	g.drawLineLowLevel(x, -halfLineLength, x, halfLineLength, lineWidth, White, true)
}

func (g *Graficadora) drawDetailLineY(y float64) {
	g.drawLineLowLevel(-halfLineLength, y, halfLineLength, y, lineWidth, White, true)
}
