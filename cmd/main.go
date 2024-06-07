package main

import (
    "flag"
    "github.com/flytegg/plugin-portal-bot/config"
    "github.com/flytegg/plugin-portal-bot/delivery/websocketserver"
    "github.com/flytegg/plugin-portal-bot/handler/interactionhandler"
    "github.com/flytegg/plugin-portal-bot/handler/messagehandler"
    "github.com/joho/godotenv"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

var (
    cfg config.Config
)

func init() {
    cfg = config.Load()
    flag.Parse()
}

func main() {
    err := godotenv.Load()
    if err != nil {
        return
    }

    // add as many handlers as you want implementing websocketserver.Handler...
    handlers := []websocketserver.Handler{
        messagehandler.New(&cfg.Discord),
        interactionhandler.New(&cfg.Discord),
    }

    server := websocketserver.New(&cfg, handlers, discordgo.IntentsAll)
    server.Serve()
    defer server.Shutdown()

    // Wait here until CTRL-C or other term signal is received.
    log.Println("Bot is now running.  Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    <-sc
}
