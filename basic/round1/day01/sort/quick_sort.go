package sort

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	midValue := a[mid]
	i := l - 1
	j := r + 1
	for i < j {
		i++
		for a[i] < midValue {
			i++
		}

		j--
		for a[j] > midValue {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	quickSort(a, l, j)
	quickSort(a, j+1, r)
}
