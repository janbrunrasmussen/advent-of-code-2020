package q

import (
	"io/ioutil"
	"reflect"
	"strings"

	//"strings"
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

	return strings.Join(s,"\n")

}

func TestParseAnswers_Simple(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	got := ParseAnswers(input)
	want := [][]string{{"abc"}, {"a","b","c"},{"ab","ac"}, {"a","a","a","a"}, {"b"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}



func TestCountAnswers_Simple(t *testing.T) {
	input := [][]string{{"abc"}, {"a","b","c"},{"ab","ac"}, {"a","a","a","a"}, {"b"}}
	got := CountAnswers(input)
	want := 11
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountAnswers_Input(t *testing.T) {
	input := ParseAnswers(loadInput())
	got := CountAnswers(input)
	want := 6259
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountUbiquitousAnswers_Simple(t *testing.T) {
	input := [][]string{{"abc"}, {"a","b","c"},{"ab","ac"}, {"a","a","a","a"}, {"b"}}
	got := CountUbiquitousAnswers(input)
	want := 6
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountUbiquitousAnswers_Input(t *testing.T) {
	input := ParseAnswers(loadInput())
	got := CountUbiquitousAnswers(input)
	want := 3178
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
