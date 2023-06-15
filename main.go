package main

import (
	"fmt"

	"github.com/51wick51/cinema/movie"
	"github.com/51wick51/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("title1")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
