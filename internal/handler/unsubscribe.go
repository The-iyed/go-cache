package handler

import "net"

func handleUnsubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: UNSUBSCRIBE <channel>\n"))
		return
	}
	channel := command[1]
	messageChannel := make(chan string)
	pubSubStore.Unsubscribe(channel, messageChannel)

	close(messageChannel)
	conn.Write([]byte("Unsubscribed from channel: " + channel + "\n"))
}
