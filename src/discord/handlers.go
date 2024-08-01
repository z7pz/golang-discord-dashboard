package discord

import "github.com/bwmarrin/discordgo"

func CreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "lesgo" {
		s.ChannelMessageSendReply(m.ChannelID, "Hi", m.Reference())
	}
}
