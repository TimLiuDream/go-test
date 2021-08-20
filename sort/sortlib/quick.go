package sortlib

// QuickSort 快速排序
// 快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
// 1. 在数组中，选择一个元素作为“基准”(pivot),一般选择第一个元素作为基准元素。设置两个游标i和j，初始时i指向数组首元素，j指向尾元素。
// 2. 从数组最右侧向前扫描，遇到小于基准值的元素停止扫描，将两者交换，然后从数组左侧开始扫描，遇到大于基准值的元素停止扫描，同样将两者交换。
// 3. i==j时分区完成，否则转2。
func QuickSort(slice []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(slice, start, end)
	QuickSort(slice, start, mid-1)
	QuickSort(slice, mid+1, end)
}

func partition(slice []int, low int, high int) int {
	i, j := low+1, high
	for {
		for slice[i] < slice[low] {
			i++
			if i == high {
				break
			}
		}
		for slice[low] < slice[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		exch(slice, i, j)
	}
	exch(slice, low, j)
	return j
}

func exch(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
