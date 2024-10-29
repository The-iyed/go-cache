# ICache

ICache is a high-performance, in-memory key-value store designed to provide fast data retrieval and caching functionality, similar to Redis. This project aims to build a simple, scalable cache solution with minimal dependencies, built entirely in Go.

## Features

- **Data Storage Commands**
  - `SET <key> <value>` - Store a key-value pair.
  - `GET <key>` - Retrieve the value of a key.
  - `DEL <key>` - Delete a key-value pair.
  - `SETEX <key> <value> <ttl>` - Store a key-value pair with an expiration time (TTL).
- **Automatic Expiration**
  - Automatically cleans up expired keys based on TTL.

## Future Features
- **Persistence** - Save cache data to disk and load on startup.
- **Advanced Data Structures** - Implement lists, sets, and other structures.
- **Replication and Clustering** - Scale out by distributing data across multiple instances.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/ICache.git
    cd ICache
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Run the server:
    ```bash
    go run cmd/server/main.go
    ```

## Usage

Connect to ICache using any TCP client (e.g., `telnet` or `nc`):

```bash
telnet localhost 6379
