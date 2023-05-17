package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Slack Slack
}

type Slack struct {
	AuthToken string
	ChannelID string
	AppToken  string
}

func New() Config {
	// load env vars
	godotenv.Load()

	return Config{
		Slack: Slack{
			AppToken:  os.Getenv("SLACK_APP_TOKEN"),
			AuthToken: os.Getenv("SLACK_AUTH_TOKEN"),
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
		},
	}
}
