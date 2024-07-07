package R2Plotter

import (
	"math"
)

var count = 0
var frameChange = 40

func (g *Graficadora) vectorsBehavior() {
	speed := 0.1
	angleChange := 0.0 // máximo cambio de ángulo aleatorio
	for i := range g.Tobs {
		// Calcula el ángulo yRelative la magnitud del vector
		dx := g.Tobs[i].GetEndX() - g.Tobs[i].GetStartX()
		dy := g.Tobs[i].GetEndY() - g.Tobs[i].GetStartY()
		angle := math.Atan2(dy, dx)
		magnitude := math.Sqrt(dx*dx + dy*dy)

		// Si va de 0 a 100 cambia para un lado, si va de 100 a 200 para el otroy asi sucesivamente, cada 100 cambia de lado
		if count < frameChange {
			angle += angleChange
		} else {
			angle -= angleChange
		}

		// Si un vector sale por un lado de la pantalla, rebota en el otro
		if g.Tobs[i].GetEndX() < -g.xUnits/2 || g.Tobs[i].GetEndX() > g.xUnits/2 {
			c := g.Tobs[i].Inicio[0]
			g.Tobs[i].Inicio[0] = g.Tobs[i].Fin[0]
			g.Tobs[i].Inicio[1] = c
		}

		// Calcula las nuevas coordenadas del punto final del vector
		g.Tobs[i].Fin[0] = g.Tobs[i].GetStartX() + magnitude*math.Cos(angle)
		g.Tobs[i].Fin[1] = g.Tobs[i].GetStartY() + magnitude*math.Sin(angle)

		// Actualiza la posición del vector
		g.Tobs[i].Inicio[0] += dx * speed
		g.Tobs[i].Inicio[1] += dy * speed
		g.Tobs[i].Fin[0] += dx * speed
		g.Tobs[i].Fin[1] += dy * speed

		if count == frameChange*2 {
			count = 0
		} else {
			count++
		}
	}
}
