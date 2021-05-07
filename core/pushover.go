package core

import (
	"fmt"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"
)

func SendPushoverMessage(token string, user string, message string, asOf time.Time) (err error) {
	client := resty.New()
	_, pushErr := client.R().
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
	if pushErr != nil {
		err = fmt.Errorf("could not push message: %s", pushErr)
	}
	return
}
