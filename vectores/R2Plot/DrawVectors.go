package R2Plotter

import (
	"image/color"
	"math/rand"
	. "vectores/vectores"
)

// drawVectors dibuja los vectores en la pantalla, le suma el punto de origen (g.Zero) a cada vector
func (g *Graficadora) drawVectors() {
	for _, vector := range g.Vectores {
		//sumarle g.zero a cada punto

		v := Vector{
			Inicio: Punto{vector.GetStartX() + g.Zero.GetX(), -vector.GetStartY() + g.Zero.GetY()},
			Fin:    Punto{vector.GetEndX() + g.Zero.GetX(), -vector.GetEndY() + g.Zero.GetY()},
		}

		g.drawArrowLowLevel(v, false, randomColor())

	}
}

func randomColor() color.Color {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
