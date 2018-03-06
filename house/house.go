package house

import "fmt"

type candidate struct {
	Name  string
	Seats int
}

// House .
type House struct {
	Candidates []candidate
}

// New .
func New(count int) *House {

	var house = new(House)

	house.Candidates = make([]candidate, count)

	return house
}

// Print simply shows the contents of the house.
func (house House) Print() {

	for _, cand := range house.Candidates {
		fmt.Printf("%10s: %5d\n", cand.Name, cand.Seats)
	}
}
