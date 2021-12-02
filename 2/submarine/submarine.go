package submarine

import (
	"strconv"
	"strings"
)

type Submarine struct {
	Depth   int32
	Forward int32
	Aim     int32
}

func (submarine *Submarine) Move(command string) {
	commands_tuple := strings.Split(command, " ")
	action := commands_tuple[0]
	value, _ := strconv.Atoi(commands_tuple[1])

	switch action {
	case "down":
		//submarine.dive(int32(value))
		submarine.aimDown(int32(value))
	case "up":
		//submarine.dive(-int32(value))
		submarine.aimUp(int32(value))
	case "forward":
		submarine.moveForward(int32(value))
	}
}

func (submarine *Submarine) dive(x int32) {
	submarine.Depth += x
}

func (submarine *Submarine) aimUp(x int32) {
	submarine.Aim -= x
}

func (submarine *Submarine) aimDown(x int32) {
	submarine.Aim += x
}

func (submarine *Submarine) moveForward(x int32) {
	submarine.Forward += x
	submarine.Depth += submarine.Aim * x
}

func (submarine *Submarine) Total() int32 {
	return submarine.Depth * submarine.Forward
}
