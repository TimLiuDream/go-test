package sortlib

func HeapSort(slice []int) {
	length := len(slice)
	// 建堆
	buildMaxHeap(slice)
	// 排序
	for i := length - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, 0, i)
	}
}

func buildMaxHeap(slice []int) {
	length := len(slice)
	for i := length / 2; i >= 0; i-- {
		heapify(slice, i, length)
	}
}

// 调整堆
func heapify(slice []int, i, length int) {
	// 从i节点开始调整
	// 为什么要 * 2
	// 因为堆是完全二叉树，所以左右子节点的位置是固定的
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	// 找出最大的节点
	if left < length && slice[left] > slice[largest] {
		largest = left
	}
	if right < length && slice[right] > slice[largest] {
		largest = right
	}
	// 如果最大的节点不是父节点，交换
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapify(slice, largest, length)
	}
}
