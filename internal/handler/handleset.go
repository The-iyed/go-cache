package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func HandleSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: SET <key> <value>\n"))
		return
	}
	kvStore.Set(command[1], command[2], 0)
	conn.Write([]byte("OK\n"))
}
