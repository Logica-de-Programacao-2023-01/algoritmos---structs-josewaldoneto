package main

import "fmt"

// Crie uma struct chamada Triângulo com os campos "base" e "altura".
// Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).

type Triangulo struct {
	base, altura float64
}

func area(t Triangulo) float64 {
	return t.base * t.altura / 2
}

func main() {
	var tri Triangulo
	fmt.Println("Base: ")
	fmt.Scan(&tri.base)
	fmt.Println("Altura: ")
	fmt.Scan(&tri.altura)
	a := area(tri)
	fmt.Printf("Área: %.2f\n", a)
}
