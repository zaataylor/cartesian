package crossproduct

import (
	"fmt"
	"testing"
)

type Job struct {
	name string
	id   int
}

func TestCorrectCrossProduct(t *testing.T) {
	sliceA := []int{1, 8}
	sliceB := []bool{true, false}
	sliceC := []int{15, 17, 18}
	sliceD := []string{"testing", "more", "things"}
	sliceE := []float64{1.7, 1.77, 1.776, 1.7769}
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

	slices := []interface{}{sliceA, sliceB, sliceC, sliceD, sliceE}
	// Construct Cross Product from these slices
	cp := NewCrossProduct(slices)
	fmt.Printf("Value of cross products themselves is:\n%s", cp)

	cp.printIndicesOnly = true
	fmt.Printf("Value of cross product indices is: \n%s", cp)

	// Construct Cross Product from these other slices
	otherSlices := []interface{}{sliceA, sliceJ}
	anotherCP := NewCrossProduct(otherSlices)
	fmt.Printf("Cross Product of otherSlices: %s", anotherCP)

	// Construct another Cross Product from yet other slices
	moreSlices := []interface{}{sliceA, sliceB}
	yetAnotherCP := NewCrossProduct(moreSlices)
	fmt.Printf("Cross Product of moreSlices: %s", yetAnotherCP)
}
