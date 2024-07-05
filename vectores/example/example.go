package main

import (
	"math/rand"
	. "vectores/R2Plot"
	. "vectores/vectores"
)

func main() {

	v1 := Vector{Inicio: Punto{0, 0}, Fin: Punto{200, 200}}
	v2 := Vector{Inicio: Punto{120, 90}, Fin: Punto{130, 100}}
	//
	NewPlotter([]Vector{v1, v2})

}

func createRandomVectors() []Vector {
	//10 vectores random con un for. Usa rand, quiero vectores entre 10 y 100 numeros
	vectors := make([]Vector, 10)
	maxn := 100.0
	minn := -100.0
	for i := range vectors {
		vectors[i] = Vector{
			Inicio: Punto{rand.Float64()*(maxn-minn) + minn, rand.Float64()*(maxn-minn) + minn},
			Fin:    Punto{rand.Float64()*(maxn-minn) + minn, rand.Float64()*(maxn-minn) + minn},
		}

	}

	return vectors
}
