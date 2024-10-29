package pubsub

import (
	"sync"
)

const BufferSize = 30

type PubSub struct {
	mu       sync.RWMutex
	channels map[string][]chan string
	buffers  map[string][]string
}

func New() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan string),
		buffers:  make(map[string][]string),
	}
}

func (ps *PubSub) Subscribe(channel string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 1)
	if buffer, ok := ps.buffers[channel]; ok {
		go func() {
			for _, msg := range buffer {
				ch <- msg
			}
		}()
	}
	ps.channels[channel] = append(ps.channels[channel], ch)
	return ch
}

func (ps *PubSub) Publish(channel string, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	ps.storeInBuffer(channel, message)

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

func (ps *PubSub) storeInBuffer(channel, message string) {
	buffer := ps.buffers[channel]
	buffer = append(buffer, message)
	if len(buffer) > BufferSize {
		buffer = buffer[1:]
	}
	ps.buffers[channel] = buffer
}
