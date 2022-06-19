package main

import "fmt"

func swap(a, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

func bubbleSort(arr []float64) []float64 {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

func findMinimum(arr []float64) int { // return index of the minimum in array
	size := len(arr)
	min := arr[0]
	index := 0
	for i := 0; i < size; i++ {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}
	return index
}

func selectionSort(arr []float64) []float64 {
	size := len(arr)
	for i := 0; i < size; i++ {
		j := findMinimum(arr[i:])
		swap(&arr[i], &arr[j+i])
	}
	return arr
}

func main() {
	arr := []float64{3, 4, 5, 12, 3, 4}
	arr = selectionSort(arr)
	fmt.Println(arr)
}
