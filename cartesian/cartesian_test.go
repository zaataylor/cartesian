package cartesian

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Job struct {
	name string
	id   int
}

func TestCorrectCartesianProduct(t *testing.T) {
	sliceInts := []int{1, 8}
	sliceBools := []bool{true, false}
	sliceStrings := []string{"testing", "more", "things"}
	job1 := Job{
		name: "test job",
		id:   1,
	}
	job2 := Job{
		name: "another test job",
		id:   2,
	}
	sliceJobs := []Job{
		job1,
		job2,
	}

	// Construct Cartesian product from these slices
	intsAndBools := []any{sliceInts, sliceBools}
	cp := NewCartesianProduct(intsAndBools)
	actual := cp.Values()
	expected := [][]any{
		{1, true},
		{1, false},
		{8, true},
		{8, false},
	}
	assert.Equal(t, expected, actual)

	// Construct Cartesian product from these other slices
	intsAndJobs := []any{sliceInts, sliceJobs}
	anotherCP := NewCartesianProduct(intsAndJobs)
	actual = anotherCP.Values()
	expected = [][]any{
		{1, job1},
		{1, job2},
		{8, job1},
		{8, job2},
	}
	assert.Equal(t, expected, actual)

	// Construct another Cartesian product
	boolsAndStrings := []any{sliceBools, sliceStrings}
	yetAnotherCP := NewCartesianProduct(boolsAndStrings)
	actual = yetAnotherCP.Values()
	expected = [][]any{
		{true, "testing"},
		{true, "more"},
		{true, "things"},
		{false, "testing"},
		{false, "more"},
		{false, "things"},
	}
	assert.Equal(t, expected, actual)

	// Construct another one!
	jobsAndStringsAndInts := []any{sliceJobs, sliceStrings, sliceInts}
	oneMoreCP := NewCartesianProduct(jobsAndStringsAndInts)
	actual = oneMoreCP.Values()
	expected = [][]any{
		{job1, "testing", 1},
		{job1, "testing", 8},
		{job1, "more", 1},
		{job1, "more", 8},
		{job1, "things", 1},
		{job1, "things", 8},
		{job2, "testing", 1},
		{job2, "testing", 8},
		{job2, "more", 1},
		{job2, "more", 8},
		{job2, "things", 1},
		{job2, "things", 8},
	}
	assert.Equal(t, expected, actual)
}

func TestCartesianProductIterator(t *testing.T) {
	sliceInts := []int{1, 8}
	sliceBools := []bool{true, false}
	slices := []any{sliceInts, sliceBools}
	cp := NewCartesianProduct(slices)
	cpi := cp.Iterator()
	// Cartesian product should have four elements, so
	// after 4 iterations, there shouldn't be anything
	// else left
	expectedNumIterations := 4
	i := 0
	for i < expectedNumIterations {
		v := cpi.Next()
		fmt.Printf("Value is: %#v\n", v)
		a, b := v[0], v[1]
		fmt.Printf("a is %v, b is %v\n", a, b)
		i += 1
	}
	assert.False(t, cpi.HasNext())

	cpi.ResetIterator()
	assert.True(t, cpi.HasNext())

	for _, v := range cp.Values() {
		log.Printf("Element is: %v\n", v)
	}

	log.Printf("Values are: %v", cp)
	cp.printIndicesOnly = true
	log.Printf("Indices are: %v", cp)
}
