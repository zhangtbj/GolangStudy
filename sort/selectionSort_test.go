package sort

import (
	"fmt"
	"testing"
)

func selectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		// loop is for j := 1 + i; j < length; j++
		for j := 1 + i; j < length; j++ {
			// comparision is arr[i] > arr[j]
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 7, 9, 3, 6, 1, 2, 5, 3, 8}
	selectionSort(arr)
	fmt.Println(arr)
}
