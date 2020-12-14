package seating

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func loadInput() string {
	data, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	//Remove last empty line
	s := strings.Split(string(data), "\n")
	s = s[:len(s)-1]

	return strings.Join(s, "\n")

}

func TestParseInput_Simple(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL`

	want := [][]string{{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"}}
	got := ParseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountOccupiedSeats_Simple(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	want := 37
	got := CountOccupiedSeats(ParseInput(input), 1, 4, false)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountOccupiedSeats_Input(t *testing.T) {
	input := loadInput()
	want := 2251
	got := CountOccupiedSeats(ParseInput(input), 1, 4, false)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountOccupiedSeatsPart2_Simple(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	want := 26
	got := CountOccupiedSeats(ParseInput(input), -1, 5, false)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountOccupiedSeatsPart2_Input(t *testing.T) {
	input := loadInput()
	want := 2019
	got := CountOccupiedSeats(ParseInput(input), -1, 5, false)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
