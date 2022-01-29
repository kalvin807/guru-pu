package bot

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePing(b *Bot, m *discordgo.MessageCreate, c *Command) {
	b.Session.ChannelMessageSend(m.ChannelID, "pong")
}
