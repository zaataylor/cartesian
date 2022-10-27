# Description

`cartesian` is a Go package that makes it easy to compute and return the [Cartesian product](https://en.wikipedia.org/wiki/Cartesian_product) of an arbitrary number of arbitrarily typed slices.

# Requires
- Go 1.18 or higher

# Example

Consider the following code:

```golang
package main

import (
	"fmt"

	"github.com/zaataylor/cartesian/cartesian"
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

	slices := []any{sliceA, sliceB, sliceJ}
	// Construct Cartesian product of these slices
	cp := cartesian.NewCartesianProduct(slices)
	fmt.Printf("Cartesian product of slices:\n%s", cp)
}
```

Running this code returns:

```
Cartesian product of slices:

[
  [1, true, {name:test job id:1}], 
  [1, true, {name:another test job id:2}], 
  [1, false, {name:test job id:1}], 
  [1, false, {name:another test job id:2}], 
  [8, true, {name:test job id:1}], 
  [8, true, {name:another test job id:2}], 
  [8, false, {name:test job id:1}], 
  [8, false, {name:another test job id:2}], 
]
 ```

# Using `cartesian`
The sections below illustrate how to use `cartesian` effectively.

## Creating the `CartesianProduct`
Put the slices you want to compute the cartesian of inside of an `[]any`-type slice. Then, provide this slice as input to `NewCartesianProduct()`. For example:
```golang
sliceA := []int{4, 5, 8}
sliceB := []bool{true, false}
input := []any{sliceA, sliceB}

cp := cartesian.NewCartesianProduct(input)
```

## Iterating over `CartesianProduct` Elements
First, use `Iterator()` to create a `CartesianProductIterator`. Then, use the iterator's `HasNext()` method as an indicator of when to continue iterating, and `Next()` to return the iterands themselves. For example:
```golang
// Construct the Cartesian product
sliceA := []int{4, 5, 8}
sliceB := []bool{true, false}
input := []any{sliceA, sliceB}
cp := cartesian.NewCartesianProduct(input)

// Construct the Cartesian Product iterator
cpi := cp.Iterator()


// Iterate over its elements and print them out
for cpi.HasNext() {
    element := cpi.Next()
    fmt.Printf("Element is: %v\n", element)
}

// Reset the iterator, if you'd like...
cpi.ResetIterator()

// ...or, create a new iterator
newCpi := cp.Iterator()
```

## Computing the full Cartesian product
When you first call `NewCartesianProduct()`, it computes the full Cartesian Product as part of its initialization process. You can then use `Values()` to print or iterate over the elements of the product. Example:
```golang
// Construct the Cartesian product
sliceA := []int{4, 5, 8}
sliceB := []bool{true, false}
input := []any{sliceA, sliceB}

cp := cartesian.NewCartesianProduct(input)

// Iterate over its elements and print them out
for _, v := range cp.Values() {
    fmt.Printf("Element is: %#v\n", v)
	// Assign values to individual variables
	firstValue, secondValue := v[0], v[1]
	fmt.Printf("First item: %v; Second item: %v", firstValue, secondValue)
}
```
