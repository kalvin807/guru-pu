package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleRegisterChannel(b *Bot, m *discordgo.MessageCreate, c *Command) {
	// upsert to redis
	err := b.Rdb.Set(*b.Ctx, m.GuildID, m.ChannelID, 0).Err()
	if err != nil {
		log.Fatalln("registerChannel", err)
	}
	// add emoji to the message
	b.Session.MessageReactionAdd(m.ChannelID, m.Message.ID, "👀")
}
