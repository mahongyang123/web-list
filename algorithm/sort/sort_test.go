package _sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	array := []int{1, 9, 5, 3, 11, 3, 6}

	fmt.Println("sort before:", array)

	Sort(array, len(array))

	fmt.Println("sort after:", array)
}

func TestBubbleSort(t *testing.T) {
	array := []int{1, 9, 5, 3, 11, 3, 6}
	//var array []int
	//	fmt.Println("33", unsafe.Pointer(&array))
	fmt.Println("bubble sort before:", array)

	BubbleSort(array, len(array))

	fmt.Println("bubble sort after:", array)
	// fmt.Println("44", unsafe.Pointer(&array))
}

func TestInsertionSort(t *testing.T) {
	array := []int{1, 9, 5, 3, 11, 3, 6}

	fmt.Println("insertionSort sort before:", array)

	InsertionSort(array, len(array))

	fmt.Println("insertionSort sort after:", array)
}

func TestSelectionSort(t *testing.T) {
	array := []int{1, 9, 5, 3, 11, 3, 6}

	fmt.Println("selectionSort sort before:", array)

	SelectionSort(array, len(array))

	fmt.Println("selectionSort sort after:", array)
}
