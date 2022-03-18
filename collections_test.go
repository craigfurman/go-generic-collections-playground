package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFunctionsOnCollection(t *testing.T) {
	l := []int{3, 10, 1, 2, 20}
	plus1 := func(i int) int {
		return i + 1
	}
	lessThan10 := func(i int) bool {
		return i < 10
	}

	// This is very ugly
	res := Reduce(
		0,
		Filter(
			Map(
				l,
				plus1,
			),
			lessThan10,
		),
		func(acc, element int) int {
			return acc + element
		},
	)

	assert.Equal(t, 9, res)
}

func TestFunctionalCollectionPipeline(t *testing.T) {
	l := []int{3, 10, 1, 2, 20}
	plus1 := func(i int) int {
		return i + 1
	}
	lessThan10 := func(i int) bool {
		return i < 10
	}
	res := Pipe4(l, PipeMap(plus1), PipeFilter(lessThan10), PipeReduce(0, func(acc, element int) int {
		return acc + element
	}))
	assert.Equal(t, 9, res)
}
