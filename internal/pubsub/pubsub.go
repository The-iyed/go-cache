package pubsub

import (
	"sync"
)

type PubSub struct {
	mu       sync.RWMutex
	channels map[string][]chan string
}

func New() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(channel string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string)
	ps.channels[channel] = append(ps.channels[channel], ch)
	return ch
}

func (ps *PubSub) Publish(channel string, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if subscribers, ok := ps.channels[channel]; ok {
		for _, subscriber := range subscribers {
			subscriber <- message
		}
	}
}

func (ps *PubSub) Unsubscribe(channel string, ch <-chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers := ps.channels[channel]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			ps.channels[channel] = append(subscribers[:i], subscribers[i+1:]...)
			close(subscriber)
			break
		}
	}
}

func (ps *PubSub) GetNumSubscribers(channel string) int {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subscribers := ps.channels[channel]
	return len(subscribers)
}
