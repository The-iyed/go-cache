package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func HandlePing(conn net.Conn, kvStore *store.KeyValueStore) {
	ping := kvStore.Ping()
	conn.Write([]byte(ping + "\n"))
}
