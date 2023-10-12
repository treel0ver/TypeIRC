// messageparser.go

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	irc "github.com/thoj/go-ircevent"
)

type MessageParser struct {
}

func NewMessageParser() *MessageParser {
	return &MessageParser{}
}

func (mp *MessageParser) ParseMessage(ircConn *irc.Connection, channel string, message string, sender string) {
	if strings.HasPrefix(message, "1") {
		numbers := []string{"a", "b", "c"}
		go func() {
			for _, v := range numbers {
				ircConn.Privmsg(channel, v)
				time.Sleep(1 * time.Second)
			}
		}()
	}

	if strings.HasPrefix(message, "2") {
		numbers := []string{"萡", "袈", "睳"}
		go func() {
			for _, v := range numbers {
				ircConn.Privmsg(channel, v)
				time.Sleep(1 * time.Second)
			}
		}()
	}

	if strings.HasPrefix(message, ".code ") {
		dataToEncode := []byte(message[6:])
		go func() {
			binaryEncoded := fmt.Sprintf("%08b", dataToEncode)
			hexEncoded := hex.EncodeToString(dataToEncode)
			base64Encoded := base64.StdEncoding.EncodeToString(dataToEncode)
			ircConn.Privmsg(channel, "binary: "+binaryEncoded+" | hexadecimal: "+hexEncoded+" | base64: "+base64Encoded)
		}()
	}

	if strings.HasPrefix(message, ".ts") {
		rand.Seed(time.Now().UnixNano())
		newtest := typingtest.test{
			textID:         rand.Intn(2),
			channel:        channel,
			sender:         sender,
			start:          time.Now(),
			users_finished: []string{""},
		}
		go func() {
			typingtest.start_test(channel, ircConn, &newtest)
		}()
	}
}
