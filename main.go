package main

import (
	"fmt"

	"github.com/worapong778/cinema/movie"
	"github.com/worapong778/cinema/ticket"
)

func main() {
	movieName := movie.FindName("M002")
	fmt.Println(movieName)
	movie.Review(movieName, 8.5)
	ticket.BuyTicket(movieName)
}
