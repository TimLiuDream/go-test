package sortlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	slice  = []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	result = []int{1, 1, 2, 2, 3, 4, 5, 5, 8, 8}
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(slice)
	assert.Equal(t, result, slice)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(slice)
	assert.Equal(t, result, slice)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(slice)
	assert.Equal(t, result, slice)
}

func TestMergeSort(t *testing.T) {
	MergeSort(slice, 0, len(slice)-1)
	assert.Equal(t, result, slice)
}

func TestShellSort(t *testing.T) {
	ShellSort(slice, len(slice))
	assert.Equal(t, result, slice)
}

func TestQuickSort(t *testing.T) {
	QuickSort(slice, 0, len(slice)-1)
	assert.Equal(t, result, slice)
}
