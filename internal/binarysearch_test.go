package internal

import "testing"

func TestEmptyBinarySearch(t *testing.T) {
	arr := []int{}
	n := uint(len(arr))
	found, at := BinarySearch(&arr, 0, 0, n-1, n/2)
	if found {
		t.Fatal(found, at)
	}
}

func TestSingleBinarySearch(t *testing.T) {
	arr := []int{1337}
	var findIdx uint = 0
	n := uint(len(arr))
	found, at := BinarySearch(&arr, arr[findIdx], 0, n-1, n/2)
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestEvenBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}
	var findIdx uint = 1
	n := uint(len(arr))
	found, at := BinarySearch(&arr, arr[findIdx], 0, n-1, n/2)
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestOddBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}
	var findIdx uint = 1
	n := uint(len(arr))
	found, at := BinarySearch(&arr, arr[findIdx], 0, n-1, n/2)
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestEvenNFoundBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}
	n := uint(len(arr))

	toFind := 3

	found, at := BinarySearch(&arr, toFind, 0, n-1, n/2)
	if found {
		t.Fatal(found, at)
	}
}

func TestOddNFoundBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}
	n := uint(len(arr))

	toFind := 2

	found, at := BinarySearch(&arr, toFind, 0, n-1, n/2)
	if found {
		t.Fatal(found, at)
	}
}

func TestUnderNFoundBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}
	n := uint(len(arr))

	toFind := 1

	found, at := BinarySearch(&arr, toFind, 0, n-1, n/2)
	if found {
		t.Fatal(found, at)
	}
}

func TestOverNFoundBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}
	n := uint(len(arr))

	toFind := 6

	found, at := BinarySearch(&arr, toFind, 0, n-1, n/2)
	if found {
		t.Fatal(found, at)
	}
}
