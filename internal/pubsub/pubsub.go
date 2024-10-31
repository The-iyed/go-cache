package pubsub

import (
	"strings"
	"sync"
)

const BufferSize = 30

type PubSub struct {
	mu       sync.RWMutex
	channels map[string][]chan string
	patterns map[string][]chan string
	buffers  map[string][]string
}

func New() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan string),
		buffers:  make(map[string][]string),
		patterns: make(map[string][]chan string),
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

func (ps *PubSub) SubscribePattern(pattern string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan string, 1)
	ps.patterns[pattern] = append(ps.patterns[pattern], ch)
	return ch
}

func (ps *PubSub) Publish(channel string, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	ps.storeInBuffer(channel, message)

	if subscribers, ok := ps.channels[channel]; ok {
		for _, subscriber := range subscribers {
			go func(sub chan string) {
				sub <- message
			}(subscriber)
		}
	}

	for pattern, subscribers := range ps.patterns {
		if matchesPattern(pattern, channel) {
			for _, subscriber := range subscribers {
				go func(sub chan string) {
					sub <- message
				}(subscriber)
			}
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

func (ps *PubSub) UnsubscribePattern(pattern string, ch <-chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if subscribers, ok := ps.patterns[pattern]; ok {
		for i, subscriber := range subscribers {
			if subscriber == ch {
				ps.patterns[pattern] = append(subscribers[:i], subscribers[i+1:]...)
				close(subscriber)
				break
			}
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

func matchesPattern(pattern, channel string) bool {
	if pattern == "*" {
		return true
	}
	parts := strings.Split(pattern, "*")
	if len(parts) == 1 {
		return pattern == channel
	}

	for i, part := range parts {
		if i == 0 && !strings.HasPrefix(channel, part) {
			return false
		}
		if i == len(parts)-1 && !strings.HasSuffix(channel, part) {
			return false
		}
		if !strings.Contains(channel, part) {
			return false
		}
	}
	return true
}
