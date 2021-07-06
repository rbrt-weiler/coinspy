package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

// SanitizeMessage checks the length of a message and returns an array of strings, each containing a part that will go through Pushover without being cut off.
func SanitizeMessage(message string) (messageParts []string, err error) {
	var messageLength int
	var singleMessage string

	messageLength = len(message)
	if messageLength > PushoverMessageLength {
		for _, line := range strings.Split(message, LineBreak) {
			if (len(line) + len(singleMessage)) > PushoverMessageLength {
				messageParts = append(messageParts, strings.TrimSpace(singleMessage))
				singleMessage = ""
			}
			singleMessage = strings.Join([]string{singleMessage, line}, LineBreak)
		}
		messageParts = append(messageParts, strings.TrimSpace(singleMessage))
	} else {
		messageParts = append(messageParts, message)
	}

	return
}

// SendPushoverMessage is used to send a Pushover notification.
func SendPushoverMessage(token string, user string, message string, asOf time.Time) (err error) {
	var messages []string
	var resp *resty.Response
	var poReply types.PushoverReply

	messages, err = SanitizeMessage(message)
	if err != nil {
		err = fmt.Errorf("could not sanitize Pushover message: %s", err)
		return
	}

	client := resty.New()
	for _, message = range messages {
		resp, err = client.R().
			SetFormData(map[string]string{
				"token":     token,
				"user":      user,
				"message":   message,
				"title":     "Coinspy",
				"sound":     "cashregister",
				"url":       ToolURL,
				"url_title": fmt.Sprintf("sent by %s v%s", ToolName, ToolVersion),
				"timestamp": strconv.FormatInt(asOf.Unix(), 10),
			}).
			Post("https://api.pushover.net/1/messages.json")
		if err != nil {
			err = fmt.Errorf("could not push message: %s", err)
			return
		}
	}

	err = json.Unmarshal(resp.Body(), &poReply)
	if err != nil {
		err = fmt.Errorf("could no unmarshal API response: %s", err)
		return
	}
	if poReply.Status != 1 {
		err = fmt.Errorf("could not push message: %s", strings.Join(poReply.Errors, ": "))
	}

	return
}
