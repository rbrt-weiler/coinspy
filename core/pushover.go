package core

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

// HTMLizeMessage adds HTML links to Yahoo finance charts to each line.
func HTMLizeMessage(message string) (htmlMessage string) {
	var urlTemplate string
	var lineTemplate string
	var currencyParser *regexp.Regexp
	var currencyParts []string
	var coin string
	var fiat string
	var url string

	urlTemplate = `https://finance.yahoo.com/quote/%s-%s/chart/`
	lineTemplate = `%s <a href="%s">&#x1f4c8</a>`
	currencyParser = regexp.MustCompile(`(\w+)\s*=\s*(\d+.\d+|\d+)\s*(\w+)`)

	for _, line := range strings.Split(message, LineBreak) {
		currencyParts = currencyParser.FindStringSubmatch(line)
		coin = currencyParts[1]
		fiat = currencyParts[3]
		url = fmt.Sprintf(urlTemplate, coin, fiat)
		line = fmt.Sprintf(lineTemplate, line, url)
		htmlMessage = strings.Join([]string{htmlMessage, line}, LineBreak)
	}

	return
}

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
	var htmlParam string
	var messages []string
	var urlTitle string
	var resp *resty.Response
	var poReply types.PushoverReply

	htmlParam = "0"
	if Config.Pushover.IncludeLinks {
		message = HTMLizeMessage(message)
		htmlParam = "1"
	}
	messages, err = SanitizeMessage(message)
	if err != nil {
		err = fmt.Errorf("could not sanitize Pushover message: %s", err)
		return
	}

	client := resty.New()
	for _, message = range messages {
		if Config.Pushover.IncludeHost {
			hostname, hostnameErr := os.Hostname()
			if hostnameErr != nil {
				hostname = "<undefined>"
			}
			urlTitle = fmt.Sprintf("sent by %s v%s via %s", ToolName, ToolVersion, hostname)
		} else {
			urlTitle = fmt.Sprintf("sent by %s v%s", ToolName, ToolVersion)
		}
		resp, err = client.R().
			SetFormData(map[string]string{
				"token":     token,
				"user":      user,
				"message":   message,
				"html":      htmlParam,
				"title":     "Coinspy",
				"sound":     "cashregister",
				"url":       ToolURL,
				"url_title": urlTitle,
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
