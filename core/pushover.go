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

// SendPushoverMessage is used to send a Pushover notification.
func SendPushoverMessage(token string, user string, message string, asOf time.Time) (err error) {
	var resp *resty.Response
	var poReply types.PushoverReply

	client := resty.New()
	resp, err = client.R().
		EnableTrace().
		SetFormData(map[string]string{
			"token":     token,
			"user":      user,
			"message":   message,
			"title":     "Coinspy",
			"sound":     "cashregister",
			"url":       ToolURL,
			"timestamp": strconv.FormatInt(asOf.Unix(), 10),
		}).
		Post("https://api.pushover.net/1/messages.json")
	if err != nil {
		err = fmt.Errorf("could not push message: %s", err)
		return
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
