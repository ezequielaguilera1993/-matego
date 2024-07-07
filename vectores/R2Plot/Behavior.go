package R2Plotter

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func (g *Graficadora) playerBehavior() {
	speed := 2.0
	rotationSpeed := 0.1 // velocidad de rotación

	// Calcula el ángulo y la magnitud del vector
	dx := g.PlayerVector.GetEndX() - g.PlayerVector.GetStartX()
	dy := g.PlayerVector.GetEndY() - g.PlayerVector.GetStartY()
	angle := math.Atan2(dy, dx)
	magnitude := math.Sqrt(dx*dx + dy*dy)

	// Calcula el vector de dirección (vector normalizado)
	dirX := dx / magnitude
	dirY := dy / magnitude

	// Actualiza la posición del vector controlado
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		// Mueve el vector a lo largo de su dirección
		g.PlayerVector.Inicio[0] += dirX * speed
		g.PlayerVector.Inicio[1] += dirY * speed
		g.PlayerVector.Fin[0] += dirX * speed
		g.PlayerVector.Fin[1] += dirY * speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		// Mueve el vector a lo largo de su dirección en sentido contrario
		g.PlayerVector.Inicio[0] -= dirX * speed
		g.PlayerVector.Inicio[1] -= dirY * speed
		g.PlayerVector.Fin[0] -= dirX * speed
		g.PlayerVector.Fin[1] -= dirY * speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// Gira el vector a la izquierda
		angle -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		// Gira el vector a la derecha
		angle += rotationSpeed
	}

	// Calcula las nuevas coordenadas del punto final del vector
	g.PlayerVector.Fin[0] = g.PlayerVector.GetStartX() + magnitude*math.Cos(angle)
	g.PlayerVector.Fin[1] = g.PlayerVector.GetStartY() + magnitude*math.Sin(angle)

}

var count = 0
var frameChange = 40

func (g *Graficadora) vectorsBehavior() {
	speed := 0.1
	angleChange := 0.0 // máximo cambio de ángulo aleatorio
	fmt.Println(count)
	for i := range g.Vectores {
		// Calcula el ángulo y la magnitud del vector
		dx := g.Vectores[i].GetEndX() - g.Vectores[i].GetStartX()
		dy := g.Vectores[i].GetEndY() - g.Vectores[i].GetStartY()
		angle := math.Atan2(dy, dx)
		magnitude := math.Sqrt(dx*dx + dy*dy)

		// Si va de 0 a 100 cambia para un lado, si va de 100 a 200 para el otroy asi sucesivamente, cada 100 cambia de lado
		if count < frameChange {
			angle += angleChange
		} else {
			angle -= angleChange
		}

		// Calcula las nuevas coordenadas del punto final del vector
		g.Vectores[i].Fin[0] = g.Vectores[i].GetStartX() + magnitude*math.Cos(angle)
		g.Vectores[i].Fin[1] = g.Vectores[i].GetStartY() + magnitude*math.Sin(angle)

		// Actualiza la posición del vector
		g.Vectores[i].Inicio[0] += dx * speed
		g.Vectores[i].Inicio[1] += dy * speed
		g.Vectores[i].Fin[0] += dx * speed
		g.Vectores[i].Fin[1] += dy * speed

		// Verifica si el vector ha salido de los límites de la pantalla
		// y ajusta las coordenadas para que aparezca en el lado opuesto
		if g.Vectores[i].Inicio[0] < 0 {
			g.Vectores[i].Inicio[0] += g.width
			g.Vectores[i].Fin[0] += g.width
		} else if g.Vectores[i].Inicio[0] > g.width {
			g.Vectores[i].Inicio[0] -= g.width
			g.Vectores[i].Fin[0] -= g.width
		}

		if g.Vectores[i].Inicio[1] < 0 {
			g.Vectores[i].Inicio[1] += g.height
			g.Vectores[i].Fin[1] += g.height
		} else if g.Vectores[i].Inicio[1] > g.height {
			g.Vectores[i].Inicio[1] -= g.height
			g.Vectores[i].Fin[1] -= g.height
		}
	}
	if count == frameChange*2 {
		count = 0
	} else {
		count++
	}
}
