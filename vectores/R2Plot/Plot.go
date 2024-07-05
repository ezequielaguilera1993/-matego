package R2Plotter

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
	"math/rand"
	. "vectores/vectores"
)

type PlotVector struct {
	Vector
	Color color.Color
}

// Graficadora es la estructura principal que contiene los vectores a graficar
type Graficadora struct {
	Vectores     []PlotVector
	PlayerVector PlotVector
	screen       *ebiten.Image
	Coordinates
	width, height float64
}

type Coordinates struct {
	xZero     float64 // la mitad del width de la pantalla
	yZero     float64 // la mitad del height de la pantalla
	xTopRight float64 // la parte más a la derecha de la pantalla
	xTopLeft  float64 // la parte más a la izquierda de la pantalla
	yTopUp    float64 // la parte más alta de la pantalla
	yTopDown  float64 // la parte más baja de la pantalla
	left      Punto
	right     Punto
	Up        Punto
	Down      Punto
	Zero      Punto
}

func (g *Graficadora) Update() error {
	g.vectorsBehavior()
	g.player()
	return nil
}

func (g *Graficadora) player() {
	speed := 10.0
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

func (g *Graficadora) Draw(screen *ebiten.Image) {
	g.screen = screen
	g.drawAxes()
	g.drawVectors()
	g.PlayerVector.Color = randomColor()

}

func (g *Graficadora) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Graficadora) drawAxes() {
	g.drawArrowBidirectional(PlotVector{Vector: Vector{
		Inicio: g.left,
		Fin:    g.right,
	},
		Color: Blue,
	})
	g.drawArrowBidirectional(PlotVector{Vector: Vector{
		Inicio: g.Up,
		Fin:    g.Down,
	},
		Color: Green,
	})
}

func NewPlotter(vectores []Vector) {
	width, height, scale := 640.0, 480.0, 1.5
	g := &Graficadora{
		width:  width,
		height: height,
	}
	g.setCoordinates()

	plotVectors := make([]PlotVector, len(vectores))
	for i, v := range vectores {
		plotVectors[i] = PlotVector{Vector: Vector{
			Inicio: Punto{v.GetStartX() + g.Zero.GetX(), -v.GetStartY() + g.Zero.GetY()},
			Fin:    Punto{v.GetEndX() + g.Zero.GetX(), -v.GetEndY() + g.Zero.GetY()},
		}, Color: randomColor()}
	}

	g.Vectores = plotVectors

	g.PlayerVector = PlotVector{Vector: Vector{
		Inicio: Punto{10 + g.Zero.GetX(), -30 + g.Zero.GetY()},
		Fin:    Punto{30 + g.Zero.GetX(), -80 + g.Zero.GetY()},
	}, Color: Green}

	ebiten.SetWindowSize(int(width*scale), int(height*scale))
	ebiten.SetWindowTitle("Graficadora de Vectores")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func randomColor() color.Color {
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func (g *Graficadora) setCoordinates() {
	width := g.width
	height := g.height
	g.xZero, g.yZero = float64(width)/2, float64(height)/2
	g.xTopRight = float64(width)
	g.xTopLeft = 0
	g.yTopUp = 0
	g.yTopDown = float64(height)
	g.left = Punto{g.xTopLeft, g.yZero}
	g.right = Punto{g.xTopRight, g.yZero}
	g.Up = Punto{g.xZero, g.yTopUp}
	g.Down = Punto{g.xZero, g.yTopDown}
	g.Zero = Punto{g.xZero, g.yZero}
}
