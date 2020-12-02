package password_validator

import (
	"regexp"
	"strconv"
	"strings"
)

type record struct {
	char     string
	lower    int
	upper    int
	password string
}

func Parse(i string) record {

	var rec record
	var err error

	s := strings.SplitN(i, "-", 2)

	rec.lower, err = strconv.Atoi(s[0])
	if err != nil {
		panic("not int")
	}
	s = strings.SplitN(s[1], " ", 2)

	rec.upper, err = strconv.Atoi(s[0])
	if err != nil {
		panic("not int")
	}
	s = strings.SplitN(s[1], ": ", 2)

	rec.char = s[0]
	rec.password = s[1]

	return rec
}

func ParseRegex(i string) record {

	var rec record
	var err error

	r := regexp.MustCompile("(^\\d*)-(\\d*)\\s([a-zA-Z]):\\s([a-zA-Z].*)")
	s := r.FindStringSubmatch(i)

	rec.lower, err = strconv.Atoi(s[1])
	if err != nil {
		panic("not int")
	}
	rec.upper, err = strconv.Atoi(s[2])
	if err != nil {
		panic("not int")
	}
	rec.char = s[3]
	rec.password = s[4]

	return rec
}

func Valid(input []string) int {

	valid := 0

	for _, i := range input {
		//r := Parse(i)
		r := ParseRegex(i)

		charCount := strings.Count(r.password, r.char)

		if charCount >= r.lower && charCount <= r.upper {
			valid++
		}

	}

	return valid
}

func ValidStrict(input []string) int {

	valid := 0

	for _, i := range input {
		//r := Parse(i)
		r := ParseRegex(i)

		if len(r.password) < r.upper {
			continue
		}

		if (string(r.password[r.lower-1]) == r.char) !=
			(string(r.password[r.upper-1]) == r.char) {
			valid++
		}
	}

	return valid
}
