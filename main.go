package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/kalvin807/guru-pu/pkg/bot"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	redisOption, err := redis.ParseURL("redis://default:bFijghCeZbxao7tS8Ugb@containers-us-west-12.railway.app:6976")
	if err != nil {
		fmt.Println("error opening connection,", err)
		panic(err)
	}
	rdb := redis.NewClient(redisOption)
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	bot := bot.MakeBot(discord, &ctx, rdb)
	bot.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Stop()
}
