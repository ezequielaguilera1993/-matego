package R2Plotter

func (g *Graficadora) drawVectors() {
	for _, vector := range g.Vectores {
		g.drawArrow(vector[0], vector[1], White)
	}
}
