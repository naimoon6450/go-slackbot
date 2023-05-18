package slackclient

import (
	"errors"
	"math/rand"
	"time"

	"github.com/slack-go/slack/slackevents"
)

// Fake list of users, populate with users from channel
// var (
// 	ListOfUsers = []string{
// 		"U01J9QGQZ9Z",
// 		"U01J9QGQZ9Z",
// 		"U01J9QGQZ9Z",
// 		"U01J9QGQZ9Z",
// 	}
// )

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
	// text := strings.ToLower(event.Text)
	// userID := event.User
	// // userID := randomUserFromList(ListOfUsers)
	// eventTS := event.TimeStamp
	// threadTS := event.ThreadTimeStamp

	// if strings.Contains(text, "/pull/") {
	// 	// tag a random user within a thread to that message
	// 	_, _, err := client.PostMessage(
	// 		event.Channel,
	// 		slack.MsgOptionText("<@"+userID+">", false),
	// 		slack.MsgOptionTS(eventTS),
	// 	)

	// 	if err != nil {
	// 		return err
	// 	}
	// } else if strings.Contains(text, "bump") && threadTS != "" {
	// 	// tag a random user within a thread to that message
	// 	threadTS := event.TimeStamp
	// 	_, _, err := client.PostMessage(
	// 		event.Channel,
	// 		slack.MsgOptionText("<@"+userID+">", false),
	// 		slack.MsgOptionTS(threadTS),
	// 	)

	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *Client) error {
	// do stuff with app mention event
	return nil
}

func randomUserFromList(list []string) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return list[r.Intn(len(list))]
}
