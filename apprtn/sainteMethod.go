package apprtn

import (
	"Apportioning/house"
	"Apportioning/ielection"
)

// SainteMethod .
type SainteMethod struct {
	base *DivisorApportioner
}

// NewSainte .
func NewSainte() *SainteMethod {

	return &SainteMethod{&DivisorApportioner{
		func(votes int, seats int) float64 {

			return float64(votes) / float64((2*seats)+1)
		}}}
}

// Apportion .
func (sainte *SainteMethod) Apportion(e ielection.IElection, n int) *house.House {

	return sainte.base.Apportion(e, n)
}
