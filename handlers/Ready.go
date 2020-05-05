package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func ReadyHandler(s *discordgo.Session, c *discordgo.Connect) {
	_ = s.UpdateStatus(0, "game")
	log.Println("bot ready lol")
}

