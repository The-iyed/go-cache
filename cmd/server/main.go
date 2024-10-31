package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-redis-v1/internal/handler"
	"github.com/go-redis-v1/internal/liststore"
	"github.com/go-redis-v1/internal/store"
)

func main() {
	kvStore := store.NewKeyValueStore()
	listStore := liststore.NewListStore()
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server is running on port 6379")
	fmt.Println("System is ready to accept connections")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handler.HandleConnection(conn, kvStore, listStore)
	}
}
