package day06

import "slices"

func OutOfBounds(pos []int, w, h int) bool {
	xBound := pos[0] < 0 || pos[0] > w-1
	yBound := pos[1] < 0 || pos[1] > h-1
	return xBound || yBound
}

func GuardWalk(pos, dir []int) []int {
	return []int{pos[0] + dir[0], pos[1] + dir[1]}
}

func ContainsIntArr(array [][]int, value []int) bool {
	return slices.ContainsFunc(array, func(member []int) bool {
		// Compare elements
		for idx, cur := range member {
			if cur != value[idx] {
				return false
			}
		}
		return true
	})
}
