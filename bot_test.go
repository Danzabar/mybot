package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// It should know when it has been mentioned by name
func TestItShouldKnowWhenItsBeenMentioned(t *testing.T) {
	m := Message{
		Text: "@TestBot",
	}

	b := &Bot{
		name: "TestBot",
	}

	res := b.isMentioned(m)
	assert.Equal(t, true, res)
}

// It should identify messages that do not mention
func TestItShouldIdentifyWhenItHasntBeenMentioned(t *testing.T) {
	m := Message{
		// Note, this doesnt *Mention* the bot directly.
		Text: "TestBot",
	}

	b := &Bot{
		name: "TestBot",
	}

	res := b.isMentioned(m)
	assert.Equal(t, false, res)
}

// It should identify keywords in a message
func TestItShouldIdentifyKeywordsInAMessage(t *testing.T) {
	m := Message{
		Text: "Some text wow keyword",
	}

	b := &Bot{
		name:     "TestBot",
		keywords: []string{"wow"},
	}

	res := b.hasKeyword(m)
	assert.Equal(t, true, res)
}
