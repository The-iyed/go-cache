package handler

import (
	"net"
	"time"

	"github.com/go-redis-v1/internal/store"
)

func handleSetEX(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 4 {
		conn.Write([]byte("Usage: SETEX <key> <value> <ttl>\n"))
		return
	}
	ttl, err := time.ParseDuration(command[3] + "s")
	if err != nil {
		conn.Write([]byte("Invalid TTL\n"))
		return
	}
	kvStore.Set(command[1], command[2], ttl)
	conn.Write([]byte("OK\n"))
}
