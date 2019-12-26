package minimum_number_of_arrows_to_burst_balloons

//在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以y坐标并不重要，因此只要知道开始和结束的x坐标就足够了。开始坐标总是小于结束坐标。平面内最多存在104个气球。
//
//一支弓箭可以沿着x轴从不同点完全垂直地射出。在坐标x处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
//
//Example:
//
//输入:
//[[10,16], [2,8], [1,6], [7,12]]
//
//输出:
//2
//
//解释:
//对于该样例，我们可以在x = 6（射爆[2,8],[1,6]两个气球）和 x = 11（射爆另外两个气球）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func FindMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	mergeSort(points)
	prev := points[0]
	result := 1
	for i := range points {
		if points[i][0] > prev[1] {
			result++
			prev = points[i]
		}
	}
	return result
}

func deepCopy(original [][]int) [][]int {
	duplicate := make([][]int, len(original))
	copy(duplicate, original)
	return duplicate
}

func mergeSort(intervals [][]int) {
	if len(intervals) <= 1 {
		return
	}
	mid := len(intervals) / 2
	left := deepCopy(intervals[:mid])
	right := deepCopy(intervals[mid:])

	mergeSort(left)
	mergeSort(right)

	l, r := 0, 0
	for i := range intervals {
		var lval []int
		var rval []int
		if l < len(left) {
			lval = left[l]
		}
		if r < len(right) {
			rval = right[r]
		}

		if (lval != nil && rval != nil && lval[1] < rval[1]) || rval == nil {
			intervals[i] = lval
			l++
		}
		if (lval != nil && rval != nil && lval[1] >= rval[1]) || lval == nil {
			intervals[i] = rval
			r++
		}
	}
}
