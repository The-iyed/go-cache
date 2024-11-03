package handler

import (
	"fmt"
	"net"
	"strconv"

	"github.com/go-redis-v1/internal/liststore"
)

func HandleLPUSH(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) < 3 {
		conn.Write([]byte("Usage: LPUSH <key> <value1> <value2> ...\n"))
		return
	}
	key := command[1]
	values := command[2:]
	listStore.LPUSH(key, values...)
	conn.Write([]byte("OK\n"))
}

func HandleRPUSH(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) < 3 {
		conn.Write([]byte("Usage: RPUSH <key> <value1> <value2> ...\n"))
		return
	}
	key := command[1]
	values := command[2:]
	listStore.RPUSH(key, values...)
	conn.Write([]byte("OK\n"))
}

func HandleLPOP(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: LPOP <key>\n"))
		return
	}
	key := command[1]
	value, ok := listStore.LPOP(key)
	if ok {
		conn.Write([]byte(value + "\n"))
	} else {
		conn.Write([]byte("nil\n"))
	}
}

func HandleRPOP(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: RPOP <key>\n"))
		return
	}
	key := command[1]
	value, ok := listStore.RPOP(key)
	if ok {
		conn.Write([]byte(value + "\n"))
	} else {
		conn.Write([]byte("nil\n"))
	}
}

func HandleLRANGE(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) != 4 {
		conn.Write([]byte("Usage: LRANGE <key> <start> <stop>\n"))
		return
	}
	key := command[1]
	start, err1 := strconv.Atoi(command[2])
	stop, err2 := strconv.Atoi(command[3])
	if err1 != nil || err2 != nil {
		conn.Write([]byte("Invalid start or stop index\n"))
		return
	}
	values, err := listStore.LRANGE(key, start, stop)
	if err != nil {
		conn.Write([]byte("nil\n"))
		return
	}
	if len(values) == 0 {
		conn.Write([]byte("nil\n"))
		return
	}
	for _, value := range values {
		conn.Write([]byte(value + "\n"))
	}
}

func HandleLLEN(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: LEN <key>\n"))
		return
	}
	key := command[1]
	length := strconv.Itoa(listStore.LLEN(key))
	conn.Write([]byte(length + "\n"))
}

func HandleLTRIM(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) != 4 {
		conn.Write([]byte("Usage: LRANGE <key> <start> <stop>\n"))
		return
	}
	key := command[1]
	start, err1 := strconv.Atoi(command[2])
	stop, err2 := strconv.Atoi(command[3])
	if err1 != nil || err2 != nil {
		conn.Write([]byte("Invalid start or stop index\n"))
		return
	}
	err := listStore.LTRIM(key, start, stop)
	if err != nil {
		conn.Write([]byte("nil\n"))
		return
	}

	conn.Write([]byte("OK\n"))

}

func HandleLINDEX(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) < 3 {
		conn.Write([]byte("ERR: LINDEX requires key and index\n"))
		return
	}
	index, err := strconv.Atoi(command[2])
	if err != nil {
		conn.Write([]byte("ERR: Invalid index\n"))
		return
	}

	value, err := listStore.LINDEX(command[1], index)
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("ERR: %v\n", err)))
	} else {
		conn.Write([]byte(fmt.Sprintf("%s\n", value)))
	}
}
