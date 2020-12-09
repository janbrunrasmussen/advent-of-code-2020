package encoder

import (
	"strconv"
	"strings"
)

func ParseInput(input string) []int {
	lines := strings.Split(input, "\n")

	n := make([]int, 0)

	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic("AAAAHH cannot parse int")
		}
		n = append(n, i)
	}

	return n
}

func IsNextValid(rng []int, n int) bool {
	for ii, i := range rng {
		for jj, j := range rng {
			if ii == jj {
				continue
			}
			if i+j == n {
				return true
			}
		}
	}

	return false
}

func FindInvalidSequence(rng []int, preamble int) int {

	for k := 0; k < len(rng)-preamble; k++ {
		next := rng[k+preamble]
		if !IsNextValid(rng[k:k+preamble], next) {
			return next
		}
	}
	return 0
}

func FindContiguesSet(rng []int, invalid int) []int {
	for i := 0; i < len(rng)-2; i++ {
		sum := rng[i]
		for j := i + 1; j < len(rng)-1; j++ {
			if rng[j] == invalid {
				continue
			}

			sum += rng[j]

			if sum == invalid {
				return rng[i : j+1]
			}
		}
	}
	return nil
}

func CalculateWeakness(rng []int) int {
	max := 0
	min := rng[0]

	for _, r := range rng {
		if r < min {
			min = r
		}
		if r > max {
			max = r
		}
	}

	return min + max
}
