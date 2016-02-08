package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that actions can be added to the bot
func TestActionsCanBeAddedToBot(t *testing.T) {
	a := Action{
		name:    "TestAction",
		keyword: "wow",
		perform: func(b *Bot, m Message) {
			postMessage(b.con, m)
		},
	}

	b := &Bot{name: "TestBot"}

	d := &ActionDictionary{bot: b}
	d.addAction(a)
	d.addToBot()

	assert.Equal(t, 1, len(b.actions))
}
