package R2Plotter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

var speed = 0.2

func (g *Graficadora) playerBehavior() {
	rotationSpeed := 0.05 // velocidad de rotación

	// Calcula el ángulo yRelative la magnitud del vector
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

	g.PlayerVector.Fin[0] = g.PlayerVector.GetStartX() + magnitude*math.Cos(angle)
	g.PlayerVector.Fin[1] = g.PlayerVector.GetStartY() + magnitude*math.Sin(angle)

	// Calcula las nuevas coordenadas del punto final del vector
	//controlador de velocidad
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.PlayerVector.Print()
		n := g.PlayerVector.X(2)
		g.PlayerVector.Vector = n // Aumenta el tamaño en un 10%
		n.Print()

	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.PlayerVector.Vector = g.PlayerVector.X(0.9) // Disminuye el tamaño en un 10%
	}
}
