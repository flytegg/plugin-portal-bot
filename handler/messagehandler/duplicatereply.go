package messagehandler

import (
    "github.com/bwmarrin/discordgo"
    "log"
    "net/http"
)

func (h Handler) DuplicateEmbedReply(s *discordgo.Session, msg *discordgo.MessageCreate) {
    if msg.Author.ID == s.State.User.ID {
        return
    }

    if len(msg.Embeds) == 0 || len(msg.Embeds[0].Fields) < 3 {
        log.Println("Message does not have enough fields to extract mongoId")
        return
    }

    var mongoId = msg.Embeds[0].Fields[2].Value

    res, err := http.Get("http://localhost:8080/v1/duplicates?id=" + mongoId)
    if err != nil {
        return
    }

    var embed = discordgo.MessageEmbed{
        Description: "Would you like to merge this duplicate?",
        Color:       0x2B2D31,
    }

    if res.StatusCode == 200 {
        buttons := []discordgo.MessageComponent{
            discordgo.ActionsRow{
                Components: []discordgo.MessageComponent{
                    discordgo.Button{
                        Label:    "Merge",
                        Emoji:    discordgo.ComponentEmoji{Name: "🔗"},
                        Style:    discordgo.SecondaryButton,
                        CustomID: "merge_" + mongoId, // CustomID for identifying the button click
                    },
                    discordgo.Button{
                        Label:    "Cancel",
                        Emoji:    discordgo.ComponentEmoji{Name: "❌"},
                        Style:    discordgo.SecondaryButton,
                        CustomID: "cancel_" + mongoId, // CustomID for identifying the button click
                    },
                },
            },
        }

        _, err := s.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
            Embed:      &embed,
            Components: buttons,
            Reference:  msg.Reference(),
        })
        if err != nil {
            log.Println(err)
        }
    }
}
