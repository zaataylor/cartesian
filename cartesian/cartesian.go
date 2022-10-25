package cartesian

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type CartesianProduct struct {
	printIndicesOnly bool
	count            int
	max              int
	length           int
	j                int
	slices           []any
	data             []int
	moduli           []int
	indices          [][]int
	values           []any
}

func NewCartesianProduct(inputSlices []any) *CartesianProduct {
	c := CartesianProduct{
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

func (c *CartesianProduct) NextIndices() []int {
	if c.count == 0 {
		c.count += 1
		tmp := make([]int, c.length)
		copy(tmp, c.data)
		c.indices = append(c.indices, tmp)
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
		c.indices = append(c.indices, tmp)
		return c.data
	}
	return nil
}

func (c *CartesianProduct) Next() []any {
	if !c.HasNext() {
		return nil
	}
	indices := c.NextIndices()
	res := make([]any, 0, len(indices))
	for i, sl := range c.slices {
		slice := reflect.ValueOf(sl)
		valueInSlice := slice.Index(indices[i]).Interface()
		res = append(res, valueInSlice)
	}
	c.values = append(c.values, res)
	return res
}

func (c *CartesianProduct) HasNext() bool {
	return c.count < c.max
}

func (c *CartesianProduct) Values() []any {
	if !c.HasNext() {
		return c.values
	}
	for c.HasNext() {
		c.Next()
	}
	return c.values
}

func (c *CartesianProduct) Indices() [][]int {
	if !c.HasNext() {
		return c.indices
	}
	for c.HasNext() {
		c.NextIndices()
	}
	return c.indices
}

func (c *CartesianProduct) String() string {
	// compute any remaining results
	c.Values()

	if c.printIndicesOnly {
		return c.createIndicesString()
	} else {
		return c.createValuesString()
	}
}

func (c *CartesianProduct) createIndicesString() string {
	s := "\n[\n"
	for _, r := range c.indices {
		b, _ := json.Marshal(r)
		s += "  " + strings.ReplaceAll(string(b), ",", ", ") + "\n"
	}
	s += "]\n"
	return s
}

func marshalAndReplace(item any) string {
	b, _ := json.Marshal(item)
	return "  " + strings.ReplaceAll(string(b), ",", ", ") + "\n"
}

func (c *CartesianProduct) createValuesString() string {
	s := "\n[\n"
	res := []string{}
	for _, r := range c.indices {
		for k, sl := range c.slices {
			slice := reflect.ValueOf(sl)
			res = append(res, fmt.Sprintf("%+v", slice.Index(r[k]).Interface()))
		}
		s += "  [" + strings.Join(res, ", ") + "], \n"
		res = []string{}
	}
	s += "]\n"
	return s
}
