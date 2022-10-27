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

	// Construct the Cartesian product

	input := []any{sliceA, sliceB}

	cp2 := cartesian.NewCartesianProduct(input)

	// Iterate over its elements and print them out
	for _, v := range cp2.Values() {
		fmt.Printf("Element is: %v\n", v)
	}

}
