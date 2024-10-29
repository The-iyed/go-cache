package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/go-redis-v1/internal/pubsub"
	"github.com/go-redis-v1/internal/store"
)

var (
	pubSubStore = pubsub.New()
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
	case "PUBLISH":
		handlePublish(conn, command)
	case "SUBSCRIBE":
		handleSubscribe(conn, command)
	case "UNSUBSCRIBE":
		handleUnsubscribe(conn, command)
	default:
		conn.Write([]byte("Unknown command\n"))
	}
}

func handlePublish(conn net.Conn, command []string) {
	if len(command) != 3 {
		conn.Write([]byte("Usage: PUBLISH <channel> <message>\n"))
		return
	}
	channel := command[1]
	message := command[2]
	pubSubStore.Publish(channel, message)
	conn.Write([]byte("Message published\n"))
}

func handleSubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: SUBSCRIBE <channel>\n"))
		return
	}
	channel := command[1]
	messageChannel := make(chan string)
	pubSubStore.Subscribe(channel)

	go func() {
		for message := range messageChannel {
			conn.Write([]byte(message + "\n"))
		}
	}()

	conn.Write([]byte("Subscribed to channel: " + channel + "\n"))
}

func handleUnsubscribe(conn net.Conn, command []string) {
	if len(command) != 2 {
		conn.Write([]byte("Usage: UNSUBSCRIBE <channel>\n"))
		return
	}
	channel := command[1]
	messageChannel := make(chan string)
	pubSubStore.Unsubscribe(channel, messageChannel)

	close(messageChannel)
	conn.Write([]byte("Unsubscribed from channel: " + channel + "\n"))
}
