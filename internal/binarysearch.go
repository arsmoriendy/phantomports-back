package internal

func BinarySearch(arr *[]int, toFind int) (bool, uint) {
	n := uint(len(*arr))
	if n == 0 {
		return false, 0
	}

	return bs(arr, toFind, 0, n-1, n/2)
}

func bs(arr *[]int, toFind int, startIdx uint, endIdx uint, pivotIdx uint) (bool, uint) {
	pivotVal := (*arr)[pivotIdx]

	if endIdx-startIdx < 2 {
		return pivotVal == toFind, pivotIdx
	}

	if toFind < pivotVal {
		return bs(arr, toFind, startIdx, pivotIdx, startIdx+(pivotIdx-startIdx)/2)
	} else if toFind > pivotVal {
		pivotIdx++
		return bs(arr, toFind, pivotIdx, endIdx, pivotIdx+(endIdx-pivotIdx)/2)
	} else {
		return true, pivotIdx
	}
}
