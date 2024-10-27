package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
    "sync"
    "time"
)

type Item struct {
  value string
  expiry time.Time
}

type KeyValueStore struct {
    data map[string]Item
    mu   sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
    store := &KeyValueStore{
        data: make(map[string]Item),
    }
    go store.Clean()
    return store 
}

func (store *KeyValueStore) Clean(){
  ticker := time.NewTicker(5 * time.Second)
  defer ticker.Stop()

  for range ticker.C {
    expired := []string{}
    now := time.Now()

    store.mu.RLock()
    for key , item := range store.data {
      if !item.expires.IsZero() && now.After(item.expires) {
        expired = append(expired, key)
      }
    }

    store.mu.RUnlock()

    if len(expired) > 0 {
      store.mu.Lock()
      for _ , key := range expired {
        delete(store.data,key)
      }
      store.mu.Unlock()
    }

  }
}

func (store *KeyValueStore) Set(key, value string,ttl time.Duration) {
    store.mu.Lock()
    defer store.mu.Unlock()

    expires := time.Time{}
    if ttl > 0 {
      expires = time.Now().Add(ttl)
    }

    store.data[key] = &Item{
      value:value,
      expires:expires
    }
}

func (store *KeyValueStore) Get(key string) (string, bool) {
    store.mu.RLock()
    defer store.mu.RUnlock()
    item , exist := store.data[key]

    if !exist {
      return "" , false  
    }

    if !item.expires.IsZero() && time.Now().After(item.expires) {
      store.mu.Lock()
      delete(store.data,key)
      store.mu.Unlock()
      return "" , false
    }

    return item.value, false
}

func (store *KeyValueStore) Delete(key string) {
    store.mu.Lock()
    defer store.mu.Unlock()
    delete(store.data, key)
}

func handleConnection(conn net.Conn, store *KeyValueStore) {
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
            store.Set(command[1], command[2],0)
            conn.Write([]byte("OK\n"))
        case "SETEX":
            if len(command) != 4 {
              conn.Write([]byte("Usage: SETEX <key> <value> <ttl>"))
              continue
            }
            ttl, err := time.ParseDuration(command[3] + "s")
            if err != nil {
              conn.Write([]byte("Invalid TTL\n"))
              continue
            }
            store.Set(command[1],command[2],command[3])
            conn.Write([]byte("OK\n"))
        case "GET":
            if len(command) != 2 {
                conn.Write([]byte("Usage: GET <key>\n"))
                continue
            }
            value, exist := store.Get(command[1])
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
            store.Delete(command[1])
            conn.Write([]byte("OK\n"))
        default:
            conn.Write([]byte("Unknown command\n"))
        }
    }
}

func main() {
    store := NewKeyValueStore()
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
        go handleConnection(conn, store)
    }
}

