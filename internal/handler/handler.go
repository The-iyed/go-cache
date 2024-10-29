package handler

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

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

func handleSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: SET <key> <value>\n"))
		return
	}
	kvStore.Set(command[1], command[2], 0)
	conn.Write([]byte("OK\n"))
}

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

func handleGet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: GET <key>\n"))
		return
	}
	value, exist := kvStore.Get(command[1])
	if !exist {
		conn.Write([]byte("(nil)\n"))
	} else {
		conn.Write([]byte(value + "\n"))
	}
}

func handleDel(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: DEL <key>\n"))
		return
	}
	kvStore.Delete(command[1])
	conn.Write([]byte("OK\n"))
}

func handleKeys(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
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

func handleExists(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: EXISTS <key>\n"))
		return
	}
	exist := kvStore.Exist(command[1])
	if exist {
		conn.Write([]byte("(integer) 1\n"))
	} else {
		conn.Write([]byte("(integer) 0\n"))
	}
}

func handleTTL(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
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

func handleFlushAll(conn net.Conn, kvStore *store.KeyValueStore) {
	kvStore.FlushAll()
	conn.Write([]byte("OK\n"))
}

func handleInfo(conn net.Conn, kvStore *store.KeyValueStore) {
	info := kvStore.Info()
	conn.Write([]byte(info + "\n"))
}

func handlePing(conn net.Conn, kvStore *store.KeyValueStore) {
	ping := kvStore.Ping()
	conn.Write([]byte(ping + "\n"))
}

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

func handleExpire(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: EXPIRE <key> <seconds>\n"))
		return
	}
	seconds, err := strconv.Atoi(command[2])
	if err != nil {
		conn.Write([]byte("Invalid seconds\n"))
		return
	}
	if kvStore.Expire(command[1], seconds) {
		conn.Write([]byte("OK\n"))
	} else {
		conn.Write([]byte("(nil)\n"))
	}
}

func handleMSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) < 3 || len(command)%2 == 0 {
		conn.Write([]byte("Usage: MSET <key1> <value1> [<key2> <value2> ...]\n"))
		return
	}
	kvStore.MSET(command[1:]...)
	conn.Write([]byte("OK\n"))
}

func handleMGet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) < 2 {
		conn.Write([]byte("Usage: MGET <key1> [<key2> ...]\n"))
		return
	}
	values := kvStore.MGET(command[1:]...)
	response := strings.Join(values, "\n") + "\n"
	conn.Write([]byte(response))
}


func handleUpdate(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: UPDATE <key> <value>\n"))
		return
	}
	oldValue := kvStore.Update(command[1], command[2])
	if oldValue == "" {
		conn.Write([]byte("Key does not exist\n"))
	} else {
		conn.Write([]byte("OK\n"))
	}
}

func handleGetSet(conn net.Conn, kvStore *store.KeyValueStore, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: GETSET <key> <value>\n"))
		return
	}
	oldValue := kvStore.GetSet(command[1], command[2])
	conn.Write([]byte(oldValue + "\n"))
}