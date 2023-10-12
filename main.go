// main.go

package main

import (
	"crypto/tls"
	"fmt"

	irc "github.com/thoj/go-ircevent"
	// Import your typingtest package
)

func main() {
	ircConn := irc.IRC("TypeX__________", "TypeX__________")
	ircConn.UseTLS = true

	ircConn.TLSConfig = &tls.Config{
		ServerName: "irc.libera.chat",
	}

	err := ircConn.Connect("irc.libera.chat:6697")
	if err != nil {
		fmt.Println("Error connecting to IRC server:", err)
		return
	}

	ircConn.AddCallback("001", func(e *irc.Event) {
		ircConn.Join("#TypeX")
	})

	messageParser := NewMessageParser()

	ircConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		message := event.Message()
		channel := event.Arguments[0]
		sender := event.Nick

		messageParser.ParseMessage(ircConn, channel, message, sender)
	})

	ircConn.Loop()
}
