package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JokeRes struct {
	err      string
	category string
	Type     string
	Joke     string
}

func main() {
	fmt.Println("Joke API")
	// joke := getJoke()
	getJoke()
}

func getJoke() {
	const url string = "https://v2.jokeapi.dev/joke/Programming,Dark?type=single"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result JokeRes
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Joke);
	fmt.Println(result.category);
}
