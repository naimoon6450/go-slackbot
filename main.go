package main

import (
	"context"
	"log"

	"github.com/naimoon6450/go-slackbot/config"
	"github.com/naimoon6450/go-slackbot/slackclient"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	cfg := config.New()
	sc := slackclient.New(cfg.Slack)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go socketListener(ctx, sc, &sc.SocketClient)

	sc.SocketClient.Run()
}

func socketListener(ctx context.Context, client *slackclient.Client, sc *slackclient.SocketClient) {
	// We'll be listening for events on the socket and blocking until we get an event
	for {
		select {
		case <-ctx.Done():
			log.Println("closing socket listener")
			return
		case event := <-sc.Events:
			switch event.Type {

			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Printf("Ignored %+v\n", event)
					continue
				}

				sc.Ack(*event.Request)

				err := slackclient.HandleSlackEvents(eventsAPIEvent, client)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
