package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis-v1/internal/jsonstore"
)

func HandleSetJSON(conn net.Conn, jsonStore *jsonstore.JSONStore, command []string) {
	if len(command) < 3 {
		conn.Write([]byte("ERR wrong number of arguments for 'SETJSON' command\n"))
		return
	}

	key := command[1]
	jsonStr := strings.Join(command[2:len(command)-1], " ")
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		conn.Write([]byte("ERR invalid JSON\n"))
		return
	}

	ttl := time.Duration(0)
	if len(command) > 3 {
		ttlParsed, err := strconv.Atoi(command[len(command)-1])
		if err == nil {
			ttl = time.Duration(ttlParsed) * time.Second
		}
	}

	if err := jsonStore.SetJSON(key, data, ttl); err != nil {
		conn.Write([]byte(fmt.Sprintf("ERR %s\n", err.Error())))
		return
	}
	conn.Write([]byte("OK\n"))
}

func HandleGetJSON(conn net.Conn, jsonStore *jsonstore.JSONStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("ERR wrong number of arguments for 'GETJSON' command\n"))
		return
	}

	key := command[1]
	var result map[string]interface{}
	err := jsonStore.GetJSON(key, &result)
	if err != nil {
		conn.Write([]byte("ERR no such key\n"))
		return
	}

	data, _ := json.Marshal(result)
	conn.Write(data)
	conn.Write([]byte("\n"))
}

func HandleDeleteJSON(conn net.Conn, jsonStore *jsonstore.JSONStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("ERR wrong number of arguments for 'DELJSON' command\n"))
		return
	}

	key := command[1]
	if jsonStore.DeleteJSON(key) {
		conn.Write([]byte(":1\n"))
	} else {
		conn.Write([]byte(":0\n"))
	}
}

func HandleUpdateJSON(conn net.Conn, jsonStore *jsonstore.JSONStore, command []string) {
	if len(command) < 4 {
		conn.Write([]byte("ERR wrong number of arguments for 'UPDATEJSON' command\n"))
		return
	}

	key := command[1]
	field := command[2]
	value := command[3]

	var valueInterface interface{}
	if err := json.Unmarshal([]byte(value), &valueInterface); err != nil {
		conn.Write([]byte("ERR invalid JSON value\n"))
		return
	}

	if err := jsonStore.UpdateJSON(key, field, valueInterface); err != nil {
		conn.Write([]byte(fmt.Sprintf("ERR %s\n", err.Error())))
		return
	}
	conn.Write([]byte("OK\n"))
}

func HandleTTLJSON(conn net.Conn, jsonStore *jsonstore.JSONStore, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("ERR wrong number of arguments for 'TTL' command\n"))
		return
	}

	key := command[1]
	ttl, err := jsonStore.TTL(key)
	if err != nil {
		conn.Write([]byte("-2\n"))
		return
	}

	conn.Write([]byte(fmt.Sprintf(":%d\n", int(ttl.Seconds()))))
}
