package main

// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
// O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado".
// Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.

type Pessoa struct {
	nome     string
	idade    int
	endereco struct {
		rua, numero, cidade, estado string
	}
}
