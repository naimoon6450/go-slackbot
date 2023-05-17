package slackclient

import (
	"log"
	"os"

	"github.com/naimoon6450/go-slackbot/config"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Client struct {
	*slack.Client
	SocketClient SocketClient
}

type SocketClient struct {
	*socketmode.Client
}

func New(cfg config.Slack) *Client {
	client := slack.New(cfg.AuthToken, slack.OptionDebug(true), slack.OptionAppLevelToken(cfg.AppToken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(
			os.Stdout, "socketmode: ",
			log.Lshortfile|log.LstdFlags,
		)))

	return &Client{
		client,
		SocketClient{
			socketClient,
		},
	}
}
