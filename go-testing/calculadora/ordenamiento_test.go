package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAscendingSort(t *testing.T) {
	errorMessage := "SortIntegerSliceInAscendingOrder function returned %v, but the expected result is %v"

	sliceToSort := []int{1, 2, 3, 4}
	expectedResult := []int{1, 2, 3, 4}
	result := SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)

	sliceToSort = []int{4, 3, 2, 1}
	expectedResult = []int{1, 2, 3, 4}
	result = SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)

	sliceToSort = []int{-6, 6, -26, 13}
	expectedResult = []int{-26, -6, 6, 13}
	result = SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)

	sliceToSort = []int{-6, 60000, -26, 1300000}
	expectedResult = []int{-26, -6, 60000, 130000}
	result = SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)

	sliceToSort = []int{-0, 60000, -0, 1300000}
	expectedResult = []int{0, 0, 60000, 130000}
	result = SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)

	sliceToSort = []int{100_000_000_001, 100_000_000_000, 100_000_000_000, 100_000_000_000}
	expectedResult = []int{100_000_000_000, 100_000_000_000, 100_000_000_000, 100_000_000_001}
	result = SortIntegerSliceInAscendingOrder(sliceToSort)
	assert.Equal(t, expectedResult, result, errorMessage, expectedResult, result)
}
