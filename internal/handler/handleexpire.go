package handler

import (
	"net"
	"strconv"

	"github.com/go-redis-v1/internal/store"
)

func HandleExpire(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: EXPIRE <key> <seconds>\n"))
		return
	}
	seconds, err := strconv.Atoi(command[2])
	if err != nil {
		conn.Write([]byte("Invalid seconds\n"))
		return
	}
	if kvStore.Expire(command[1], seconds) {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("(nil)\n"))
	}
}
