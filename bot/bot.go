package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/germanpachec0/discord-bot/config"
)

type Bot struct {
}

func NewBot(cfg config.AppCredentials) (*discordgo.Session, error) {
	return discordgo.New("Bot " + cfg.Key)
}
