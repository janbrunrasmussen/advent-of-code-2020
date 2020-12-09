package compiler

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

func TestNewInstructions_Simple(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	got := NewInstructions(input)
	want := []Instruction{
		{"nop", +0, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", +4, 0, 0},
		{"acc", +3, 0, 0},
		{"jmp", -3, 0, 0},
		{"acc", -99, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", -4, 0, 0},
		{"acc", +6, 0, 0},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestRunBootCode_Simple(t *testing.T) {
	input := []Instruction{
		{"nop", +0, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", +4, 0, 0},
		{"acc", +3, 0, 0},
		{"jmp", -3, 0, 0},
		{"acc", -99, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", -4, 0, 0},
		{"acc", +6, 0, 0},
	}
	got, gotExitCode := RunBootCode(input)
	want := 5
	wantExitCode := 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
	if !reflect.DeepEqual(gotExitCode, wantExitCode) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestRunBootCode_Input(t *testing.T) {
	input := NewInstructions(loadInput())
	got, gotExitCode := RunBootCode(input)
	want := 1801
	wantExitCode := 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
	if !reflect.DeepEqual(gotExitCode, wantExitCode) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFixBootCode_Simple(t *testing.T) {
	input := []Instruction{
		{"nop", +0, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", +4, 0, 0},
		{"acc", +3, 0, 0},
		{"jmp", -3, 0, 0},
		{"acc", -99, 0, 0},
		{"acc", +1, 0, 0},
		{"jmp", -4, 0, 0},
		{"acc", +6, 0, 0},
	}
	got := FixBootCode(input)
	want := 8

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestFixBootCode_Input(t *testing.T) {
	input := NewInstructions(loadInput())
	got := FixBootCode(input)
	want := 2060

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
