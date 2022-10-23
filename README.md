# Description

`crossproduct` is a Go package that makes it easy to compute and return the [cross products](https://en.wikipedia.org/wiki/Cross_product) of an arbitrary number of arbitrarily typed slices.

# Example

Consider the following code:

```golang
package main

import (
	"fmt"

	"github.com/zaataylor/crossproduct/crossproduct"
)

type Job struct {
	name string
	id   int
}

func main() {
	sliceA := []int{1, 8}
	sliceB := []bool{true, false}
	sliceJ := []Job{
		{
			name: "test job",
			id:   1,
		},
		{
			name: "another test job",
			id:   2,
		},
	}

	slices := []interface{}{sliceA, sliceB, sliceJ}
	// Construct cross product of these slices
	cp := crossproduct.NewCrossProduct(slices)
	fmt.Printf("Cross Product of slices:\n%s", cp)
}
```

Running this code returns:

```
Cross Product of slices:
[ 
  (1, true, {test job 1}), 
  (1, true, {another test job 2}), 
  (1, false, {test job 1}), 
  (1, false, {another test job 2}), 
  (8, true, {test job 1}), 
  (8, true, {another test job 2}), 
  (8, false, {test job 1}), 
  (8, false, {another test job 2}), 
 ]
 ```

