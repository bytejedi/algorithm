// https://leetcode-cn.com/problems/positions-of-large-groups/
package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	l := len(s)
	if l < 3 {
		return nil
	}

	var res [][]int
	var i, j int
	for i < l {
		if s[i] != s[j] {
			if i-j >= 3 {
				res = append(res, []int{j, i - 1})
			}
			j = i
		} else {
			i++
		}
	}
	if i-j >= 3 {
		res = append(res, []int{j, i - 1})
	}
	return res
}

func main() {
	fmt.Println(largeGroupPositions("baaasfseeeeee"))
}
