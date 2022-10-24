package calculadora

import "sort"

func SortIntegerSliceInAscendingOrder(integerSlice []int) []int {
	sort.Slice(integerSlice, func(i, j int) bool {
		return integerSlice[i] < integerSlice[j]
	})
	return integerSlice
}
