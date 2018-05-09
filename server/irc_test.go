package server

import (
	"fmt"
	"testing"
	"time"
)

func TestIRC(t *testing.T) {
	irc := NewIRC("chatfs-test", "irc.freenode.net")

	err := irc.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer irc.Close()

	irc.Join("#test-channel")
	go func() {
		msgChan := irc.RecvChan()
		for msg := range msgChan {
			fmt.Println("get a message:", msg)
		}
	}()

	sendChan := irc.SendChan()
	var msg Message
	msg.Nick = "#test-channel"
	msg.Text = "hello world 123"
	sendChan <- msg

	time.Sleep(time.Second * 60)
}
