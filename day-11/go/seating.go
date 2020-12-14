package seating

import (
	"fmt"
	"strings"
)

func ParseInput(input string) [][]string {

	lines := strings.Split(input, "\n")
	m := make([][]string, len(lines))

	for i, l := range lines {
		m[i] = strings.Split(l, "")
	}

	return m
}

func CountOccupiedSeats(seats [][]string, reach, threshold int, print bool) int {

	// If we don't set reach, set to number of rows
	if reach == -1 {
		reach = len(seats)
	}

	// Let's make clean copy
	next := make([][]string, len(seats))
	for i := 0; i < len(seats); i++ {
		innerLen := len(seats[i])
		next[i] = make([]string, innerLen)
		for j := 0; j < innerLen; j++ {
			next[i][j] = seats[i][j]
		}
	}

	directions := map[string][]int{
		"U":  {-1, 0},
		"UR": {-1, 1},
		"R":  {0, 1},
		"DR": {1, 1},
		"D":  {1, 0},
		"DL": {1, -1},
		"L":  {0, -1},
		"UL": {-1, -1},
	}

	iterations := 0
	for {

		current := make([][]string, len(seats))
		for i := 0; i < len(seats); i++ {
			innerLen := len(seats[i])
			current[i] = make([]string, innerLen)
			for j := 0; j < innerLen; j++ {
				current[i][j] = next[i][j]
			}
		}

		diff := 0
		iterations++
		occupied := 0
		for i, c := range current {
			for j := range c {
				seat := current[i][j]
				if seat == "." {
					continue
				}
				adj := 0
				adjSeats := make(map[string]string, 0)
				for q := 1; q <= reach; q++ {
					for dir, k := range directions {

						// Direction has already been checked
						if _, ok := adjSeats[dir]; ok {
							continue
						}

						ii := k[0] * q
						jj := k[1] * q

						adjSeat := ""
						//out of range downwards
						if i+ii < 0 || j+jj < 0 {
							adjSeat = "-1"
						}
						//out of range upwards
						if i+ii >= len(seats) || j+jj >= len(seats[i]) {
							adjSeat = "-1"
						}

						if adjSeat == "" {
							adjSeat = current[i+ii][j+jj]
						}

						if adjSeat == "." {
							continue
						}

						adjSeats[dir] = adjSeat

						if adjSeat == "#" {
							adj++
						}

						if seat == "L" && adj > 0 {
							break
						}
						if adj > threshold {
							break
						}
					}

				}

				if seat == "L" && adj == 0 {
					next[i][j] = "#"
					diff++
				} else if seat == "#" && adj >= threshold {
					next[i][j] = "L"
					diff++
					occupied++
				} else if seat == "#" {
					occupied++
				}
			}
		}

		if print {
			fmt.Println("======", iterations, diff)
			for _, n := range next {
				fmt.Println(strings.Join(n, ""))
			}
		}
		if diff == 0 {
			return occupied
		}
	}
}
