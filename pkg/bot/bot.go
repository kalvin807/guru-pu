package bot

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
)

type Bot struct {
	Session *discordgo.Session
	Ctx     *context.Context
	Rdb     *redis.Client
}

func (b *Bot) onMessageReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	// Ignore all messages created by the bot itself
	if e.UserID == s.State.User.ID {
		return
	}
	HandleReactionAdd(b, e)
}

func (b *Bot) onMessageReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	// Ignore all messages created by the bot itself
	if e.UserID == s.State.User.ID {
		return
	}
	HandleReactionRemove(b, e)
}

func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	msg := m.Content
	if strings.HasPrefix(msg, "!gurupu:") {
		ProcessCommand(b, m)
	}
}

func (b *Bot) Run() error {
	// Register handlers
	b.Session.AddHandler(b.onMessageCreate)
	b.Session.AddHandler(b.onMessageReactionAdd)
	b.Session.AddHandler(b.onMessageReactionRemove)

	// Register intents
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions

	// Listen to discord websocket
	err := b.Session.Open()
	if err != nil {
		return err
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	return nil
}

func (b *Bot) Stop() {
	b.Session.Close()
	b.Rdb.Close()
}

func MakeBot(s *discordgo.Session, ctx *context.Context, rdb *redis.Client) *Bot {
	return &Bot{
		Session: s,
		Ctx:     ctx,
		Rdb:     rdb,
	}
}

func ProcessCommand(b *Bot, m *discordgo.MessageCreate) {
	s := b.Session
	cmd := NewCommand(m.Content)

	log.Println("got command", cmd.Command, cmd.Options)

	// Case empty command
	if cmd.Command == "" {
		// TODO: handle error, ignore for now
		return
	}

	handler, ok := handlerMaps[cmd.Command]

	// Case, unknown command
	if !ok {
		s.ChannelMessageSend(m.ChannelID, "what7usay")
		return
	}

	handler(b, m, cmd)
}
