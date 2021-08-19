package sortlib

func MergeSort(slice []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	MergeSort(slice, 0, mid)
	MergeSort(slice, mid+1, end)
	merge(slice, start, mid, end)
}

func merge(slice []int, start, mid, end int) {
	var tmp []int
	i := start
	j := mid + 1

	for i <= mid && j <= end {
		if slice[i] < slice[j] {
			tmp = append(tmp, slice[i])
			i++
		} else {
			tmp = append(tmp, slice[j])
			j++
		}
	}

	if i > mid {
		tmp = append(tmp, slice[j:end+1]...)
	} else {
		tmp = append(tmp, slice[i:mid+1]...)
	}

	copy(slice[start:end+1], tmp)
}
