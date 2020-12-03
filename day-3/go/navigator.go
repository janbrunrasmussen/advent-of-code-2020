package navigator

import (
	"fmt"
	"math"
	"strings"
)

func ParseMap(input string) [][]string {

	lines := strings.Split(input, "\n")
	m := make([][]string, len(lines))

	for i, l := range lines {
		m[i] = strings.Split(l, "")
	}

	return m
}

func CountTrees(input string, stepRight, stepDown int, print bool) int {

	m := ParseMap(input)
	wm := make([][]string, len(m))

	//expand map (number of maps: rows * steps right / diveded by columns - and then round up)
	for i, r := range m {
		for j := 0; j < int(math.Ceil(float64(stepRight)*float64(len(m))/float64(len(m[0])))); j++ {
			wm[i] = append(wm[i], r...)
		}
	}

	pos := []int{0, 0}
	trees := 0
	
	for {
		marker := "O"

		if wm[pos[0]][pos[1]] == "#" {
			marker = "X"
			trees++
		}

		//mark position
		wm[pos[0]][pos[1]] = marker

		//new position
		pos[0] += stepDown
		pos[1] += stepRight

		if pos[0] >= len(wm) {
			break
		}
	}

	// print path
	if print {
		for _, r := range wm {
			fmt.Println(r)
		}
	}

	return trees
}

func CountMoreTrees(input string) int {

	steps := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	MoreTrees := 1
	for _, s := range steps {
		MoreTrees = MoreTrees * CountTrees(input, s[0], s[1], false)
	}

	return MoreTrees
}
