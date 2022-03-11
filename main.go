package main

import (
	"fmt"
)

func findSequencial(nums []int) []int {
	var sequen []int //nill
	var lastIndex int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			sequen = append(sequen, nums[i])
			lastIndex = i
		}
	}
	sequen = append(sequen, nums[lastIndex+1])
	return sequen
}

func reverseNum(num []int) []int {
	if len(num) == 0 {
		return num
	}
	return append(reverseNum(num[1:]), num[0])
}

func findReverse(nums, arr []int) bool {
	for i := 0; i < len(arr)-len(nums); i++ {
		compare := arr[i : len(nums)-1]
		fmt.Println("Compare", compare)
		for j := 0; j < len(nums); j++ {
			if compare[j] != nums[j] {
				return false
			}
		}
	}
	return true
}

func maxArray(arr []int) int {
	var max int
	sequence := findSequencial(arr)
	fmt.Println(sequence)
	revsequence := reverseNum(sequence)
	fmt.Println(revsequence)
	isExist := findReverse(arr, revsequence)
	fmt.Println(isExist)
	if isExist {
		max = sequence[len(sequence)-1]
	} else {
		max = -1
	}

	return max
}

func main() {
	fmt.Println(maxArray([]int{7, 1, 2, 3, 8, 6, 3, 2, 1})) // 3
	fmt.Println(maxArray([]int{7, 1, 2, 9, 7, 2, 1}))       // 2
}
