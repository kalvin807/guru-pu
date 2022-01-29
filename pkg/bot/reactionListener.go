package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
)

func HandleReactionAdd(b *Bot, e *discordgo.MessageReactionAdd) {
	chanIDRecord, err := b.Rdb.Get(*b.Ctx, e.GuildID).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("HandleReactionAdd", err)
	}
	// ignore if message is not in registered chan
	if err == redis.Nil || chanIDRecord != e.ChannelID {
		return
	}
	// ignore if the message does not register to a roleID
	roleIDRecord, err := b.Rdb.Get(*b.Ctx, e.MessageID).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("HandleReactionAdd", err)
	}
	if err == redis.Nil {
		return
	}

	// Add this role to the sender
	if err := b.Session.GuildMemberRoleAdd(e.GuildID, e.UserID, roleIDRecord); err != nil {
		log.Fatalln("HandleReactionAdd", err)
	}
}

func HandleReactionRemove(b *Bot, e *discordgo.MessageReactionRemove) {
	chanIDRecord, err := b.Rdb.Get(*b.Ctx, e.GuildID).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("HandleReactionRemove", err)
	}
	// ignore if message is not in registered chan
	if err == redis.Nil || chanIDRecord != e.ChannelID {
		return
	}
	// ignore if the messageID was not registered in redis (not created by this bot)
	roleIDRecord, err := b.Rdb.Get(*b.Ctx, e.MessageID).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("HandleReactionRemove", err)
	}
	if err == redis.Nil {
		return
	}
	// remove this role from the sender
	if err := b.Session.GuildMemberRoleRemove(e.GuildID, e.UserID, roleIDRecord); err != nil {
		log.Fatalln("HandleReactionRemove", err)
	}
}
