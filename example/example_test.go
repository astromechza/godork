package example_test

import (
	"fmt"

	"github.com/AstromechZA/godork/example"
)

// Example_of_stuff shows an example of various random things.
func Example_of_stuff() {
	fmt.Println(example.T0, example.T1, example.T2)
	// Output: 0 1 2
}

func ExampleTopLevelFunk() {
	example.TopLevelFunk(1, "hi")
	// Output: 1 ["hi"]
}

func ExampleSomeType_PointerReceiver() {
	t := example.SomeType{}
	t.PointerReceiver(1)
	// Output: 1
}
