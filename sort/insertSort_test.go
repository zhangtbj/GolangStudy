package sort

import (
	"fmt"
	"testing"
)

func insertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		// loop is j := i; j > 0; j--
		for j := i; j > 0; j-- {
			// comparision is arr[j] > arr[j - 1]
			if arr[j] > arr[j - 1] {
				break
			}
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
	}
}

func TestInsertSort(t *testing.T) {
	arr := []int{4, 7, 9, 3, 6, 1, 2, 5, 3, 8}
	insertSort(arr)
	fmt.Println(arr)
}
