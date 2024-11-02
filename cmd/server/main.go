package main

import (
	"log"
	"net"

	"github.com/go-redis-v1/internal/handler"
	"github.com/go-redis-v1/internal/jsonstore"
	"github.com/go-redis-v1/internal/liststore"
	"github.com/go-redis-v1/internal/store"
	"github.com/go-redis-v1/logger"
)

func main() {
	kvStore := store.NewKeyValueStore()
	listStore := liststore.NewListStore()
	jsonStore := jsonstore.NewJSONStore()
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	logger.Info("Server is running on port 6379")
	logger.Info("Server initialized")
	logger.Info("Ready to accept connections tcp")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handler.HandleConnection(conn, kvStore, listStore, jsonStore)
	}
}
