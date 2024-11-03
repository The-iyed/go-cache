package handler

import (
	"fmt"
	"net"

	"github.com/go-redis-v1/internal/store"
)

func HandleTTL(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: TTL <key>\n"))
		return
	}
	ttl := kvStore.TTL(command[1])
	if ttl == -2 {
		conn.Write([]byte("(nil)\n"))
	} else {
		conn.Write([]byte(fmt.Sprintf("(integer) %d\n", ttl)))
	}
}
