package handler

import (
	"net"
)

func handlePublish(conn net.Conn, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: PUBLISH <channel> <message>\n"))
		return
	}
	channel := command[1]
	message := command[2]
	pubSubStore.Publish(channel, message)
	conn.Write([]byte("Message published\n"))
}
