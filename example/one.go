package example

import "fmt"

// Numbered test constants
const (
	T0 int = iota
	T1
	T2
)

// Something is a test string used in various scenarios.
const Something = "word"

// DefaultName is a mutable variable used to store a string used in various default scenarios when an override is
// not provided.
var DefaultName = "godork"

// Some other variables.
var (
	Hi   = "Hello"
	Byte = "Good bye"
)

// SomeType is a struct holding the state of the example struct type.
// A bunch of content here
//    indent
// stuff.
type SomeType struct {
	// External is the verbleburble
	External int
	internal int
}

// ValueReceiver is an example of a method with a value rcvr
func (s SomeType) ValueReceiver(x int) {
	fmt.Println(x)
}

// PointerReceiver is an example of a method with a pointer rcvr
func (s *SomeType) PointerReceiver(y int) {
	fmt.Println(y)
}

// TopLevelFunk takes the args and does a thing with them!
func TopLevelFunk(derp int, names ...string) string {
	return fmt.Sprintf("%v, %v", derp, names)
}
