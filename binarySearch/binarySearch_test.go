package main

import "testing"

func TestBinarySerach(t *testing.T) {
	testCases := []struct {
		searchItem     int
		searchArr      []int
		expectedResult string
	}{
		{
			4,
			[]int{1, 2, 3, 4},
			"3 index",
		},
		{
			-1,
			[]int{1, 2, 3, 4},
			"not found",
		},
	}

	for _, val := range testCases {
		res := BinarySearch(val.searchItem, val.searchArr)
		if val.expectedResult != res {
			t.Errorf("Unexpected values. Got %v want %v\n", val.expectedResult, res)
		}
	}
}
