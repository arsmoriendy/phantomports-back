package internal

import "testing"

func TestEmptyBinarySearch(t *testing.T) {
	arr := []int{}
	found, at := BinarySearch(&arr, 0)
	if found {
		t.Fatal(found, at)
	}
}

func TestSingleBinarySearch(t *testing.T) {
	arr := []int{1337}
	var findIdx uint = 0
	found, at := BinarySearch(&arr, arr[findIdx])
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestEvenBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}
	var findIdx uint = 1
	found, at := BinarySearch(&arr, arr[findIdx])
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestOddBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}
	var findIdx uint = 1
	found, at := BinarySearch(&arr, arr[findIdx])
	if !found || at != findIdx {
		t.Fatal(found, at)
	}
}

func TestEvenNFoundBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}

	toFind := 3

	found, at := BinarySearch(&arr, toFind)
	if found {
		t.Fatal(found, at)
	}
}

func TestOddNFoundBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}

	toFind := 2

	found, at := BinarySearch(&arr, toFind)
	if found {
		t.Fatal(found, at)
	}
}

func TestUnderNFoundBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8}

	toFind := 1

	found, at := BinarySearch(&arr, toFind)
	if found {
		t.Fatal(found, at)
	}
}

func TestOverNFoundBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5}

	toFind := 6

	found, at := BinarySearch(&arr, toFind)
	if found {
		t.Fatal(found, at)
	}
}
