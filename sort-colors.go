package main

func SortColors(nums []int) {
	i, j, current := 0, len(nums)-1, 0
	var tmp int
	for current <= j {
		if nums[current] == 0 {
			tmp = nums[current]
			nums[current] = nums[i]
			nums[i] = tmp
			current++
			i++
		} else if nums[current] == 2 {
			tmp = nums[current]
			nums[current] = nums[j]
			nums[j] = tmp
			j--
		} else {
			current++
		}
	}
}
