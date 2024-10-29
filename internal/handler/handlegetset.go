package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleGetSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: GETSET <key> <value>\n"))
		return
	}
	oldValue := kvStore.GetSet(command[1], command[2])
	conn.Write([]byte(oldValue + "\n"))
}
