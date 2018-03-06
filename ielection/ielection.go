package ielection

// IElection .
type IElection interface {
	Size() int
	Name(int) string
	Support(int) int
}
