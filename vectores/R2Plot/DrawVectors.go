package R2Plotter

// drawVectors dibuja los vectores en la pantalla, le suma el punto de origen (g.Zero) a cada vector
func (g *Graficadora) drawVectors() {

	for _, vector := range g.Vectores {
		g.drawArrow(vector)
	}

	g.drawArrow(g.PlayerVector)

}
