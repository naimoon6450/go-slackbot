package githubclient

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/naimoon6450/go-slackbot/config"
	"golang.org/x/oauth2"
)

type Client struct {
	*github.Client
	GrowthMembers []string
	GrowthTeamID  int
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
