package main

import (
	"golang.org/x/net/websocket"
	"strings"
)

// Generic Bot Struct
type Bot struct {
	name     string
	keywords []string
	message  Message
	con      websocket.Conn
}

// Create a new bot
func NewBot(con websocket.Conn, name string) {
	return &Bot{
		name,
	}
}

func (b *Bot) respond(m Message) {
	b.message = m

	// We only want to respond to messages
	if m.Type == "message" {

		if b.isMentioned || b.hasKeyword {

			if b.actionKey == nil {
				return
			}

			// Here we have an action key
		}
	}
}

// Checks whether the bot is mentioned personally
func (b *Bot) isMentioned(m Message) bool {
	return strings.Contains(m.Text, "@"+b.name)
}

// Checks whether there is an action/keyword in the message
func (b *Bot) hasKeyword(m Message) bool {
	for _, keyword := range b.keywords {
		if strings.Contains(m.Text, keyword) {
			b.actionKey = keyword
			return true
		}
	}
	return false
}
