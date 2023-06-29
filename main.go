package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	print("Bem vinda(o) ao sistema de busca de filmes!")

	GetTop250Movies()
}

func GetTop250Movies() {
	res, err := http.Get("https://imdb-api.com/en/API/Top250Movies/{apiKey}")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	fmt.Print("Body: ", string(resBody))
}
