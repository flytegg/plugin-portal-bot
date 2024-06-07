package interactionhandler

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "log"
    "net/http"
    "strings"
)

func (h Handler) HandleButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {

    customID := i.MessageComponentData().CustomID
    fmt.Println("Button customID:", customID)

    action := strings.Split(customID, "_")[0]
    mongoId := strings.Split(customID, "_")[1]

    var url string

    switch action {
    case "merge":
        url = "http://localhost:8080/v1/duplicates/merge?id=" + mongoId
    case "cancel":
        url = "http://localhost:8080/v1/duplicates/cancel?id=" + mongoId
    default:
        log.Println("Unknown action:", action)
        return
    }

    res, err := http.Post(url, "application/json", nil)
    if err != nil || res.StatusCode != http.StatusOK {
        log.Printf("Failed to perform action %s on %s: %v", action, mongoId, err)
        return
    }

    responseContent := "Action " + action + " has been performed successfully."
    err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: responseContent,
        },
    })
    if err != nil {
        log.Println("Failed to send interaction response:", err)
    }
}
