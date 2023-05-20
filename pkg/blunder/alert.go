package blunder

import (
	"github.com/DenChenn/blunder/pkg/options"
	"github.com/bwmarrin/discordgo"
	"os"
)

func (mr *MatchResult) Alert(opt *options.Alert) *MatchResult {
	if opt != nil {
		switch true {
		case opt.DiscordAlert:
			if err := SendDiscordAlert(mr.Result, nil); err != nil {
				return &MatchResult{
					IsMatched: false,
					Result:    ErrDiscordAlertFailed.WithCustomMessage(err.Error()),
				}
			}
		}
	}
	return mr
}

func (mr *MatchResult) AlertWithLog(log string, opt *options.Alert) *MatchResult {
	if opt != nil {
		switch true {
		case opt.DiscordAlert:
			if err := SendDiscordAlert(mr.Result, &log); err != nil {
				return &MatchResult{
					IsMatched: false,
					Result:    ErrDiscordAlertFailed.WithCustomMessage(err.Error()),
				}
			}
		}
	}
	return mr
}

func SendDiscordAlert(happened Error, log *string) error {
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	channelId := os.Getenv("DISCORD_CHANNEL_ID")

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return err
	}
	defer discord.Close()

	var fd []*discordgo.MessageEmbedField
	fd = append(fd, &discordgo.MessageEmbedField{
		Name:  "Error ID",
		Value: happened.GetId(),
	})

	fd = append(fd, &discordgo.MessageEmbedField{
		Name:  "Message",
		Value: happened.GetMessage(),
	})

	if log != nil {
		fd = append(fd, &discordgo.MessageEmbedField{
			Name:  "Log",
			Value: *log,
		})
	}

	if _, err = discord.ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
		Title:  "🚒Bad things happened 🚒",
		Color:  0xff0000,
		Fields: fd,
	}); err != nil {
		return err
	}

	return nil
}
