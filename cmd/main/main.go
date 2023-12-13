package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sandy-sunday/kyuutei-kenetsukan/config"
)

func replaceToVx(s *discordgo.Session, m *discordgo.MessageCreate, url string) {
	if !strings.Contains(m.Content, url) {
		return
	}
	newContent := strings.ReplaceAll(m.Content, url, "https://vxtwitter.com/")
	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.ChannelMessageSend(m.ChannelID, newContent)
}

func main() {
	token, err := config.GetBotToken()
	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New(token)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)

	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		log.Printf("Message: %v", m.Content)
		// Stateuser is the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}
		replaceToVx(s, m, "https://twitter.com/")
		replaceToVx(s, m, "https://x.com/")
	})

	discord.StateEnabled = true
	discord.Identify.Intents = discordgo.IntentsGuildMessages
	discord.Open()
	log.Println("Opened")

	<-make(chan struct{})
}
