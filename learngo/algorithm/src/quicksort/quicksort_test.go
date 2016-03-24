package quicksort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{1,2,3,4,5,12,15}, QuickSort([]int{15,12,5,3,4,2,1}))
}

func TestPartition(t *testing.T) {
	slice := []int{1, 7, 3}
	assert.Equal(t, 0, partition(slice,0, 2))
	assert.Equal(t, []int{1, 7, 3}, slice)
}

func TestQuickSort2(t *testing.T) {
	var v []int = [] int {123,321,1,2,3,10,7}
	var r []int = [] int {1,2,3,7,10,123,321}

	QuickSort(v)

	for i := 0; i <len(v); i++ {
		if v[i] != v[i] {
			t.Error("QuickSort() Error",v,"ExpectedL",r)
		}
	}
}

func TestQuickSort3(t *testing.T) {
	var v []int = []int{0}
	var r []int = []int{0}
	QuickSort(v)
	if v[0] != r[0] {
		t.Error("QuickSort() Error.", v, "Expected:", r)
	}
}
