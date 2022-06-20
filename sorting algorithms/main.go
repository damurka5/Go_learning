package main

import (
	"fmt"
)

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

func insertionSort(arr []float64) []float64 {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
	return arr
}

func merge(arr []float64, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]float64, n1)
	M := make([]float64, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[p+i]
	}
	for i := 0; i < n2; i++ {
		M[i] = arr[q+i+1]
	}

	i := 0
	j := 0
	k := p

	for i < n1 && j < n2 {
		if L[i] <= M[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = M[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = M[j]
		j++
		k++
	}
}

func mergeSort(arr []float64, l, r int) {
	if l < r {
		m := l + (r-l)/2

		mergeSort(arr, l, m)
		mergeSort(arr, l+1, r)

		merge(arr, l, m, r)
	}
}

func main() {
	arr := []float64{3, 4, 5, 12, 3, 4}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
