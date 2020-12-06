package seater

import (
	"sort"
)

type Seat struct {
	Row    int
	Column int
	ID     int
}

func NewSeat(row, column int) Seat {
	return Seat{
		Row:    row,
		Column: column,
		ID:     row*8 + column,
	}
}

func ParseSeat(input string) Seat {

	r := solveBSP(0, 127, "B", "F", input[:7])
	c := solveBSP(0, 7, "R", "L", input[7:])

	s := NewSeat(r, c)
	return s
}

func solveBSP(l, u int, uC, lC string, seq string) int {
	for _, c := range seq {
		m := (u - l) / 2
		if string(c) == lC {
			u = l + m
		} else if string(c) == uC {
			l = u - m
		}
	}

	return l
}

func MaxSeatID(input []string) int {
	maxID := 0

	for _, i := range input {
		s := ParseSeat(i)

		if s.ID > maxID {
			maxID = s.ID
		}
	}

	return maxID
}

func FindEmptySeatID(input []string) int {
	IDs := make([]int, 0)
	for _, i := range input {
		IDs = append(IDs, ParseSeat(i).ID)
	}

	sort.Ints(IDs)

	for i := 0; i < len(IDs)-1; i++ {
		if IDs[i]+1 != IDs[i+1] {
			return IDs[i] + 1
		}
	}

	return 0
}
