package main

import (
	"golang.org/x/net/websocket"
	"strings"
)

// Generic Bot Struct
type Bot struct {
	name      string
	keywords  []string
	message   Message
	actionKey string
	con       *websocket.Conn
	actions   map[string]Action
}

// Adds an action to the bot
func (b *Bot) addAction(a Action) {
	b.actions[a.keyword] = a
}

// Responds to a message
func (b *Bot) respond(m Message) {
	b.message = m

	// We only want to respond to messages
	if m.Type == "message" {

		// We were mentioned in a message
		if b.isMentioned(m) {

			// check if we also found a keyword in the message
			if b.hasKeyword(m) {
				b.actions[b.actionKey].run(b, m)
				return
			}

			// Here we have no keyword but we were mentioned
			return
		}

		// We found a keyword
		if b.hasKeyword(m) {
			// We have an action to perform
			b.actions[b.actionKey].run(b, m)
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
