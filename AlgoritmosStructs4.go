package main

// Crie uma struct chamada Playlist com os campos "nome" e "músicas".
// O campo "músicas" deve ser um slice de structs, cada uma representando uma música com os campos "título", "artista" e "duração".
// Escreva uma função que receba uma Playlist como parâmetro e imprima o nome da playlist, o tempo total da playlist e as informações de cada música.

type Playlist struct {
	nome    string
	musicas []struct {
		titulo, artista string
		duracao         float64
	}
}
