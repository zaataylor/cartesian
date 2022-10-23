package crossproduct

import (
	"fmt"
	"testing"
)

func TestCorrectCrossProduct(t *testing.T) {
	sliceA := []int{1, 8}
	sliceB := []bool{true, false, false, true}
	sliceC := []int{15, 17, 18}
	sliceD := []string{"testing", "more", "things"}
	sliceE := []float64{1.7, 1.77, 1.776, 1.7769}

	slices := []interface{}{sliceA, sliceB, sliceC, sliceD, sliceE}

	cp := NewCrossProduct(slices)
	fmt.Printf("Value of cross products themselves is:\n%s", cp)

	cp.printIndicesOnly = true
	fmt.Printf("Value of cross product indices is: \n%s", cp)
}
