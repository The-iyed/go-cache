package handler

import (
	"net"
)

func HandleSubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: SUBSCRIBE <channel>\n"))
		return
	}
	channel := command[1]
	messageChannel := make(chan string)
	pubSubStore.Subscribe(channel)

	go func() {
		for message := range messageChannel {
			conn.Write([]byte(message + "\n"))
		}
	}()

	conn.Write([]byte("Subscribed to channel: " + channel + "\n"))
}
