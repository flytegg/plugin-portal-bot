package config

type Discord struct {
    GuildID string
    Token   string
}

type Config struct {
    Discord Discord
}

func New(discord Discord) Config {
    return Config{
        Discord: discord,
    }
}
