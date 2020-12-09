package encoder

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
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	got := ParseInput(input)
	want := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

// 26 would be a valid next number, as it could be 1 plus 25 (or many other pairs, like 2 and 24).
// 49 would be a valid next number, as it is the sum of 24 and 25.
// 100 would not be valid; no two of the previous 25 numbers sum to 100.
// 50 would also not be valid; although 25 appears in the previous 25 numbers, the two numbers in the pair must be different.

func TestNextIsValid_Table(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	var tests = []struct {
		name   string
		number int
		want   bool
	}{
		{"Test 1", 26, true},
		{"Test 2", 49, true},
		{"Test 3", 100, false},
		{"Test 4", 50, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNextValid(input, tt.number)
			if got != tt.want {
				t.Errorf("got: %v, want: %v.", got, tt.want)
			}
		})
	}
}

func TestFindInvalidSequence_Simple1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 100}
	got := FindInvalidSequence(input, 25)
	want := 100
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFindInvalidSequence_Table(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	var tests = []struct {
		name     string
		preamble int
		want     int
	}{
		{"Test 1", 5, 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindInvalidSequence(input, tt.preamble)
			if got != tt.want {
				t.Errorf("got: %v, want: %v.", got, tt.want)
			}
		})
	}
}

func TestFindInvalidSequence_Simple(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	got := FindInvalidSequence(input, 5)
	want := 127
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFindInvalidSequence_Input(t *testing.T) {
	input := ParseInput(loadInput())
	got := FindInvalidSequence(input, 25)
	want := 217430975
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFindContiguesSet_Simple(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	got := FindContiguesSet(input, 127)
	want := []int{15, 25, 47, 40}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFindWeakness_Simple(t *testing.T) {
	input := []int{15, 25, 47, 40}
	got := CalculateWeakness(input)
	want := 62
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFindWeakness_Input(t *testing.T) {
	input := FindContiguesSet(ParseInput(loadInput()), 217430975)
	got := CalculateWeakness(input)
	want := 28509180
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
