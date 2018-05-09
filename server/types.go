package server

import "time"

type Message struct {
	Nick string    `json:"nick"`
	Time time.Time `json:"time"`
	Text string    `json:"text"`
}
