package slackclient

import (
	"context"
	"log"

	"github.com/naimoon6450/go-slackbot/githubclient"
	"github.com/slack-go/slack"
)

func HandleSlashCommand(ctx context.Context, cmd slack.SlashCommand, client *Client, ghc *githubclient.Client) error {
	switch cmd.Command {
	case "/pulls":
		return handlePullsCmd(ctx, cmd, client, ghc)
	default:
		return nil
	}
}

func handlePullsCmd(ctx context.Context, cmd slack.SlashCommand, client *Client, ghc *githubclient.Client) error {
	prListStr, err := ghc.BuildPRMessage(ctx)
	if err != nil {
		return err
	}

	_, resp, err := client.PostMessage(cmd.ChannelID, slack.MsgOptionText(prListStr, false))

	if err != nil {
		return err
	}

	log.Println(resp)

	return nil
}

func HelperSlackPost(ctx context.Context, client *Client, ghc *githubclient.Client, channelID string) error {
	prListStr, err := ghc.BuildPRMessage(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	_, resp, err := client.PostMessage(channelID, slack.MsgOptionText(prListStr, false))

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(resp)

	return nil
}
