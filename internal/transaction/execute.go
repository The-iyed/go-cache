package transaction

import (
	"net"

	handler "github.com/go-redis-v1/internal/Handler"
	"github.com/go-redis-v1/internal/jsonstore"
	"github.com/go-redis-v1/internal/liststore"
	"github.com/go-redis-v1/internal/store"
)

func (t *Transaction) Execute(conn net.Conn,
	kvStore *store.KeyValueStore,
	listStore *liststore.ListStore,
	jsonStore *jsonstore.JSONStore) {

	if !t.IsActive {
		conn.Write([]byte("ERROR: Transaction already executed or discarded\n"))
		return
	}
	for _, cmd := range t.Commands {
		command := cmd.Args
		switch cmd.Name {
		case "SET":
			handler.HandleSet(conn, kvStore, command)
		case "SETEX":
			handler.HandleSetEX(conn, kvStore, command)
		case "GET":
			handler.HandleGet(conn, kvStore, command)
		case "DEL":
			handler.HandleDel(conn, kvStore, command)
		case "KEYS":
			handler.HandleKeys(conn, kvStore, command)
		case "EXISTS":
			handler.HandleExists(conn, kvStore, command)
		case "TTL":
			handler.HandleTTL(conn, kvStore, command)
		case "FLUSHALL":
			handler.HandleFlushAll(conn, kvStore)
		case "INFO":
			handler.HandleInfo(conn, kvStore)
		case "PING":
			handler.HandlePing(conn, kvStore)
		case "PERSIST":
			handler.HandlePersist(conn, kvStore, command)
		case "EXPIRE":
			handler.HandleExpire(conn, kvStore, command)
		case "MSET":
			handler.HandleMSet(conn, kvStore, command)
		case "MGET":
			handler.HandleMGet(conn, kvStore, command)
		case "UPDATE":
			handler.HandleUpdate(conn, kvStore, command)
		case "GETSET":
			handler.HandleGetSet(conn, kvStore, command)
		case "PUBLISH":
			handler.HandlePublish(conn, command)
		case "SUBSCRIBE":
			handler.HandleSubscribe(conn, command)
		case "UNSUBSCRIBE":
			handler.HandleUnsubscribe(conn, command)
		case "GETNSUM":
			handler.HandleGetNumSub(conn, command)
		case "PSUBSCRIBE":
			handler.HandlePatternSubscribe(conn, command)
		case "PUNSUBSCRIBE":
			handler.HandlePatternUnsubscribe(conn, command)
		case "LPUSH":
			handler.HandleLPUSH(conn, command, listStore)
		case "RPUSH":
			handler.HandleRPUSH(conn, command, listStore)
		case "LPOP":
			handler.HandleLPOP(conn, command, listStore)
		case "RPOP":
			handler.HandleLPOP(conn, command, listStore)
		case "LRANGE":
			handler.HandleLRANGE(conn, command, listStore)
		case "LLEN":
			handler.HandleLLEN(conn, command, listStore)
		case "LTRIM":
			handler.HandleLTRIM(conn, command, listStore)
		case "LINDEX":
			handler.HandleLINDEX(conn, command, listStore)
		case "JSON.SET":
			handler.HandleSetJSON(conn, jsonStore, command)
		case "JSON.GET":
			handler.HandleGetJSON(conn, jsonStore, command)
		case "JSON.DEL":
			handler.HandleDeleteJSON(conn, jsonStore, command)
		case "JSON.UPDATE":
			handler.HandleUpdateJSON(conn, jsonStore, command)
		case "JSON.TTL":
			handler.HandleTTLJSON(conn, jsonStore, command)
		default:
			conn.Write([]byte("ERROR: Unknown command " + cmd.Name + "\n"))
			return
		}
	}
	t.IsActive = false

}
