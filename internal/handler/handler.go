package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/go-redis-v1/internal/store"
)

func HandleConnection(conn net.Conn, kvStore *store.KeyValueStore) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected!")
			return
		}

		command := strings.Fields(strings.TrimSpace(input))
		if len(command) < 1 {
			conn.Write([]byte("Invalid command\n"))
			continue
		}

		handleCommand(conn, kvStore, command)
	}
}

func handleCommand(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	switch strings.ToUpper(command[0]) {
	case "SET":
		handleSet(conn, kvStore, command)
	case "SETEX":
		handleSetEX(conn, kvStore, command)
	case "GET":
		handleGet(conn, kvStore, command)
	case "DEL":
		handleDel(conn, kvStore, command)
	case "KEYS":
		handleKeys(conn, kvStore, command)
	case "EXISTS":
		handleExists(conn, kvStore, command)
	case "TTL":
		handleTTL(conn, kvStore, command)
	case "FLUSHALL":
		handleFlushAll(conn, kvStore)
	case "INFO":
		handleInfo(conn, kvStore)
	case "PING":
		handlePing(conn, kvStore)
	case "PERSIST":
		handlePersist(conn, kvStore, command)
	case "EXPIRE":
		handleExpire(conn, kvStore, command)
	case "MSET":
		handleMSet(conn, kvStore, command)
	case "MGET":
		handleMGet(conn, kvStore, command)
	case "UPDATE":
		handleUpdate(conn, kvStore, command)
	case "GETSET":
		handleGetSet(conn, kvStore, command)
	default:
		conn.Write([]byte("Unknown command\n"))
	}
}

