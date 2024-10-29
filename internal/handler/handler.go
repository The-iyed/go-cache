package handler

import (
	"bufio"
	"fmt"
	"net"
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
		default:
			conn.Write([]byte("Unknown command\n"))
		}
	}
}
