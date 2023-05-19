# A Slack bot created using Go!

This bot is a simple bot that can be used to send messages to a Slack channel. It is written in Go and uses the [Slack API](https://api.slack.com/).

Currently, the cmd/growthbot will send a message containing the PRs of a given list of github users from the last 7 days. It will post a message daily at 9am and 4pm.