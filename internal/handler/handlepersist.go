package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)


func handlePersist(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: PERSIST <key>\n"))
		return
	}
	if kvStore.Persist(command[1]) {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("(nil)\n"))
	}
}