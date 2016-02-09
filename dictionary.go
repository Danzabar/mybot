package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// The Dictionary Struct, provides randomized messages
type Dictionary struct {
	file     string
	messages []string
}

// Creates a new Dictionary and loads from file
func NewDict(file string) *Dictionary {
	d := &Dictionary{
		file: file,
	}

	d.readFromFile()

	return d
}

// Reads the json file associated with the dictionary
func (d *Dictionary) readFromFile() {
	file, err := ioutil.ReadFile(d.file)

	if err != nil {
		fmt.Printf("Unable to fetch dictionary: %v", err)
		os.Exit(1)
	}

	var messages []string

	json.Unmarshal(file, &messages)

	for _, mes := range messages {
		d.messages = append(d.messages, mes)
	}
}

// Selects a random entry from the list
func (d *Dictionary) pick() {
}
