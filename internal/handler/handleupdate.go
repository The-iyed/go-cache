package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleUpdate(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: UPDATE <key> <value>\n"))
		return
	}
	oldValue := kvStore.Update(command[1], command[2])
	if oldValue == "" {
		conn.Write([]byte("Key does not exist\n"))
	} else {
		conn.Write([]byte("OK\n"))
	}
}
