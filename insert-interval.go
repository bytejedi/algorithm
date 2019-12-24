package main

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) < 1 {
		return [][]int{newInterval}
	}
	i, j := 0, len(intervals)
	for i < j {
		mid := i + (j-i)/2
		if intervals[mid][0] < newInterval[0] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	intervals = append(intervals, []int{})
	copy(intervals[i+1:], intervals[i:])
	intervals[i] = newInterval

	merged := [][]int{intervals[0]}
	intervals = intervals[1:]
	for i := range intervals {
		if intervals[i][0] <= merged[len(merged)-1][1] { // need merge
			if merged[len(merged)-1][1] < intervals[i][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
