package main

import (
	"math/rand"
	. "vectores/R2Plot"
	. "vectores/vectores"
)

func main() {

	//v1 := V(Punto{10, 10})
	//v2 := V(Punto{20, 20})

	NewPlotter(nil, true, true)

}

func createRandomVectors() []Vector {
	//10 vectores random con un for. Usa rand, quiero vectores entre 10 y 100 numeros

	vectors := make([]Vector, 60)
	maxn := 20.0
	minn := -20.0
	for i := range vectors {
		vectors[i] = Vector{
			Inicio: Punto{rand.Float64()*(maxn-minn) + minn, rand.Float64()*(maxn-minn) + minn},
			Fin:    Punto{rand.Float64()*(maxn-minn) + minn, rand.Float64()*(maxn-minn) + minn},
		}

	}

	return vectors
}
