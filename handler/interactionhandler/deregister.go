package interactionhandler

import "github.com/bwmarrin/discordgo"

func (h Handler) DeRegister(session *discordgo.Session) {
    h.TearDownCommands(session)

    for _, remove := range h.handlers {
        remove()
    }
}
