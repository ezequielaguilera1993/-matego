package vectores

import (
	"fmt"
	"math"
)

// Punto representa un punto en el plano
type Punto []float64

// Vector representa un vector en el plano. Idéntico a un par ordenado.
type Vector struct {
	Inicio, Fin Punto
}

// V crea un nuevo vector con puntos de inicio y fin.
// Si no se proporcionan puntos, se utilizan valores predeterminados.
func V(puntos ...Punto) Vector {
	switch len(puntos) {
	case 1:
		return Vector{Inicio: Punto{0, 0}, Fin: puntos[0]}
	case 2:
		return Vector{Inicio: puntos[0], Fin: puntos[1]}
	default:
		panic("solo se permiten uno o dos puntos para crear un vector")
	}
}

// X calcula el producto escalar de dos vectores
func (v Vector) X(vector Vector) float64 {
	return v.GetStartX()*vector.GetStartX() + v.GetEndY()*vector.GetEndY()
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
	return math.Sqrt(math.Pow(v.GetEndX()-v.GetStartX(), 2) + math.Pow(v.GetEndY()-v.GetStartY(), 2))
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

// GetStartN devuelve la coordenada n del primer punto un vector
func (v Vector) GetStartN(n int) float64 {
	return v.Inicio.GetN(n)
}

// GetEndN devuelve la coordenada n del segundo punto un vector
func (v Vector) GetEndN(n int) float64 {
	return v.Fin.GetN(n)
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

// GetN devuelve la coordenada n de un vector
func (p Punto) GetN(n int) float64 {
	return p[n]
}
