package sort

func mergeSort(a []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(a, l, mid)
	mergeSort(a, mid+1, r)
	tmp := make([]int, r-l+1)
	k := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = a[j]
		k++
		j++
	}
	copy(a[l:r+1], tmp)
}
