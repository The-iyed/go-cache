package handler

import (
	"fmt"
	"net"

	"github.com/go-redis-v1/internal/jsonstore"
	"github.com/go-redis-v1/internal/liststore"
	"github.com/go-redis-v1/internal/store"
	"github.com/go-redis-v1/internal/transaction"
)

func CommitTransaction(conn net.Conn,
	kvStore *store.KeyValueStore,
	listStore *liststore.ListStore,
	jsonStore *jsonstore.JSONStore,
	transaction *transaction.Transaction) {

	if !transaction.IsActive {
		conn.Write([]byte("ERROR: Transaction already executed or discarded\n"))
		return
	}
	for _, cmd := range transaction.Commands {
		command := cmd.Args
		switch cmd.Name {
		case "SET":
			HandleSet(conn, kvStore, command)
		case "SETEX":
			HandleSetEX(conn, kvStore, command)
		case "GET":
			HandleGet(conn, kvStore, command)
		case "DEL":
			HandleDel(conn, kvStore, command)
		case "KEYS":
			HandleKeys(conn, kvStore, command)
		case "EXISTS":
			HandleExists(conn, kvStore, command)
		case "TTL":
			HandleTTL(conn, kvStore, command)
		case "FLUSHALL":
			HandleFlushAll(conn, kvStore)
		case "INFO":
			HandleInfo(conn, kvStore)
		case "PING":
			HandlePing(conn, kvStore)
		case "PERSIST":
			HandlePersist(conn, kvStore, command)
		case "EXPIRE":
			HandleExpire(conn, kvStore, command)
		case "MSET":
			HandleMSet(conn, kvStore, command)
		case "MGET":
			HandleMGet(conn, kvStore, command)
		case "UPDATE":
			HandleUpdate(conn, kvStore, command)
		case "GETSET":
			HandleGetSet(conn, kvStore, command)
		case "PUBLISH":
			HandlePublish(conn, command)
		case "SUBSCRIBE":
			HandleSubscribe(conn, command)
		case "UNSUBSCRIBE":
			HandleUnsubscribe(conn, command)
		case "GETNSUM":
			HandleGetNumSub(conn, command)
		case "PSUBSCRIBE":
			HandlePatternSubscribe(conn, command)
		case "PUNSUBSCRIBE":
			HandlePatternUnsubscribe(conn, command)
		case "LPUSH":
			HandleLPUSH(conn, command, listStore)
		case "RPUSH":
			HandleRPUSH(conn, command, listStore)
		case "LPOP":
			HandleLPOP(conn, command, listStore)
		case "RPOP":
			HandleLPOP(conn, command, listStore)
		case "LRANGE":
			HandleLRANGE(conn, command, listStore)
		case "LLEN":
			HandleLLEN(conn, command, listStore)
		case "LTRIM":
			HandleLTRIM(conn, command, listStore)
		case "LINDEX":
			HandleLINDEX(conn, command, listStore)
		case "JSON.SET":
			HandleSetJSON(conn, jsonStore, command)
		case "JSON.GET":
			HandleGetJSON(conn, jsonStore, command)
		case "JSON.DEL":
			HandleDeleteJSON(conn, jsonStore, command)
		case "JSON.UPDATE":
			HandleUpdateJSON(conn, jsonStore, command)
		case "JSON.TTL":
			HandleTTLJSON(conn, jsonStore, command)
		default:
			conn.Write([]byte("ERROR: Unknown command " + cmd.Name + "\n"))
			transaction.IsActive = false
			return
		}
	}
	transaction.IsActive = false
	transaction.Commands = nil
	conn.Write([]byte("OK: Transaction committed\n"))
}
