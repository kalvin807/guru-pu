package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const delimiter = " "

type CommandHandler func(b *Bot, m *discordgo.MessageCreate, c *Command)

var handlerMaps map[string]CommandHandler = map[string]CommandHandler{
	"ping":     HandlePing,
	"register": HandleRegisterChannel,
	"add":      HandleRegisterGroup,
}

type Command struct {
	Command string
	Options []string
}

func NewCommand(s string) *Command {
	strs := strings.Split(s, delimiter)
	// "!gurupu" check ensure here always has at least one element
	command := strs[0][8:]

	return &Command{
		Command: command,
		Options: strs[1:],
	}
}
