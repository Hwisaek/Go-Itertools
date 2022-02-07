package itertools

import (
	"sort"
)

func Permutation(arr []int, r int) (result [][]int) {
	sort.Ints(arr)
	check := make([]int, len(arr))
	recursive([]int{}, check, r, &arr, &result)
	return result
}

func recursive(chosen []int, check []int, r int, arr *[]int, result *[][]int) {
	if len(chosen) == r {
		copySlice := make([]int, len(chosen))
		copy(copySlice, chosen)
		*result = append(*result, copySlice)
		return
	}

	for i := 0; i < len(*arr); i++ {
		if check[i] == 0 {
			chosen = append(chosen, (*arr)[i])
			check[i] = 1
			recursive(chosen, check, r, arr, result)
			check[i] = 0
			chosen = chosen[:len(chosen)-1]
		}
	}
}
