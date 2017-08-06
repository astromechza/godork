# package `example`

```golang
import "github.com/AstromechZA/godork/example"
```

Package example is an example of various documentation abilities and edge cases.

This is paragraph text in the middle of things.

	here is some indented stuff

blurp.


## Examples

### `Example_of_stuff`

Example_of_stuff shows an example of various random things.

```golang
package main

import (
	"fmt"
	"github.com/AstromechZA/godork/example"
)

func main() {
	fmt.Println(example.T0, example.T1, example.T2)
}

```


## Constants

Numbered test constants

```golang
const (
	T0	int	= iota
	T1
	T2
)
```

Something is a test string used in various scenarios.

```golang
const Something = "word"
```


## Variables

Some other variables.

```golang
var (
	Hi	= "Hello"
	Byte	= "Good bye"
)
```

DefaultName is a mutable variable used to store a string used in various default scenarios when an override is
not provided.

```golang
var DefaultName = "godork"
```


## Functions

### `func CheckSync(s realsync.Locker)`

CheckSync locks and unlocks a lock



### `func GimmeBuffer() *bytes.Buffer`

GimmeBuffer gives you a bufferino



### `func TopLevelFunk(derp int, names ...string) string`

TopLevelFunk takes the args and does a thing with them!


#### `ExampleTopLevelFunk`


```golang
package main

import (
	"github.com/AstromechZA/godork/example"
)

func main() {
	example.TopLevelFunk(1, "hi")
}

```


