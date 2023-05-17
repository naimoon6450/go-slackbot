package main

import (
	"github.com/naimoon6450/go-slackbot/config"
	"github.com/naimoon6450/go-slackbot/internal/slackclient"
)

func main() {
	cfg := config.Config{}
	slackClient := slackclient.New(cfg)
}
