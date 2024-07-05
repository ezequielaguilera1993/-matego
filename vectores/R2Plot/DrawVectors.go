package R2Plotter

import . "vectores/vectores"

// drawVectors dibuja los vectores en la pantalla, le suma el punto de origen (g.Zero) a cada vector
func (g *Graficadora) drawVectors() {
	for _, vector := range g.Vectores {
		//sumarle g.zero a cada punto

		v := Vector{
			Inicio: Punto{vector.GetStartX() + g.Zero.GetX(), vector.GetStartY() + g.Zero.GetY()},
			Fin:    Punto{vector.GetEndX() + g.Zero.GetX(), vector.GetEndY() + g.Zero.GetY()},
		}

		g.drawArrowLowLevel(v, false, White)

	}
}
