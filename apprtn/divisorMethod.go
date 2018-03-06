package apprtn

import (
	"Apportioning/house"
	"Apportioning/ielection"
)

// DivisorApportioner .
type DivisorApportioner struct {
	divisor func(int, int) float64
}

// NewDivisorApportioner creates a new DivisorApportioner using a given divisor.
func NewDivisorApportioner(divisor func(int, int) float64) *DivisorApportioner {

	return &DivisorApportioner{divisor}
}

// Apportion apportions seats to candidates based on their support from an election, returning a House.
func (apprtn DivisorApportioner) Apportion(election ielection.IElection, numSeats int) *house.House {

	var numCandidates = election.Size()

	var house = house.New(numCandidates)
	var remainders = make([]float64, numCandidates)

	for i := 0; i < numCandidates; i++ {
		var houseCand = &house.Candidates[i]

		houseCand.Name = election.Name(i)
		houseCand.Seats = 0

		remainders[i] = float64(election.Support(i))
	}

	for i := 0; i < numSeats; i++ {

		var largest = findLargestIndex(remainders)

		var largestCandidate = &house.Candidates[largest]

		largestCandidate.Seats++

		remainders[largest] = apprtn.divisor(election.Support(largest), largestCandidate.Seats)
	}

	return house
}

// findLargestIndex finds the index of a slice with the largest float64 value, and returns it.
func findLargestIndex(list []float64) int {
	var index = 0
	var largestValue = list[index]

	for i, val := range list {
		if val > largestValue {
			largestValue = val
			index = i
		}
	}

	return index
}
