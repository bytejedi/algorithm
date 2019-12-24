package main

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	mergeSort(intervals)
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

func DeepCopy(original [][]int) [][]int {
	duplicate := make([][]int, len(original))
	copy(duplicate, original)
	return duplicate
}

func mergeSort(intervals [][]int) {
	if len(intervals) > 1 {
		h := len(intervals) / 2
		left := DeepCopy(intervals[:h])
		right := DeepCopy(intervals[h:])
		mergeSort(left)
		mergeSort(right)

		l, r := 0, 0
		for i := 0; i < len(intervals); i++ {
			var lval []int
			var rval []int
			if l < len(left) {
				lval = left[l]
			}
			if r < len(right) {
				rval = right[r]
			}

			if (lval != nil && rval != nil && lval[0] < rval[0]) || rval == nil {
				intervals[i] = lval
				l += 1
			}
			if (rval != nil && lval != nil && lval[0] >= rval[0]) || lval == nil {
				intervals[i] = rval
				r += 1
			}
		}
	}
}
