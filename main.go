package main

import "fmt"

func main() {
	fmt.Println("hello world")

	game := Game{
		profit:    0,
		pollution: 0}

	fmt.Printf("Game: %d", game.profit)
}

type Game struct {
	profit    int
	pollution int
}
