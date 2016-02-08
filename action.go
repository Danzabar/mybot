package main

// A single action
type Action struct {
	name    string
	keyword string
	perform func(b *Bot, m Message)
}

// An action dictionary
type ActionDictionary struct {
	actions []Action
	bot     *Bot
}

func (a Action) run(b *Bot, m Message) {
	go a.perform(b, m)
}

// Adds an action to the dictionary
func (d *ActionDictionary) addAction(a Action) {
	d.actions = append(d.actions, a)
}

// Adds the actions to the bot
func (d *ActionDictionary) addToBot() {
	// Set up the map on the bot struct
	d.bot.actions = make(map[string]Action, len(d.actions))

	for _, a := range d.actions {
		// Add keywords
		d.bot.keywords = append(d.bot.keywords, a.keyword)
		// Add to action map
		d.bot.actions[a.keyword] = a
	}
}
