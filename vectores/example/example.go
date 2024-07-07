package main

import (
	"math"
	"math/rand"
	. "vectores/R2Plot"
	. "vectores/vectores"
)

func main() {

	ResX := 20
	ResY := 20

	NewPlotter(createRandomVectors(), Config{
		PlayerEnabled:     true,
		AnimationEnabled:  true,
		AxesEnabled:       true,
		AxesDetailEnabled: true,
		TobsEnabled:       true,
		ResolutionX:       ResX,
		ResolutionY:       ResY,
	})

}

func createRandomVectors() []Vector {
	vectors := make([]Vector, 1)

	for i := range vectors {
		// Genera un ángulo aleatorio entre 0 y 2π
		angle := 2 * math.Pi * rand.Float64()

		// Genera una longitud aleatoria entre 0.5 y 4
		length := 0.5 + 3.5*rand.Float64()

		// Genera un punto de origen aleatorio
		maxSeparation := 19.0
		originX := rand.Float64()*maxSeparation - maxSeparation/2 // Ajusta el rango según sea necesario
		originY := rand.Float64()*maxSeparation - maxSeparation/2 // Ajusta el rango según sea necesario

		// Crea un vector con el origen, la longitud y el ángulo generados
		vectors[i] = V(Punto{originX, originY}, Punto{originX + length*math.Cos(angle), originY + length*math.Sin(angle)})
	}

	return vectors
}
