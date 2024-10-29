# ICache

ICache is a high-performance, in-memory key-value store designed to provide fast data retrieval and caching functionality, similar to Redis. This project aims to build a simple, scalable cache solution with minimal dependencies, built entirely in Go.

## Features

- **Data Storage Commands**
  - `SET <key> <value>` - Store a key-value pair.
  - `SETEX <key> <value> <ttl>` - Store a key-value pair with an expiration time (TTL).
  - `GET <key>` - Retrieve the value of a key.
  - `DEL <key>` - Delete a key-value pair.
  - `EXISTS <key>` - Check if the specified key exists.
  - `TTL <key>` - Get the remaining time to live for a key (returns the TTL in seconds, `-1` if the key has no expiration, or `-2` if the key does not exist).
  - `KEYS <pattern>` - Retrieve all keys matching a pattern, with responses formatted like Redis.
  - `FLUSHALL` - Delete all keys from all databases.
  - `INFO` - Get information and statistics about the server, including memory usage and other details.
  - `PING` - Check the connection to the server.
  - `EXPIRE <key> <seconds>` - Set a timeout on the specified key after which it will be automatically deleted.
  - `PERSIST <key>` - Remove the expiration from a key, making it persistent.
  - `MSET <key1> <value1> [<key2> <value2> ...]` - Set multiple key-value pairs at once.
  - `MGET <key1> [<key2> ...]` - Retrieve multiple values for the given keys.
  - `UPDATE <key> <value>` - Update the value of a key if it exists.
  - `GETSET <key> <value>` - Set the value of a key and return its old value.

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
