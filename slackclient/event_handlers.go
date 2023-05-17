package slackclient

import (
	"errors"

	"github.com/slack-go/slack/slackevents"
)

func HandleSlackEvents(
	event slackevents.EventsAPIEvent,
	client *Client,
) error {
	switch event.Type {
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			return handleMessageEvent(ev, client)
		case *slackevents.AppMentionEvent:
			return handleAppMentionEvent(ev, client)
		}
	default:
		return errors.New("unknown event type")
	}

	return nil
}

func handleMessageEvent(event *slackevents.MessageEvent, client *Client) error {
	// do stuff with message event
	return nil
}

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *Client) error {
	// do stuff with app mention event
	return nil
}
