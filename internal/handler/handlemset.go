package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleMSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) < 3 || len(command)%2 == 0 {
		conn.Write([]byte("Usage: MSET <key1> <value1> [<key2> <value2> ...]\n"))
		return
	}
	kvStore.MSET(command[1:]...)
	conn.Write([]byte("OK\n"))
}
