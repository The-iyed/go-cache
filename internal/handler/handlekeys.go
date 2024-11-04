package handler

import (
	"fmt"
	"net"
	"strings"

	"github.com/go-redis-v1/internal/store"
)

func HandleKeys(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: KEYS <pattern>\n"))
		return
	}
	keys := kvStore.Keys(command[1])
	var response strings.Builder
	for i, key := range keys {
		response.WriteString(fmt.Sprintf("%d) \"%s\"\n", i+1, key))
	}
	conn.Write([]byte(response.String()))
}
