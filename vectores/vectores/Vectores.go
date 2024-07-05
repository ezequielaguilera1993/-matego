package vectores

import (
	"fmt"
	"math"
)

// Punto representa un punto en el plano
type Punto []float64

// ParOrdenado representa un par ordenado en el plano
type ParOrdenado struct {
	Inicio, Fin Punto
}

// Vector representa un vector en el plano. Idéntico a un par ordenado.
type Vector ParOrdenado

// V crea un vector a partir de dos puntos
func V(parOrdenado ParOrdenado) Vector {
	return Vector(parOrdenado)
}

// X calcula el producto escalar de dos vectores
func (v Vector) X(vector Vector) float64 {
	return v.GetStartX()*vector.GetStartX() + v.GetStartY()*vector.GetEndY()
}

// Add suma dos vectores
func (v Vector) Add(vector Vector) Vector {
	return Vector{
		Inicio: Punto{v.GetStartX() + vector.GetStartX(), v.GetStartY() + vector.GetStartY()},
		Fin:    Punto{v.GetEndX() + vector.GetEndX(), v.GetEndY() + vector.GetEndY()},
	}
}

// Sub resta dos vectores
func (v Vector) Sub(vector Vector) Vector {
	return Vector{
		Inicio: Punto{v.GetStartX() - vector.GetStartX(), v.GetStartY() - vector.GetStartY()},
		Fin:    Punto{v.GetEndX() - vector.GetEndX(), v.GetEndY() - vector.GetEndY()},
	}
}

// Magnitude calcula la magnitud de un vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X(v))
}

// AngleBetweenVectors calcula el ángulo entre dos vectores
func (v Vector) AngleBetweenVectors(vector Vector) float64 {
	return math.Acos(v.X(vector) / (v.Magnitude() * vector.Magnitude()))
}

// IsOrthogonal determina si dos vectores son ortogonales
func (v Vector) IsOrthogonal(vector Vector) bool {
	return v.X(vector) == 0
}

// Print imprime un vector
func (v Vector) Print() {
	fmt.Println(v)
}

// GetStartX devuelve la coordenada x del primer punto un vector
func (v Vector) GetStartX() float64 {
	return v.Inicio.GetX()
}

// GetStartY devuelve la coordenada y del primer punto un vector
func (v Vector) GetStartY() float64 {
	return v.Inicio.GetY()
}

// GetEndX devuelve la coordenada x del segundo punto un vector
func (v Vector) GetEndX() float64 {
	return v.Fin.GetX()
}

// GetEndY devuelve la coordenada y del segundo punto un vector
func (v Vector) GetEndY() float64 {
	return v.Fin.GetY()
}

// GetX devuelve la coordenada x de un vector
func (p Punto) GetX() float64 {
	return p[0]
}

// GetY devuelve la coordenada y de un vector
func (p Punto) GetY() float64 {
	return p[1]
}

// GetZ devuelve la coordenada z de un vector
func (p Punto) GetZ() float64 {
	return p[2]
}
