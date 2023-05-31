package sortlib

// BubbleSort 冒泡排序
// 每轮都得到数列中的最大值，同时将其置于最后，然后对剩余部分进行排序。
func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ { // 循环 n 次
		for j := 0; j < len(slice)-i-1; j++ { // 进行排序
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
