package main

func findSequencial(nums []int) []int {
	var sequen []int
	var lastIndex int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i]+1 {
			sequen = append(sequen, nums[i])
			lastIndex = i
		} else {
			if sequen != nil {
				break
			}
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

func findReverse(arr []int, nums []int) bool {
	var isExist bool
	for i := 0; i <= len(arr)-len(nums); i++ {
		isExist = true
		compare := arr[i : len(nums)+i]
		for j := 0; j < len(nums); j++ {
			if compare[j] != nums[j] {
				isExist = false
			}
		}
		if isExist {
			return true
		}
	}
	return false
}

func mainProgram(arr []int) int {
	var max int = -1
	sequence := findSequencial(arr)
	revsequence := reverseNum(sequence)
	for i := len(revsequence) - 1; i >= 0; i-- {
		isExist := findReverse(arr, revsequence[i:])
		if isExist {
			max = revsequence[i]
		}
	}
	return max
}
