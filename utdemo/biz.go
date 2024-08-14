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

func HeapSort(input []int) {
	// build heap
	buildHeap(input)

	// sort
	for i := len(input) - 1; i > 0; i-- {
		input[i], input[0] = input[0], input[i]
		heapify(input, i-1, 0)
	}
}

func buildHeap(input []int) {
	for i := (len(input) - 1) / 2; i >= 0; i-- {
		heapify(input, len(input)-1, i)
	}
}

// input[0:n]中index=i进行向下堆化
func heapify(input []int, n, i int) {
	for {
		left, right := 2*i+1, 2*i+2
		maxPos := i
		if left <= n && input[left] > input[maxPos] {
			maxPos = left
		}
		if right <= n && input[right] > input[maxPos] {
			maxPos = right
		}
		if i == maxPos {
			return
		}
		input[maxPos], input[i] = input[i], input[maxPos]
		i = maxPos
	}
}

// 请求外部依赖 进行处理逻辑