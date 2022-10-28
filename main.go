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

	// Example 1: Construct Cartesian product of these
	// slices and print all its elements
	slices := []any{sliceA, sliceB, sliceJ}
	cp := cartesian.NewCartesianProduct(slices)
	fmt.Printf("Cartesian product of slices:\n%s", cp)

	// Example 2: Construct the Cartesian product and
	// iterate over elements one by one
	input := []any{sliceA, sliceB}
	cp2 := cartesian.NewCartesianProduct(input)
	for _, v := range cp2.Values() {
		fmt.Printf("Element is: %v\n", v)
	}

	// Example 3: Functions, too (using Indices())
	sliceS := []string{"MSFT", "GOOG", "META"}
	sliceF := make([]func(bool, string) string, 0)
	sliceF = append(sliceF, isStonksFunc)
	sliceF = append(sliceF, isOneFunc)

	anotherInput := []any{sliceF, sliceB, sliceS}
	cp3 := cartesian.NewCartesianProduct(anotherInput)
	for _, indexes := range cp3.Indices() {
		fn, boolArg, stringArg := sliceF[indexes[0]], sliceB[indexes[1]], sliceS[indexes[2]]
		fmt.Printf("Function: %v, boolArg: %v, stringArg: %v --> Result: %v\n", cartesian.GetFunctionName(fn), boolArg, stringArg, fn(boolArg, stringArg))
	}
}

func isStonksFunc(isStonks bool, company string) string {
	if isStonks {
		return fmt.Sprintf("%s is stonks! 👍", company)
	} else {
		return fmt.Sprintf("%s is NOT stonks! 👎", company)
	}
}

func isOneFunc(isOne bool, _ string) string {
	if isOne {
		return "isOne"
	}
	return "isZero"
}
