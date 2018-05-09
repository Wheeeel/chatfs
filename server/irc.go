package server

import (
	ircclient "github.com/fluffle/goirc/client"
	"github.com/pkg/errors"
)

type IRC struct {
	client   *ircclient.Conn
	recvChan chan Message
	sendChan chan Message
	server   string
}

func NewIRC(nickname, server string) (irc *IRC) {
	irc = new(IRC)
	irc.client = ircclient.SimpleClient(nickname)
	irc.server = server
	irc.recvChan = make(chan Message, 1024)
	irc.sendChan = make(chan Message, 1024)
	return
}

func (irc *IRC) Join(c string) {
	irc.client.Join(c)
	return
}

func (irc *IRC) RecvChan() (msgChan <-chan Message) {
	msgChan = irc.recvChan
	return
}

func (irc *IRC) SendChan() (msgChan chan<- Message) {
	msgChan = irc.sendChan
	return
}

func (irc *IRC) handleRecv(conn *ircclient.Conn, line *ircclient.Line) {
	var msg Message
	msg.Nick = line.Nick
	msg.Time = line.Time
	msg.Text = line.Text()

	irc.recvChan <- msg
}

func (irc *IRC) handleSend() {
Out:
	for {
		select {
		case m, ok := <-irc.sendChan:
			if !ok {
				break Out
			}
			irc.client.Privmsg(m.Nick, m.Text)
		}
	}
}

func (irc *IRC) Start() (err error) {
	err = irc.client.ConnectTo(irc.server)
	if err != nil {
		err = errors.Wrap(err, "irc start error")
		return
	}

	irc.client.HandleFunc(ircclient.PRIVMSG, irc.handleRecv)
	irc.client.HandleFunc(ircclient.NOTICE, irc.handleRecv)
	go func() {
		irc.handleSend()
	}()
	return
}

func (irc *IRC) Close() (err error) {
	err = irc.client.Close()
	if err != nil {
		err = errors.Wrap(err, "irc close error")
		return
	}
	close(irc.recvChan)
	close(irc.sendChan)
	return
}

func (irc *IRC) Me() string {
	return irc.client.Me().Nick
}
