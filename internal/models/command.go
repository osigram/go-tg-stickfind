package models

import (
	"errors"
	"regexp"
	"strings"
)

type InputCommand struct {
	Command string
	Params  []string
}

func NewInputCommand(message string) (*InputCommand, error) {
	// Command has a structure like this: /command@bot param1 param2...
	if len(message) < 2 {
		return nil, errors.New("message is empty")
	}

	if message[0] != '/' || message[1] == ' ' {
		return nil, errors.New("message is not a command")
	}

	messageBySpace := strings.Split(message, " ")

	commandWithBotNick := messageBySpace[0]
	command := regexp.MustCompile(`[a-zA-Z0-9_]+`).FindString(commandWithBotNick)

	params := messageBySpace[1:]

	return &InputCommand{
		Command: command,
		Params:  params,
	}, nil
}
