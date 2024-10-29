package handler

import (
	"net"
	"strings"

	"github.com/go-redis-v1/internal/store"
)

func handleMGet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) < 2 {
		conn.Write([]byte("Usage: MGET <key1> [<key2> ...]\n"))
		return
	}
	values := kvStore.MGET(command[1:]...)
	response := strings.Join(values, "\n") + "\n"
	conn.Write([]byte(response))
}
