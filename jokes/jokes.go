package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	ID        int
	Type      string
	Setup     string
	Punchline string
}

func RandomJokeOfType(jokeType string) (*Joke, error) {
	resp, err := http.Get(fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%s/random", jokeType))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := []*Joke{{}}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no jokes to return")
	}
	return data[0], nil
}
