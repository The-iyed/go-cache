package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func HandleGet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: GET <key>\n"))
		return
	}
	value, exist := kvStore.Get(command[1])
	if !exist {
		conn.Write([]byte("(nil)\n"))
	} else {
		conn.Write([]byte(value + "\n"))
	}
}
