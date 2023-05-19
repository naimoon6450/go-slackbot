package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/naimoon6450/go-slackbot/config"
	"github.com/naimoon6450/go-slackbot/githubclient"
	"github.com/naimoon6450/go-slackbot/slackclient"
	"github.com/robfig/cron/v3"
)

func main() {
	ctx := context.Background()

	cfg := config.New()
	ghc := githubclient.New(ctx, cfg.Github)
	sc := slackclient.New(cfg.Slack)

	c := cron.New()

	// c.AddFunc("0 0 9,16 * * *", func() { slackclient.BuildPRMessage(ctx, ghc) })
	c.AddFunc("*/1 * * * *", func() {
		log.Println("Firing post to slack")
		slackclient.HelperSlackPost(ctx, sc, ghc, "SOMECHANNEL")
	})

	log.Println("Starting cron")
	c.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	c.Stop()
}
