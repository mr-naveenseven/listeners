package listener

import (
	"fmt"
	"listeners/message"
	"sync"
	"time"
)

type Listener struct {
	Name     string
	Message  *message.Message
	ticker   *time.Ticker
	Done     chan struct{}
	wg       *sync.WaitGroup
	stopOnce sync.Once
}

func NewListener(name string, tickerTime time.Duration, Message *message.Message, wg *sync.WaitGroup) *Listener {
	return &Listener{
		Name:    name,
		ticker:  time.NewTicker(tickerTime),
		Done:    make(chan struct{}),
		Message: Message,
		wg:      wg,
	}
}

func (l *Listener) listenerfunc() {
	defer l.wg.Done()
	for {
		select {
		case <-l.Done:
			fmt.Printf("%s stopped!\n", l.Name)
			return
		case <-l.ticker.C:
			// fmt.Println("Tick at", t.String())
			fmt.Printf("Message data from %s: %s\n", l.Name, l.Message.ReadMessage())
		}
	}
}

func (l *Listener) Start() {
	l.wg.Add(1)
	go l.listenerfunc()
}

func (l *Listener) Stop() {
	l.stopOnce.Do(func() {
		l.ticker.Stop()
		close(l.Done)
	})
}
