package main

func merge(A []int, m int, B []int, n int) {
	num := m + n - 1
	i := m - 1
	j := n - 1

	for i >= 0 && j >= 0 {
		if A[i] >= B[j] {
			A[num] = A[i]
			num--
			i--
		} else {
			A[num] = B[j]
			num--
			j--
		}
	}

	// B数组依然有剩余
	if j >= 0 {
		for k := 0; k <= j; k++ {
			A[k] = B[k]
		}
	}
}

func main() {
}
