package messagehandler

import (
    "github.com/flytegg/plugin-portal-bot/config"
)

type Handler struct {
    config  *config.Discord
    actions []func()
}

func New(cfg *config.Discord) *Handler {
    return &Handler{
        config: cfg,
    }
}
