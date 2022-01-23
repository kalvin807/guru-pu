package command

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, c *Command) {
	s.ChannelMessageSend(m.ChannelID, "pong")
}
