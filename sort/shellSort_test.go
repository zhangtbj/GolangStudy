package sort

import (
	"fmt"
	"testing"
)

func ShellSort(arr []int) {
	length := len(arr)
	increment := 1
	for increment < length/3 { //寻找合适的间隔h
		increment = 3*increment + 1
	}
	for increment >= 1 {
		//将数组变为间隔h个元素有序
		for i := increment; i < length; i++ {
			//间隔h插入排序
			for j := i; j >= increment; j -= increment {
				if arr[j] < arr[j-increment] {
					arr[j], arr[j-increment] = arr[j-increment], arr[j]
				}
			}
		}
		increment /= 3
	}
}

func TestShellSort(t *testing.T) {
	arr := []int{4, 7, 9, 3, 6, 1, 2, 5, 3, 8}
	ShellSort(arr)
	fmt.Println(arr)
}
