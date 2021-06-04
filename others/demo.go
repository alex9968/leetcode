package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m sync.RWMutex
	buffer int
	timeout time.Duration
	subscribers map[subscriber]topicFunc
}


func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer: buffer,
		timeout: publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.send2Topic(sub, topic, v, &wg)
	}
	wg.Wait()
}
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) send2Topic(sub chan interface{}, topic topicFunc, msg interface{}, wg *sync.WaitGroup	 ) {
	defer wg.Done()
	if topic != nil && !topic(msg) {
		return
	}

	select {
	case sub <- msg:
		case <-time.After(p.timeout):
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//
	all := p.Subscribe()

	golang := p.SubscribeTopic( func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return  strings.Contains(s, "golang")
		}
		return false
	})


	p.Publish("golang 111")
	p.Publish("msg 111")
	p.Publish("msg2 111")

	go func() {
		for msg := range all {
			fmt.Println("Grobal:", msg)
		}
	}()
	go func() {
		for msg := range golang {
			fmt.Println("Golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}




