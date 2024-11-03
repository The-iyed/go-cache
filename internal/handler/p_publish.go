package handler

import (
	"fmt"
	"net"
)

func HandlePatternSubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Invalid PSUBSCRIBE command\n"))
		return
	}
	pattern := command[1]
	ch := pubSubStore.SubscribePattern(pattern)

	go func() {
		for msg := range ch {
			conn.Write([]byte(fmt.Sprintf("Message: %s\n", msg)))
		}
	}()
	conn.Write([]byte(fmt.Sprintf("Subscribed to pattern: %s\n", pattern)))
}
