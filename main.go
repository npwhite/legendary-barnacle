// package main
//
// import (
// 	"flag"
// 	"fmt"
//
// 	"github.com/npwhite/legendary-barnacle/jokes"
// )
//
// // legal joke types
// const (
// 	GeneralJoke     = "general"
// 	ProgrammingJoke = "programming"
// )
//
// var flag_jokeType string
//
// func init() {
// 	flag.StringVar(&flag_jokeType, "type", GeneralJoke, "The type of joke to get. Must be one of: general, programming")
// 	flag.Parse()
// }
//
// func main() {
// 	joke, err := jokes.RandomJokeOfType(flag_jokeType)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(joke.Setup)
// 	fmt.Println(joke.Punchline)
// }

package main

import (
	"fmt"
	"time"
	"math/rand"
)

var greetings = []string{
	"hello!",
	"hey!",
	"welcome!",
	"hello there!",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := greetings[rand.Intn(len(greetings))]
	fmt.Println(g)
}
