package navigator

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
	input := `F10
N3
F7
R90
F11`

	want := []Instruction{{"F", 10}, {"N", 3}, {"F", 7}, {"R", 90}, {"F", 11}}
	got := ParseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestManhattenDistance_Simple(t *testing.T) {
	input := `F10
N3
F7
R90
F11`
	want := 25
	got := ManhattenDistance(ParseInput(input))

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestManhattenDistance_Input(t *testing.T) {
	input := loadInput()
	want := 521
	got := ManhattenDistance(ParseInput(input))

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestManhattenDistance2_Simple(t *testing.T) {
	input := `F10
N3
F7
R90
F11`
	want := 286
	got := ManhattenDistance2(ParseInput(input))

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestManhattenDistance2_Input(t *testing.T) {
	input := loadInput()
	want := 22848
	got := ManhattenDistance2(ParseInput(input))

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
