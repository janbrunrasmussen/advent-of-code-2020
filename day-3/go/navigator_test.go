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
	return strings.Join(s[:len(s)-1], "\n")
}

func TestParseMap_Simple(t *testing.T) {
	input := `..##
#...`

	want := [][]string{{".", ".", "#", "#"}, {"#", ".", ".", "."}}
	got := ParseMap(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountTrees_Simple(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	want := 7
	got := CountTrees(input,3,1, true)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountTrees_Input(t *testing.T) {
	input := loadInput()
	want := 237
	got := CountTrees(input,3,1, false)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountMoreTrees_Simple(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	want := 336
	got := CountMoreTrees(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountMoreTrees_Input(t *testing.T) {
	input := loadInput()

	want := 2106818610
	got := CountMoreTrees(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}
