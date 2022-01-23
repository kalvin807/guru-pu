package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kalvin807/guru-pu/pkg/command"
)

type commandHandler func(s *discordgo.Session, m *discordgo.MessageCreate, c *command.Command)

var handlerMaps map[string]commandHandler = map[string]commandHandler{
	"ping": command.HandlePing,
}

func ProcessCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := command.NewCommand(m.Content)

	// Case empty command
	if cmd.Command == "" {
		// TODO: handle error
		return
	}

	handler, ok := handlerMaps[cmd.Command]

	// Case, unknown command
	if !ok {
		s.ChannelMessageSend(m.ChannelID, "what7usay")
		return
	}

	handler(s, m, cmd)
}
