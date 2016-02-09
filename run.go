package main

import (
	"fmt"
	"log"
	"os"
)

// The Ramsey Bot
var ramsey = &Bot{
	defMsg: "What are you? An idiot sandwich",
}

// Default Action Dictionary
var actionDict = &ActionDictionary{
	bot: ramsey,
}

var wowAction = Action{
	name:    "Wow response",
	keyword: "wow",
	perform: func(b *Bot, m Message) {
		m.Text = "wow wow wow"
		postMessage(b.con, m)
	},
}

var nsfwAction = Action{
	name:    "Not Safe For Work",
	keyword: "nsfw",
	perform: func(b *Bot, m Message) {
		m.Text = ""
		postMessage(b.con, m)
	},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, id := slackConnect(os.Args[1])
	fmt.Println("mybot ready, ^C exits")

	ramsey.name = id
	ramsey.con = ws

	actionDict.addToBot()

	for {
		// read each incoming message
		m, err := getMessage(ws)

		if err != nil {
			log.Fatal(err)
		}

		ramsey.respond(m)
	}
}

func addActions() {
	actionDict.addAction(wowAction)
}
