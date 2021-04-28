package sort

import (
	"fmt"
	"testing"
)

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		// loop is for j := 1; j < length - i; j++
		for j := 1; j < length - i; j++ {
			// comparision is: arr[j] < arr[j - 1]
			if arr[j] < arr[j - 1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 7, 9, 3, 6, 1, 2, 5, 3, 8}
	bubbleSort(arr)
	fmt.Println(arr)
}
