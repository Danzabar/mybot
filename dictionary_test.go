package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItReadsMessagesFromFile(t *testing.T) {
	d := NewDict("./test/testgreetings.json")

	assert.Equal(t, 3, len(d.messages))
}
