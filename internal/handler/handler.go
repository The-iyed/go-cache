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

		switch strings.ToUpper(command[0]) {
		case "SET":
			if len(command) != 3 {
				conn.Write([]byte("Usage: SET <key> <value>\n"))
				continue
			}
			kvStore.Set(command[1], command[2], 0)
			conn.Write([]byte("OK\n"))
		case "SETEX":
			if len(command) != 4 {
				conn.Write([]byte("Usage: SETEX <key> <value> <ttl>\n"))
				continue
			}
			ttl, err := time.ParseDuration(command[3] + "s")
			if err != nil {
				conn.Write([]byte("Invalid TTL\n"))
				continue
			}
			kvStore.Set(command[1], command[2], ttl)
			conn.Write([]byte("OK\n"))
		case "GET":
			if len(command) != 2 {
				conn.Write([]byte("Usage: GET <key>\n"))
				continue
			}
			value, exist := kvStore.Get(command[1])
			if !exist {
				conn.Write([]byte("(nil)\n"))
			} else {
				conn.Write([]byte(value + "\n"))
			}
		case "DEL":
			if len(command) != 2 {
				conn.Write([]byte("Usage: DEL <key>\n"))
				continue
			}
			kvStore.Delete(command[1])
			conn.Write([]byte("OK\n"))
		case "KEYS":
			if len(command) != 2 {
				conn.Write([]byte("Usage: KEYS <pattern>\n"))
				continue
			}
			keys := kvStore.Keys(command[1])
			var response strings.Builder
			for i, key := range keys {
				response.WriteString(fmt.Sprintf("%d) \"%s\"\n", i+1, key))
			}
			conn.Write([]byte(response.String()))
		case "EXISTS":
			if len(command) != 2 {
				conn.Write([]byte("Usage: EXISTS <key>\n"))
				continue
			}
			exist := kvStore.Exist(command[1])
			if exist {
				conn.Write([]byte("(integer) 1\n"))
			} else {
				conn.Write([]byte("(integer) 0\n"))
			}
		case "TTL":
			if len(command) != 2 {
				conn.Write([]byte("Usage: TTL <key>\n"))
				continue
			}
			ttl := kvStore.TTL(command[1])
			if ttl == -2 {
				conn.Write([]byte("(nil)\n"))
			} else {
				conn.Write([]byte(fmt.Sprintf("(integer) %d\n", ttl)))
			}
		case "FLUSHALL":
			if len(command) != 1 {
				conn.Write([]byte("Usage: FLUSHALL\n"))
				continue
			}
			kvStore.FlushAll()
			conn.Write([]byte("OK\n"))
		case "INFO":
			if len(command) != 1 {
				conn.Write([]byte("Usage: INFO\n"))
				continue
			}
			info := kvStore.Info()
			conn.Write([]byte(info + "\n"))
		case "PING":
			if len(command) != 1 {
				conn.Write([]byte("Usage: PING\n"))
				continue
			}
			ping := kvStore.Ping()
			conn.Write([]byte(ping + "\n"))
		case "PERSIST":
			if len(command) != 2 {
				conn.Write([]byte("Usage: PERSIST <key>\n"))
				continue
			}
			if kvStore.Persist(command[1]) {
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("(nil)\n"))
			}
		case "EXPIRE":
			if len(command) != 3 {
				conn.Write([]byte("Usage: EXPIRE <key> <seconds>\n"))
				continue
			}
			seconds, err := strconv.Atoi(command[2])
			if err != nil {
				conn.Write([]byte("Invalid seconds\n"))
				continue
			}
			if kvStore.Expire(command[1], seconds) {
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("(nil)\n"))
			}
		case "MSET":
			if len(command) < 3 || len(command)%2 == 0 {
				conn.Write([]byte("Usage: MSET <key1> <value1> [<key2> <value2> ...]\n"))
				continue
			}
			kvStore.MSET(command[1:]...)
			conn.Write([]byte("OK\n"))

		case "MGET":
			if len(command) < 2 {
				conn.Write([]byte("Usage: MGET <key1> [<key2> ...]\n"))
				continue
			}
			values := kvStore.MGET(command[1:]...)
			response := strings.Join(values, "\n") + "\n"
			conn.Write([]byte(response))
		case "UPDATE":
			if len(command) != 3 {
				conn.Write([]byte("Usage: UPDATE <key> <value>\n"))
				continue
			}
			oldValue := kvStore.Update(command[1], command[2])
			if oldValue == "" {
				conn.Write([]byte("Key does not exist\n"))
			} else {
				conn.Write([]byte("OK\n"))
			}

		case "GETSET":
			if len(command) != 3 {
				conn.Write([]byte("Usage: GETSET <key> <value>\n"))
				continue
			}
			oldValue := kvStore.GetSet(command[1], command[2])
			conn.Write([]byte(oldValue + "\n"))
		default:
			conn.Write([]byte("Unknown command\n"))
		}
	}
}
