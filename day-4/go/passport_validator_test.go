package passport_validator

import (
	"io/ioutil"
	"reflect"
	//"strings"
	"testing"
)

func loadInput() string {
	data, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	return string(data)

	// //Remove last empty line
	// s := strings.Split(string(data), "\n")
	// return s[:len(s)-1]
}

func TestParsePassport_Simple(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	want := []string{
		`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`,
		`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`,
		`hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`,
		`hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`,
	}
	got := ParsePassport(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestIsValid_Table(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"Test 1", `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`, true},
		{"Test 2", `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`, false},
		{"Test 3", `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`, true},
		{"Test 4", `hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.input)
			if got != tt.want {
				t.Errorf("got: %v, want: %v.", got, tt.want)
			}
		})
	}
}

func TestCountValid_Simple(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	want := 2
	got := CountValidPassports(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountValid_Input(t *testing.T) {
	input := loadInput()
	want := 196
	got := CountValidPassports(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestIsFieldValid_Table(t *testing.T) {
	var tests = []struct {
		name  string
		Field string
		Value string
		want  bool
	}{
		{"Field byr 1 valid", "byr", "2002", true},
		{"Field byr 2 invalid", "byr", "2003", false},
		{"Field hgt 1 valid", "hgt", "60in", true},
		{"Field hgt 2 valid", "hgt", "190cm", true},
		{"Field hgt 3 invalid", "hgt", "190in", false},
		{"Field hgt 4 invalid", "hgt", "190", false},
		{"Field hcl 1 valid", "hcl", "#123abc", true},
		{"Field hcl 2 invalid", "hcl", "#123abz", false},
		{"Field hcl 3 invalid", "hcl", "123abc", false},
		{"Field ecl 1 valid", "ecl", "brn", true},
		{"Field ecl 2 invalid", "ecl", "wat", false},
		{"Field pid 1 valid", "pid", "000000001", true},
		{"Field pid 2invalid", "pid", "0123456789", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFieldValid(tt.Field, tt.Value)
			if got != tt.want {
				t.Errorf("got: %v, want: %v.", got, tt.want)
			}
		})
	}
}

func TestParsePassportFields_Simple(t *testing.T) {
	input := `ecl:gry pid:860033327
eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`
	want := map[string]string{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}

	got := ParsePassportFields(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v.", got, want)
	}

}

func TestCountValidStrict_Simple(t *testing.T) {
	input := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	want := 4
	got := CountValidPassportsStrict(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}

func TestCountValidStrict_Input(t *testing.T) {
	input := loadInput()
	want := 150
	got := CountValidPassportsStrict(input)

	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
	}
}
