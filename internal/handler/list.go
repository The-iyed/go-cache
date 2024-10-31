package handler

import (
	"net"
	"strconv"

	"github.com/go-redis-v1/internal/liststore"
)


func handleLPUSH(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) < 3 {
		conn.Write([]byte("Usage: LPUSH <key> <value1> <value2> ...\n"))
		return
	}
	key := command[1]
	values := command[2:]
	listStore.LPUSH(key, values...) 
	conn.Write([]byte("OK\n"))
}


func handleRPUSH(conn net.Conn, command []string, listStore *liststore.ListStore) {
	if len(command) < 3 {
		conn.Write([]byte("Usage: RPUSH <key> <value1> <value2> ...\n"))
		return
	}
	key := command[1]
	values := command[2:]
	listStore.RPUSH(key, values...) 
	conn.Write([]byte("OK\n"))
}


func handleLPOP(conn net.Conn, command []string, listStore *liststore.ListStore) {
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


func handleRPOP(conn net.Conn, command []string, listStore *liststore.ListStore) {
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


func handleLRANGE(conn net.Conn, command []string, listStore *liststore.ListStore) {
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
	values := listStore.LRANGE(key, start, stop)
	if len(values) == 0 {
		conn.Write([]byte("nil\n"))  
		return
	}
	for _, value := range values {
		conn.Write([]byte(value + "\n"))  
	}
}
