package navigator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	Action string
	Value  int
}

type Position struct {
	// East-West
	ew int
	// North South
	ns int
}

func ParseInput(input string) []Instruction {

	ins := make([]Instruction, 0)

	lines := strings.Split(input, "\n")

	for _, l := range lines {
		value, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("Counld not parse int: %v\n", value)
			panic("ARRHGGHH")
		}

		ins = append(ins, Instruction{string(l[0]), value})

	}
	return ins
}

// Action N means to move north by the given value.
// Action S means to move south by the given value.
// Action E means to move east by the given value.
// Action W means to move west by the given value.
// Action L means to turn left the given number of degrees.
// Action R means to turn right the given number of degrees.
// Action F means to move forward by the given value in the direction the ship is currently facing.
func ManhattenDistance(ins []Instruction) int {

	dir := 90
	pos := Position{0, 0}
	dirs := map[int]string{
		0:   "N",
		90:  "E",
		180: "S",
		270: "W",
	}

	for _, i := range ins {
		switch i.Action {
		case "L":
			dir = dir - i.Value
		case "R":
			dir = dir + i.Value
		case "F":
			pos.move(dirs[dir], i.Value)
		default:
			pos.move(i.Action, i.Value)
		}

		// Baseline degrees
		if dir < 0 {
			dir = 360 + dir
		}
		if dir >= 360 {
			dir = 360 - dir
		}
		dir = abs(dir)
	}

	return abs(pos.ns) + abs(pos.ew)
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func (p *Position) move(dir string, val int) {
	switch dir {
	case "N":
		p.ns += val
	case "S":
		p.ns -= val
	case "E":
		p.ew += val
	case "W":
		p.ew -= val
	}
}

// Action N means to move the waypoint north by the given value.
// Action S means to move the waypoint south by the given value.
// Action E means to move the waypoint east by the given value.
// Action W means to move the waypoint west by the given value.
// Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
// Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
// Action F means to move forward to the waypoint a number of times equal to the given value.
func ManhattenDistance2(ins []Instruction) int {

	pos := Position{0, 0}
	wp := Position{10, 1}

	for _, i := range ins {
		switch i.Action {
		case "L":
			wp.rotate(-1 * i.Value)
		case "R":
			wp.rotate(i.Value)
		case "F":
			pos.moveVessel(wp, i.Value)
		default:
			wp.moveWaypoint(i.Action, i.Value)
		}
	}

	return abs(pos.ns) + abs(pos.ew)
}

func (p *Position) moveVessel(wp Position, val int) {
	p.ew += wp.ew * val
	p.ns += wp.ns * val
}

func (p *Position) moveWaypoint(dir string, val int) {
	switch dir {
	case "N":
		p.ns += val
	case "S":
		p.ns -= val
	case "E":
		p.ew += val
	case "W":
		p.ew -= val
	}
}

// 90 (x,y)=(−y,x)
// 180 (x,y)=(−x,−y)
// 270 (x,y)=(y,−x)
func (p *Position) rotate(deg int) {
	ew := p.ew
	ns := p.ns

	switch deg {
	case 90, -270:
		p.ew = ns
		p.ns = -ew
	case 180, -180:
		p.ew = -ew
		p.ns = -ns
	case -90, 270:
		p.ew = -ns
		p.ns = ew
	}
}
