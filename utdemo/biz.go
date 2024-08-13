package utdemo

func QuickSort(input []int) {
	recrusive(input, 0, len(input)-1)
}

func recrusive(input []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(input, l, r)
	recrusive(input, l, pivot-1)
	recrusive(input, pivot+1, r)
}

func partition(input []int, l, r int) int {
	pivot := input[r]

	i := l
	for j := l; j < r; j++ {
		if input[j] < pivot {
			if i != j {
				input[i], input[j] = input[j], input[i]
			}
			i++
		}
	}
	input[i], input[r] = input[r], input[i]
	return i
}
