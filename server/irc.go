package server

import (
	irc "gopkg.in/sorcix/irc.v2"
)

type IRC struct {
	client *irc.Conn
}

func NewIRC() (irc *IRC) {
	return &IRC{}
}

func (irc *IRC) Login() (err error) {
	return
}

func (irc *IRC) Logout() (err error) {
	return
}

func (irc *IRC) MessageChannel() (msg chan *Message, err error) {
	return
}
