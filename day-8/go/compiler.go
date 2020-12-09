package compiler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Name       string
	Value      int
	Executions int
	Sequence   int
}

func NewInstructions(input string) []Instruction {
	ins := make([]Instruction, 0)

	re := regexp.MustCompile("(^\\w{3}) ((\\+|\\-)\\d*$)")
	for _, i := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(i)
		v, err := strconv.Atoi(m[2])
		if err != nil {
			fmt.Printf("string is not int: %s \n", m[2])
			panic("AARGGH - not int")
		}
		ins = append(ins, Instruction{m[1], v, 0, 0})
	}

	return ins
}

func RunBootCode(ins []Instruction) (int, int) {

	exitCode := -1
	i := 0
	seq := 1
	ACC := 0

	for exitCode == -1 {

		// code is looping
		if ins[i].Executions > 0 {
			exitCode = 1
			break
		}

		ins[i].Sequence = seq
		seq++

		switch n := ins[i]; n.Name {
		case "nop":
			i++
		case "acc":
			ACC += n.Value
			ins[i].Executions++
			i++
		case "jmp":
			ins[i].Executions++
			i += n.Value
		}

		//end of sequence
		if i == len(ins) {
			exitCode = 0
			break
		}
	}

	// Print the execution
	// for j, in := range ins {
	// 	loop := ""
	// 	if i == j {
	// 		loop = "*"
	// 	}
	// 	fmt.Println(in, loop)
	// }

	return ACC, exitCode
}

func FixBootCode(ins []Instruction) int {

	for i := 0; i < len(ins); i++ {
		mutatedIns := make([]Instruction, len(ins))
		copy(mutatedIns, ins)

		switch ins[i].Name {
		case "jmp":
			mutatedIns[i].Name = "nop"
		case "nop":
			mutatedIns[i].Name = "jmp"
		default:
			continue
		}

		ACC, exitCode := RunBootCode(mutatedIns)
		if exitCode == 0 {
			return ACC
		}
	}

	return 0
}
