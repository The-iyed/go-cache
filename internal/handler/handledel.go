package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleDel(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: DEL <key>\n"))
		return
	}
	kvStore.Delete(command[1])
	conn.Write([]byte("OK\n"))
}