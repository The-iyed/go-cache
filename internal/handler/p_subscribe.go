package handler

import (
	"fmt"
	"net"
)

func handlePatternUnsubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Invalid PUNSUBSCRIBE command\n"))
		return
	}
	pattern := command[1]
	pubSubStore.UnsubscribePattern(pattern, nil)
	conn.Write([]byte(fmt.Sprintf("Unsubscribed from pattern: %s\n", pattern)))
}
