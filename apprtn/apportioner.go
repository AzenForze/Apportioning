package apprtn

import (
	"Apportioning/house"
	"Apportioning/ielection"
)

// Apportioner .
type Apportioner interface {
	Apportion(ielection.IElection, int) *house.House
}
