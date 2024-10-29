package handler

import (
	"net"
	"strconv"
)

func handleGetNumSub(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: GETNSUB <channel> \n"))
		return
	}
	channel := command[1]
	number := pubSubStore.GetNumSubscribers(channel)
	value := strconv.Itoa(number)
	conn.Write([]byte(value + "\n"))
}
