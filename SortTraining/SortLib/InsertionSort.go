package SortLib

//插入排序
//每轮将数列未排序部分的第一个数字插入已排好的部分。
func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		temp := slice[i]
		j := i
		for j > 0 && temp < slice[j-1] {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = temp
	}
}
