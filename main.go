package main

import (
	"Apportioning/apprtn"
	"Apportioning/election"
	"fmt"
)

func main() {
	var dh = apprtn.NewDHondt()
	var sainte = apprtn.NewSainte()

	var college = new(election.Election)

	college.Emplace("Yellow", 47000)
	college.Emplace("White", 16000)
	college.Emplace("Red", 15900)
	college.Emplace("Green", 12000)
	college.Emplace("Blue", 6000)
	college.Emplace("Pink", 3100)

	var seats = 10

	college.ApportionSeats(dh, seats).Print()
	fmt.Println()
	college.ApportionSeats(sainte, seats).Print()
}
