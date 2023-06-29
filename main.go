package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Items struct {
	Items []Filme `json:"items"`
}

type Filme struct {
	Titulo         string `json:"title"`
	TituloCompleto string `json:"fullTitle"`
	Ano            string `json:"year"`
	Nota           string `json:"imDbRating"`
	Imagem         string `json:"image"`
	Elenco         string `json:"crew"`
	Ranking        string `json:"rank"`
}

func main() {
	print("Bem vinda(o) ao sistema de busca de filmes!")

	GetTop250Movies()
}

func GetTop250Movies() []Filme {
	res, err := http.Get("https://imdb-api.com/en/API/Top250Movies/{apiKey}")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	filmes := Items{}

	err = json.Unmarshal(resBody, &filmes)

	return filmes.Items
}
