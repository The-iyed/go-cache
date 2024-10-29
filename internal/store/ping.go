package store

func (store *KeyValueStore) Ping() string {
	return "PONG"
}
