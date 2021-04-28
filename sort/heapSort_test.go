package sort

import (
	"fmt"
	"testing"
)

// 堆排序
func heapSort(tree []int, n int) {
	buildHeap(tree, n)
	// 重复：第一个结点和最后一个结点交换位置，然后重新调整堆排序，i--去掉最后一个结点
	for i := n - 1; i >= 0; i-- {
		swap(tree, i, 0)
		heapify(tree, i, 0)
	}
}

// 创建堆
func buildHeap(tree []int, n int) {
	lastNode := n - 1            // 最后一个结点
	parent := (lastNode - 1) / 2 // 最后一个结点的父节点
	// 往最后一个父节点开始，不断往前创建结点
	for i := parent; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

// 调整堆,第i个结点
func heapify(tree []int, n, i int) {
	if i >= n {
		return
	}
	// c1左子结点，c2右左子结点
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	if max != i {
		swap(tree, max, i)
		heapify(tree, n, max)
	}
}

// 交换数组位置
func swap(tree []int, i, j int) {
	tree[i], tree[j] = tree[j], tree[i]
}

func TestHeapSort(t *testing.T) {
	arr := []int{4, 7, 9, 3, 6, 1, 2, 5, 3, 8}
	heapSort(arr, len(arr))
	fmt.Println(arr)
}
