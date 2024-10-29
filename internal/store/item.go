package store

import "time"

type Item struct {
	value   string
	expires time.Time
}

func (i *Item) IsExpired() bool {
	return !i.expires.IsZero() && time.Now().After(i.expires)
}
