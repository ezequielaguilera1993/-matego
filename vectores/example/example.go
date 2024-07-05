package main

import (
	. "vectores/R2Plot"
	. "vectores/vectores"
)

func main() {

	v1 := Vector{Inicio: Punto{0, 0}, Fin: Punto{100, 100}}
	v2 := Vector{Inicio: Punto{0, 0}, Fin: Punto{100, -100}}

	v1.Add(v2).Print()

	NewPlotter([]Vector{v1, v2})

}
