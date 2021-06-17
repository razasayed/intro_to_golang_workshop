package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://official-joke-api.appspot.com/jokes/programming/random"

type Joke struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type Jokes []Joke

func main() {
	// Make the http request
	res, err := http.Get(BASE_URL)
	if err != nil {
		fmt.Println("Unable to fetch data")
		return
	}

	// Read the response body
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to parse response body")
		return
	}

	// Parse response body as json
	var jokes Jokes
	if err := json.Unmarshal(bytes, &jokes); err != nil {
		fmt.Println("Error parsing json", err)
		return
	}

	// Print result
	fmt.Println(jokes[0].Setup, jokes[0].Punchline)
}
