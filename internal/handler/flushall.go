package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleFlushAll(conn net.Conn, kvStore *store.KeyValueStore) {
	kvStore.FlushAll()
	conn.Write([]byte("OK\n"))
}
