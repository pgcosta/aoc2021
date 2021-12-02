package submarine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondExample(t *testing.T) {
	sub := Submarine{}

	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	sub.Move(commands[0])
	sub.Move(commands[1])
	sub.Move(commands[2])
	sub.Move(commands[3])
	sub.Move(commands[4])
	sub.Move(commands[5])

	assert.Equal(t, sub.Forward, int32(15))
	assert.Equal(t, sub.Depth, int32(60))
	assert.Equal(t, sub.Total(), int32(900))
}
