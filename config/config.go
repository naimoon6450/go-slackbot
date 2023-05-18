package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Slack  Slack
	Github Github
}

type Slack struct {
	AuthToken string
	ChannelID string
	AppToken  string
}

type Github struct {
	AccessToken   string
	GrowthMembers []string
	GrowthTeamID  int
}

func New() Config {
	godotenv.Load()

	growthMemStr := os.Getenv("GROWTH_MEMBERS")
	growthMembers := strings.Split(growthMemStr, ",")

	growthTeamIDToInt, _ := strconv.Atoi(os.Getenv("GROWTH_TEAM_ID"))

	return Config{
		Slack: Slack{
			AppToken:  os.Getenv("SLACK_APP_TOKEN"),
			AuthToken: os.Getenv("SLACK_AUTH_TOKEN"),
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
		},
		Github: Github{
			AccessToken:   os.Getenv("GITHUB_ACCESS_TOKEN"),
			GrowthMembers: growthMembers,
			GrowthTeamID:  growthTeamIDToInt,
		},
	}
}
