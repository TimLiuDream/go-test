package sortlib

// BubbleSort 冒泡排序
// 每轮都得到数列中的最大值，同时将其置于最后，然后对剩余部分进行排序。
func BubbleSort(slice []int) {
	var isFinish bool

	for !isFinish {
		isFinish = true //当已经全部排序完，isFinish不会再改变为false，那就跳出循环
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				temp := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = temp
				isFinish = false //有元素变动就是未完成，继续
			}
		}
	}
}

// BubbleSort 冒泡排序
// 每轮都得到数列中的最大值，同时将其置于最后，然后对剩余部分进行排序。
func BubbleSort1(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i; j < len(slice); j++ {
			if slice[i] > slice[j] {
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}
	}
}
