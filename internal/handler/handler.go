package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/go-redis-v1/internal/jsonstore"
	"github.com/go-redis-v1/internal/liststore"
	"github.com/go-redis-v1/internal/pubsub"
	"github.com/go-redis-v1/internal/store"
	"github.com/go-redis-v1/internal/transaction"
)

var (
	pubSubStore = pubsub.New()
)

func HandleConnection(conn net.Conn,
	kvStore *store.KeyValueStore,
	listStore *liststore.ListStore,
	jsonStore *jsonstore.JSONStore,
	transaction *transaction.Transaction) {
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
		case "MULTI":
			transaction.StartTransaction()
			conn.Write([]byte("OK: Transaction started\n"))
		case "EXEC":
			CommitTransaction(conn, kvStore, listStore, jsonStore, transaction)
		case "DISCARD":
			err := transaction.AbortTransaction()
			if err != nil {
				conn.Write([]byte("ERR: " + err.Error() + "\n"))
				return
			}
			conn.Write([]byte("OK: Transaction discarded\n"))
		default:
			HandleCommand(conn, kvStore, listStore, jsonStore, transaction, command)
		}
	}
}

func HandleCommand(conn net.Conn,
	kvStore *store.KeyValueStore,
	listStore *liststore.ListStore,
	jsonStore *jsonstore.JSONStore,
	transaction *transaction.Transaction,
	command []string) {

	if transaction.IsActive {
		transaction.AddCommand(strings.ToUpper(command[0]), command)
		conn.Write([]byte("QUEUED\n"))
		return
	}

	switch strings.ToUpper(command[0]) {
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
		HandleRPOP(conn, command, listStore)
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
		conn.Write([]byte("Unknown command\n"))
	}
}
