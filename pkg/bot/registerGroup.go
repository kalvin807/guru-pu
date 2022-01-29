package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
)

func HandleRegisterGroup(b *Bot, m *discordgo.MessageCreate, c *Command) {
	val, err := b.Rdb.Get(*b.Ctx, m.GuildID).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("registerGroup", err)
	}
	// ignore if message is not in registered chan
	if err == redis.Nil || val != m.ChannelID {
		return
	}
	// ignore if option is empty
	if len(c.Options) == 0 {
		return
	}
	// role name will be the first option
	roleName := c.Options[0]
	roleID, err := createGuildRole(b.Session, m.GuildID, roleName)
	if err != nil {
		log.Fatalln(err)
	}

	if roleID != "" {
		err := b.Rdb.Set(*b.Ctx, m.ID, roleID, 0).Err()
		if err != nil {
			log.Fatalln("registerGroup", err)
		}
	}

	// add emoji to the if success
	b.Session.MessageReactionAdd(m.ChannelID, m.Message.ID, "👌")
}

func createGuildRole(s *discordgo.Session, guildID string, name string) (string, error) {

	allRoles, err := s.GuildRoles(guildID)
	// skip create if already exist
	oldRole := func(rs []*discordgo.Role) *discordgo.Role {
		for _, r := range rs {
			if r.Name == name {
				return r
			}
		}
		return nil
	}(allRoles)

	if oldRole != nil {
		return "", nil
	}

	defaultRole := func(rs []*discordgo.Role) *discordgo.Role {
		for _, r := range rs {
			if r.Name == "@everyone" {
				return r
			}
		}
		return nil
	}(allRoles)
	if defaultRole == nil || err != nil {
		log.Fatalln(err)
	}

	role, err := s.GuildRoleCreate(guildID)
	if err != nil {
		log.Fatalln(err)
	}

	// It inherent all things from default
	// except name, and it is always mentionable
	if _, err := s.GuildRoleEdit(guildID,
		role.ID,
		name,
		defaultRole.Color,
		defaultRole.Hoist,
		defaultRole.Permissions,
		true); err != nil {
		log.Fatalln(err)
	}

	return role.ID, nil
}
