package main

import (
	"fmt"
	"math"
)

// Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo (área = pi * raio * raio).
// Dica: use a constante math.Pi para representar o número pi.

type Circulo struct {
	raio float64
}

func areadocirculo(c Circulo) float64 {
	pi := math.Pi
	return pi * c.raio * c.raio
}

func main() {
	var circ Circulo
	fmt.Print("Digite o raio do círculo: ")
	fmt.Scan(&circ.raio)
	a := areadocirculo(circ)
	fmt.Printf("O valor da área é, aproximadamente, %.2f\n", a)
}
