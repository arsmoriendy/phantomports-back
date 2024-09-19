package internal

func BinarySearch(arr *[]int, toFind int, startIdx uint, endIdx uint, pivotIdx uint) (bool, uint) {
	pivotVal := (*arr)[pivotIdx]

	if endIdx-startIdx < 2 {
		return pivotVal == toFind, pivotIdx
	}

	if toFind < pivotVal {
		return BinarySearch(arr, toFind, startIdx, pivotIdx, startIdx+(pivotIdx-startIdx)/2)
	} else if toFind > pivotVal {
		pivotIdx++
		return BinarySearch(arr, toFind, pivotIdx, endIdx, pivotIdx+(endIdx-pivotIdx)/2)
	} else {
		return true, pivotIdx
	}
}
