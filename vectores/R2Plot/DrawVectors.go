package R2Plotter

// drawTobs dibuja los vectores en la pantalla, le suma el punto de origen (g.Origin) a cada vector
func (g *Graficadora) drawTobs() {

	for _, vector := range g.Tobs {
		g.drawArrowSimple(vector)
	}

}
