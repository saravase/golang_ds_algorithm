package main

import (
	"fmt"
	"golang_ds_algorithm/searching"
	"golang_ds_algorithm/sorting"
)

func main() {
	fmt.Println(sorting.BubbleSort([]int{12, 123, 21, 3, 213, 12, 3, 21, 312, 31}))
	fmt.Println(sorting.SelectionSort([]int{12, 123, 21, 3, 213, 12, 3, 21, 312, 31}))
	fmt.Println(sorting.InsertionSort([]int{12, 123, 21, 3, 213, 12, 3, 21, 312, 31}))

	fmt.Println(searching.LinearSearch(sorting.BubbleSort([]int{12, 123, 21, 3, 213, 12, 3, 21, 312, 31}), 312))
	fmt.Println(searching.BinarySearch(sorting.BubbleSort([]int{12, 123, 21, 3, 213, 12, 3, 21, 312, 31}), 31))
}
