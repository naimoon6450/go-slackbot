package slackclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/github"
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

	daysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	text := slack.TextBlockObject{
		Type: "mrkdwn",
	}

	prListStr := "*[Open Growth PRs]*\n\n"
	for _, member := range ghc.GrowthMembers {
		opts := &github.SearchOptions{
			ListOptions: github.ListOptions{PerPage: 20},
		}
		query := fmt.Sprintf("is:pr is:open draft:false org:%s author:%s created:>%s", "UseFedora", member, daysAgo)
		prs, _, err := ghc.Search.Issues(ctx, query, opts)
		if err != nil {
			log.Fatal(err)
		}

		if len(prs.Issues) == 0 {
			continue
		}

		memberPRStr := fmt.Sprintf("*%s*\n", member)
		for _, pr := range prs.Issues {
			loc, _ := time.LoadLocation("America/New_York")
			formattedTime := pr.GetCreatedAt().In(loc).Format("2006-01-02 03:04:05 PM")
			memberPRStr += fmt.Sprintf("%s | %s\n", pr.GetHTMLURL(), formattedTime)
		}

		prListStr += memberPRStr
	}

	text.Text = prListStr

	_, resp, err := client.PostMessage(cmd.ChannelID, slack.MsgOptionText(prListStr, false))

	if err != nil {
		return err
	}

	log.Println(resp)

	return nil
}
