package handler

import (
	"net"

	"github.com/go-redis-v1/internal/store"
)

func handleInfo(conn net.Conn, kvStore *store.KeyValueStore) {
	info := kvStore.Info()
	conn.Write([]byte(info + "\n"))
}
