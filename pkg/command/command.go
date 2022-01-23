package command

import "strings"

const delimiter = " "

type Command struct {
	Command string
	Options []string
}

func NewCommand(s string) *Command {
	strs := strings.Split(s, delimiter)

	// "!" check ensure here always has at least one element
	command := strs[0][1:]

	return &Command{
		Command: command,
		Options: strs[1:],
	}
}
