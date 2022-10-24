package crossproduct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Job struct {
	name string
	id   int
}

func TestCorrectCrossProduct(t *testing.T) {
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

	// Construct Cross Product from these slices
	intsAndBools := []any{sliceInts, sliceBools}
	cp := NewCrossProduct(intsAndBools)
	actual := cp.Values()
	expected := []any{
		[]any{1, true},
		[]any{1, false},
		[]any{8, true},
		[]any{8, false},
	}
	assert.Equal(t, expected, actual)

	// Construct Cross Product from these other slices
	intsAndJobs := []any{sliceInts, sliceJobs}
	anotherCP := NewCrossProduct(intsAndJobs)
	actual = anotherCP.Values()
	expected = []any{
		[]any{1, job1},
		[]any{1, job2},
		[]any{8, job1},
		[]any{8, job2},
	}
	assert.Equal(t, expected, actual)

	// Construct another Cross Product
	boolsAndStrings := []any{sliceBools, sliceStrings}
	yetAnotherCP := NewCrossProduct(boolsAndStrings)
	actual = yetAnotherCP.Values()
	expected = []any{
		[]any{true, "testing"},
		[]any{true, "more"},
		[]any{true, "things"},
		[]any{false, "testing"},
		[]any{false, "more"},
		[]any{false, "things"},
	}
	assert.Equal(t, expected, actual)

	// Construct another one!
	jobsAndStringsAndInts := []any{sliceJobs, sliceStrings, sliceInts}
	oneMoreCP := NewCrossProduct(jobsAndStringsAndInts)
	actual = oneMoreCP.Values()
	expected = []any{
		[]any{job1, "testing", 1},
		[]any{job1, "testing", 8},
		[]any{job1, "more", 1},
		[]any{job1, "more", 8},
		[]any{job1, "things", 1},
		[]any{job1, "things", 8},
		[]any{job2, "testing", 1},
		[]any{job2, "testing", 8},
		[]any{job2, "more", 1},
		[]any{job2, "more", 8},
		[]any{job2, "things", 1},
		[]any{job2, "things", 8},
	}
	assert.Equal(t, expected, actual)
}
