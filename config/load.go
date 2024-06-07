package config

import (
    "github.com/joho/godotenv"
    "os"
)

func Load() Config {
    err := godotenv.Load(".env")
    if err != nil {
        return Config{}
    }

    return Config{
        Discord{
            GuildID: os.Getenv("GUILD_ID"),
            Token:   os.Getenv("BOT_TOKEN"),
        },
    }

}
