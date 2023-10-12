package main

import (
	"fmt"
	"time"

	irc "github.com/thoj/go-ircevent"
)

// BEGIN to be continued

var texts = []string{"hehe", "123"}

type test struct {
	textID         int
	channel        string
	sender         string
	start          time.Time
	users_finished []string
}

func start_test(channel string, ircConn *irc.Connection, test *test) {
	ircConn.Privmsg(channel, fmt.Sprintf("textID: %d sender: %s ", test.textID, test.sender))
}

// END to be continued
