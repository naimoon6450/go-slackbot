package githubclient

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"github.com/naimoon6450/go-slackbot/config"
	"golang.org/x/oauth2"
)

type Client struct {
	*github.Client
	GrowthMembers []string
	GrowthTeamID  int
}

// TODO: make things less hardcoded
func (c *Client) BuildPRMessage(ctx context.Context) (string, error) {
	daysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	prListStr := "*[Open Growth PRs]*\n\n"
	for _, member := range c.GrowthMembers {
		opts := &github.SearchOptions{
			ListOptions: github.ListOptions{PerPage: 20},
		}
		query := fmt.Sprintf("is:pr is:open draft:false org:%s author:%s created:>%s", "UseFedora", member, daysAgo)
		prs, _, err := c.Search.Issues(ctx, query, opts)
		if err != nil {
			return "", err
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

	return prListStr, nil
}

func New(ctx context.Context, cfg config.Github) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Client{
		client,
		cfg.GrowthMembers,
		cfg.GrowthTeamID,
	}
}
