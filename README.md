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
  - `GETNSUM <channel>` - Retrieve the number of subscribers for a given channel.

- **List Operations**
  - `LPUSH <key> <value>` - Insert a value at the head of the list.
  - `RPUSH <key> <value>` - Insert a value at the tail of the list.
  - `LPOP <key>` - Remove and return the first element of the list.
  - `RPOP <key>` - Remove and return the last element of the list.
  - `LRANGE <key> <start> <stop>` - Retrieve a range of elements from the list.
  - `LLEN <key>` - Returns the length of the list stored at key. If key does not exist, it is interpreted as an empty list and 0 is returned.
  - `LTRIM <key> <start> <stop>` - rim an existing list so that it will contain only the specified range of elements specified.
  - `LINDEX <key> <index>` - Retrieve the element at a specific index in the list. Supports negative indices, where -1 is the last element, -2 is the second last, and so on.

- **JSON Store Operations**
  - `JSON.SET <key> <value>` - Store a JSON object at the specified key.
  - `JSON.GET <key>` - Retrieve a JSON object from the specified key.
  - `JSON.DEL <key>` - Delete a JSON object from the store.
  - `JSON.UPDATE <key> <field> <value>` - Update a field in a JSON object.
  - `JSON.TTL <key>` - Get the remaining time to live for a JSON object.

- **Set Operations**
  - `SADD <key> <value>` - Add a member to a set.
  - `SREM <key> <value>` - Remove a member from a set.
  - `SMEMBERS <key>` - Get all the members of a set.

- **Pattern-Based Subscription**
  - `PSUBSCRIBE <pattern>` - Subscribe to channels that match a given pattern. This allows a client to receive messages published to any channel matching the pattern.
  - `PUNSUBSCRIBE <pattern>` - Unsubscribe from channels that match the specified pattern, stopping message delivery.

- **Automatic Expiration**
  - Automatically cleans up expired keys based on TTL.

- **Pub/Sub Feature**
  - **SUBSCRIBE <channel>** - Subscribe to a specified channel to receive messages published to that channel.
  - **PUBLISH <channel> <message>** - Publish a message to a specified channel that all subscribers will receive.
  - **UNSUBSCRIBE <channel>** - Unsubscribe from a specified channel, stopping message delivery.
  - **Message Persistence** - Stores recent messages in a buffer for each channel, allowing new subscribers to catch up on recent messages upon joining.

- **Transaction Support**
  - `MULTI` - Start a transaction. This command marks the beginning of a transaction block.
  - `EXEC` - Execute all commands issued after `MULTI` and discard the transaction if an error occurs.
  - `DISCARD` - Abort the transaction, discarding all commands issued after `MULTI`.

## Future Features
- **Persistence** - Save cache data to disk and load on startup.
- **Advanced Data Structures** - Implement more complex data structures such as sorted sets and hashes.
- **Replication and Clustering** - Scale out by distributing data across multiple instances.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/ICache.git
    cd ICache
    ```
2. Run the server:
    ```bash
    docker compose up
    ```

## Usage

Connect to ICache using any TCP client (e.g., `telnet` or `nc`):

```bash
telnet localhost 6379
