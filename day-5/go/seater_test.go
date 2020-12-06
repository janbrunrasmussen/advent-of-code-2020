package seater

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func loadInput() []string {
	data, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	//Remove last empty line
	s := strings.Split(string(data), "\n")
	return s[:len(s)-1]
}

func TestParseSeat_Simple(t *testing.T) {
	got := ParseSeat("FBFBBFFRLR")
	want := Seat{Row: 44, Column: 5, ID: 357}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestParseSeat_Table(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  Seat
	}{
		{"test 1", "BFFFBBFRRR", Seat{70, 7, 567}},
		{"test 1", "FFFBBBFRRR", Seat{14, 7, 119}},
		{"test 1", "BBFFBBFRLL", Seat{102, 4, 820}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSeat(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v.", got, tt.want)
			}
		})
	}
}

func TestMaxseatID_Simple(t *testing.T) {
	got := MaxSeatID([]string{"BFFFBBFRRR", "FBFBBFFRLR"})
	want := 567
	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestMaxseatID_Input(t *testing.T) {
	got := MaxSeatID(loadInput())
	want := 890
	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestEmptySeats_Input(t *testing.T) {
	got := FindEmptySeatID(loadInput())
	want := 651
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
