package cartesian

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Cartesian Product
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
	// compute Cartesian and values upfront
	c.computeCartesianProduct()
	return &c
}

func (c *CartesianProduct) computeCartesianProduct() {
	for c.count < c.max {
		if c.count == 0 {
			c.count += 1
			tmp := make([]int, c.length)
			copy(tmp, c.data)
			c.indices = append(c.indices, tmp)
			c.values = append(c.values, c.getValues(tmp))
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
			tmp := make([]int, c.length)
			copy(tmp, c.data)
			c.indices = append(c.indices, tmp)
			c.values = append(c.values, c.getValues(c.indices[c.count]))
			c.count += 1
			c.j = 1
		}
	}
}

func (c *CartesianProduct) getValues(indices []int) []any {
	res := make([]any, 0, len(indices))
	for i, sl := range c.slices {
		slice := reflect.ValueOf(sl)
		valueInSlice := slice.Index(indices[i]).Interface()
		res = append(res, valueInSlice)
	}
	return res
}

func (c *CartesianProduct) Values() []any {
	return c.values
}

func (c *CartesianProduct) Indices() [][]int {
	return c.indices
}

func (c *CartesianProduct) String() string {
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

// Iterators
type CartesianProductIterator struct {
	iteratorCount int
	max           int
	cartesian     *CartesianProduct
}

func (c *CartesianProduct) Iterator() *CartesianProductIterator {
	return &CartesianProductIterator{
		iteratorCount: 0,
		max:           c.max,
		cartesian:     c,
	}
}

func (cpi *CartesianProductIterator) ResetIterator() {
	cpi.iteratorCount = 0
}

func (cpi *CartesianProductIterator) NextIndices() []int {
	if !cpi.HasNext() {
		return nil
	}
	indices := cpi.cartesian.indices[cpi.iteratorCount]
	cpi.iteratorCount += 1
	return indices
}

func (cpi *CartesianProductIterator) Next() []any {
	if !cpi.HasNext() {
		return nil
	}
	indices := cpi.NextIndices()
	return cpi.cartesian.getValues(indices)
}

func (cpi *CartesianProductIterator) HasNext() bool {
	return cpi.iteratorCount < cpi.max
}
