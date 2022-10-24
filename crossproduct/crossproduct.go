package crossproduct

import (
	"fmt"
	"reflect"
	"strings"
)

type CrossProduct struct {
	printIndicesOnly bool
	count            int
	max              int
	length           int
	j                int
	slices           []any
	data             []int
	moduli           []int
	results          [][]int
}

func NewCrossProduct(inputSlices []any) *CrossProduct {
	c := CrossProduct{
		printIndicesOnly: false,
		count:            0,
		j:                1,
		max:              1,
		slices:           inputSlices,
		length:           len(inputSlices),
		data:             make([]int, len(inputSlices)),
		moduli:           make([]int, len(inputSlices)),
	}
	for i, sl := range inputSlices {
		slice := reflect.ValueOf(sl)
		c.moduli[i] = slice.Len()
		c.max *= slice.Len()
	}
	return &c
}

func (c *CrossProduct) NextIndices() []int {
	if c.count == 0 {
		c.count += 1
		tmp := make([]int, c.length)
		copy(tmp, c.data)
		c.results = append(c.results, tmp)
		return c.data
	}
	if c.count < c.max {
		// increment by "1", then take modulus
		v := (c.data[c.length-c.j] + 1) % c.moduli[c.length-c.j]
		c.data[c.length-c.j] = v
		// carry the "1" if v is 0
		if v == 0 {
			for v == 0 && c.length-c.j > 0 {
				// shift down 1 (i.e. one to the left)
				c.j += 1
				// increment by "1", then take modulus
				v = (c.data[c.length-c.j] + 1) % c.moduli[c.length-c.j]
				c.data[c.length-c.j] = v
			}
		}
		c.count += 1
		c.j = 1
		tmp := make([]int, c.length)
		copy(tmp, c.data)
		c.results = append(c.results, tmp)
		return c.data
	}
	return nil
}

func (c *CrossProduct) Next() []any {
	indices := c.NextIndices()
	if indices == nil {
		return nil
	}
	res := make([]any, 0, len(indices))
	for i, sl := range c.slices {
		slice := reflect.ValueOf(sl)
		v := slice.Index(indices[i]).Interface()
		res = append(res, v)
	}
	return res
}

func (c *CrossProduct) Values() []any {
	values := []any{}
	v := c.Next()
	for v != nil {
		values = append(values, v)
		v = c.Next()
	}
	return values
}

func (c *CrossProduct) Indices() [][]int {
	cpResult := c.NextIndices()
	for cpResult != nil {
		cpResult = c.NextIndices()
	}
	return c.results
}

func (c *CrossProduct) String() string {
	// compute remaining results
	c.Indices()

	// add results to the string
	s := "[\n"
	if c.printIndicesOnly {
		// print indices into each set in cross product
		for i, r := range c.results {
			if i == len(c.results)-1 {
				s += fmt.Sprintf("  %v", r)
			} else {
				s += fmt.Sprintf("  %v, \n", r)
			}
		}
	} else {
		// print the actual cross product values. More expensive
		res := []string{}
		for _, r := range c.results {
			for k, sl := range c.slices {
				slice := reflect.ValueOf(sl)
				res = append(res, fmt.Sprintf("%v", slice.Index(r[k])))
			}
			s += "  (" + strings.Join(res, ", ") + "), \n"
			res = []string{}
		}

	}
	s += "]\n"
	return s
}
