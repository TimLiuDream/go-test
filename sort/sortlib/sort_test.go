package sortlib

import (
	"reflect"
	"testing"
)

var (
	slice  = []int{1, 5, 8, 3, 2, 1, 5, 8, 2, 4}
	result = []int{1, 1, 2, 2, 3, 4, 5, 5, 8, 8}
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(slice)
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("BubbleSort failed. Got %v, expected %v", slice, result)
	}
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(slice)
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("InsertionSort failed. Got %v, expected %v", slice, result)
	}
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(slice)
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("SelectionSort failed. Got %v, expected %v", slice, result)
	}
}

func TestMergeSort(t *testing.T) {
	MergeSort(slice, 0, len(slice)-1)
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("MergeSort failed. Got %v, expected %v", slice, result)
	}
}

func TestShellSort(t *testing.T) {
	ShellSort(slice, len(slice))
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("ShellSort failed. Got %v, expected %v", slice, result)
	}
}

func TestQuickSort(t *testing.T) {
	QuickSort(slice, 0, len(slice)-1)
	if !reflect.DeepEqual(result, slice) {
		t.Errorf("QuickSort failed. Got %v, expected %v", slice, result)
	}
}
