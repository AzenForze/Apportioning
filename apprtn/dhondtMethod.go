package apprtn

import (
	"Apportioning/house"
	"Apportioning/ielection"
)

// DHondtMethod .
type DHondtMethod struct {
	base *DivisorApportioner
}

// NewDHondt .
func NewDHondt() *DHondtMethod {

	return &DHondtMethod{&DivisorApportioner{
		func(votes int, seats int) float64 {

			return float64(votes) / float64(seats+1)
		}}}
}

// Apportion .
func (dhondt *DHondtMethod) Apportion(e ielection.IElection, n int) *house.House {

	return dhondt.base.Apportion(e, n)
}
