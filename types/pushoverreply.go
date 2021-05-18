package types

// PushoverReply encapsulates an answer from the Pushover API.
type PushoverReply struct {
	User    string   `json:"user"`
	Errors  []string `json:"errors"`
	Status  int      `json:"status"`
	Request string   `json:"request"`
}
