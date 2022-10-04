package main

import "fmt"

func BinarySearch(searchedItem int, searchedArr []int) string {
	start := 0
	end := len(searchedArr) - 1

	for start <= end {
		mid := (start + end) / 2
		if searchedArr[mid] == searchedItem {
			return fmt.Sprintf("%d index", mid)
		} else if searchedArr[mid] > searchedItem {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return "not found"
}

func main() {
	fmt.Println(BinarySearch(2, []int{1, 2, 3, 4}))
}
