package slackclient

import (
	"github.com/naimoon6450/go-slackbot/config"
	"github.com/slack-go/slack"
)

type Client struct {
	*slack.Client
}

func New(cfg config.Config) *Client {
	client := slack.New("token")
	return &Client{
		Client: client,
	}
}
