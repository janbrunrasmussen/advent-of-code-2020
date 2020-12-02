package password_validator

import (
	"io/ioutil"
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

func TestParse_Simple(t *testing.T) {
	input := "1-3 a: abcde"
	want := record{"a", 1, 3, "abcde"}
	got := Parse(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestParse_Regex(t *testing.T) {
	input := "1-3 a: abcde"
	want := record{"a", 1, 3, "abcde"}
	got := ParseRegex(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestValid_Simple(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: cccccccc"}
	want := 2
	got := Valid(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestValid_Input(t *testing.T) {
	input := loadInput()
	want := 383
	got := Valid(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestValidStrict_Simple(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: cccccccc"}
	want := 1
	got := ValidStrict(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestValidStrict_Input(t *testing.T) {
	input := loadInput()
	want := 272
	got := ValidStrict(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
