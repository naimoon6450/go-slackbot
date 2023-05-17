package main

import (
	"github.com/naimoon6450/go-slackbot/config"
	"github.com/naimoon6450/go-slackbot/slackclient"
)

func main() {
	cfg := config.Config{}
	_ = slackclient.New(cfg)
}
