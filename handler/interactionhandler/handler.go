package interactionhandler

import (
    "github.com/flytegg/plugin-portal-bot/config"
)

type Handler struct {
    config     *config.Discord
    env        *config.Discord
    handlers   []func() // we need handlers func in order to deregister them later
    commandIDs []string // we need command ids in order to deregister them later
}

func New(cfg *config.Discord) *Handler {
    return &Handler{
        config: cfg,
    }
}
