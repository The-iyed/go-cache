package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleExists(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: EXISTS <key>\n"))
		return
	}
	exist := kvStore.Exist(command[1])
	if exist {
		conn.Write([]byte("(integer) 1\n"))
	} else {
		conn.Write([]byte("(integer) 0\n"))
	}
}
