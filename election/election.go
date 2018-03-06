package election

import (
	"Apportioning/apprtn"
	"Apportioning/house"
)

type candidate struct {
	name    string
	support int
}

// Election .
type Election struct {
	candidate []candidate
}

// Add adds an existing candidate to the election.
func (election *Election) Add(candidate candidate) {
	election.candidate = append(election.candidate, candidate)
}

// Emplace creates a new candidate using the given name and support,
// and adds them to the election
func (election *Election) Emplace(name string, support int) {
	election.Add(candidate{name, support})
}

// Size .
func (election *Election) Size() int {
	return len(election.candidate)
}

// Name .
func (election *Election) Name(index int) string {
	return election.candidate[index].name
}

// Support .
func (election *Election) Support(index int) int {
	return election.candidate[index].support
}

// ApportionSeats apportions any number of seats using a given method
// based on the support of each candidate in the election.
func (election *Election) ApportionSeats(method apprtn.Apportioner, seats int) *house.House {
	return method.Apportion(election, seats)
}
