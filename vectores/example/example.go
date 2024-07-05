package main

import (
	. "vectores/2dPlot"
	. "vectores/vectores"
)

func main() {

	v1 := V(ParOrdenado{Inicio: Punto{0, 0}, Fin: Punto{100, 100}})
	v2 := V(ParOrdenado{Inicio: Punto{5, 10}, Fin: Punto{105, 105}})

	NewPlotter([]Vector{v1, v2})

}
