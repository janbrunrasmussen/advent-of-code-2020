package passport_validator

import (
	"regexp"
	"strconv"
	"strings"
)

var RequiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func ParsePassport(input string) []string {
	p := strings.Split(input, "\n\n")
	return p
}

func IsValid(passport string) bool {

	for _, rf := range RequiredFields {
		if !strings.Contains(passport, rf) {
			return false
		}
	}
	return true
}

func CountValidPassports(input string) int {
	count := 0
	for _, p := range ParsePassport(input) {
		if IsValid(p) {
			count++
		}
	}
	return count
}

func ParsePassportFields(passport string) map[string]string {
	p := make(map[string]string, 0)

	s := strings.Replace(passport, "\n", " ", -1)
	fields := strings.Split(s, " ")

	for _, f := range fields {
		kv := strings.Split(f, ":")
		if len(kv) != 2 {
			continue
		}
		p[kv[0]] = kv[1]
	}

	return p
}

func IsFieldValid(field, value string) bool {
	switch field {
	case "byr":
		return len(value) == 4 && integer(value) >= 1920 && integer(value) <= 2002
	case "iyr":
		return len(value) == 4 && integer(value) >= 2010 && integer(value) <= 2020
	case "eyr":
		return len(value) == 4 && integer(value) >= 2020 && integer(value) <= 2030
	case "hgt":
		suffix := string(value[len(value)-2:])
		val := string(value[:len(value)-2])
		return (suffix == "cm" && integer(val) >= 150 && integer(val) <= 193) ||
			(suffix == "in" && integer(val) >= 59 && integer(val) <= 76)
	case "hcl":
		re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
		return re.MatchString(value)
	case "ecl":
		set := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, s := range set {
			if value == s {
				return true
			}
		}
		return false
	case "pid":
		re := regexp.MustCompile(`^[0-9]{9}$`)
		return re.MatchString(value)
	case "cid":
		return true
	}

	return false
}

func integer(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("AARHHH!")
	}
	return v
}

func CountValidPassportsStrict(input string) int {
	count := 0
	for _, p := range ParsePassport(input) {
		if !IsValid(p) {
			continue
		}
		// There is a loop to much here, because we don't need to
		// return all the values and the test if they are valid
		// we could instead parse and through away the rest of invalid
		// passports.
		// Next iteration...
		fields := ParsePassportFields(p)
		valid := false
		for k, v := range fields {
			valid = IsFieldValid(k, v)
			if !valid {
				break
			}
		}

		if valid {
			count++
		}
	}
	return count
}
