package types

// custom types
type Integer = int   // alias to type int
type Location string // new type based on type int - this can have methods

// type method
func (from Location) DistanceTo(to Location) Location {
	// mocked distance
	return "120km"
}
